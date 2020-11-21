package bitmapdb

import (
	"bytes"
	"encoding/binary"
	"sort"

	"github.com/RoaringBitmap/roaring"
	"github.com/c2h5oh/datasize"
	"github.com/ledgerwatch/turbo-geth/common"
	"github.com/ledgerwatch/turbo-geth/common/math"
	"github.com/ledgerwatch/turbo-geth/ethdb"
)

const ChunkLimit = uint64(1900 * datasize.B) // threshold after which appear LMDB OverflowPages

func ChunkIterator(bm *roaring.Bitmap, target uint64) func() *roaring.Bitmap {
	return func() *roaring.Bitmap {
		return CutLeft(bm, target)
	}
}

// CutLeft - cut from bitmap `targetSize` bytes from left
// removing lft part from `bm`
// returns nil on zero cardinality
func CutLeft(bm *roaring.Bitmap, sizeLimit uint64) *roaring.Bitmap {
	if bm.GetCardinality() == 0 {
		return nil
	}

	sz := bm.GetSerializedSizeInBytes()
	if sz <= sizeLimit {
		lft := roaring.New()
		lft.Or(bm)
		bm.Clear()
		return lft
	}

	from := uint64(bm.Minimum())
	minMax := bm.Maximum() - bm.Minimum()             // +1 because AddRange has semantic [from,to)
	to := sort.Search(int(minMax), func(i int) bool { // can be optimized to avoid "too small steps", but let's leave it for readability
		lft := roaring.New() // bitmap.Clear() method intentionally not used here, because then serialized size of bitmap getting bigger
		lft.AddRange(from, from+uint64(i)+1)
		lft.And(bm)
		lft.RunOptimize()
		return lft.GetSerializedSizeInBytes() > sizeLimit
	})

	lft := roaring.New()
	lft.AddRange(from, from+uint64(to)+1)
	lft.And(bm)
	bm.RemoveRange(from, from+uint64(to)+1)
	lft.RunOptimize()
	return lft
}

func WalkChunks(bm *roaring.Bitmap, sizeLimit uint64, f func(chunk *roaring.Bitmap, isLast bool) error) error {
	for bm.GetCardinality() > 0 {
		if err := f(CutLeft(bm, sizeLimit), bm.GetCardinality() == 0); err != nil {
			return err
		}
	}
	return nil
}

func WalkChunkWithKeys(k []byte, m *roaring.Bitmap, sizeLimit uint64, f func(chunkKey []byte, chunk *roaring.Bitmap) error) error {
	return WalkChunks(m, sizeLimit, func(chunk *roaring.Bitmap, isLast bool) error {
		chunkKey := make([]byte, len(k)+4)
		copy(chunkKey, k)
		if isLast {
			binary.BigEndian.PutUint32(chunkKey[len(k):], ^uint32(0))
		} else {
			binary.BigEndian.PutUint32(chunkKey[len(k):], chunk.Maximum())
		}
		return f(chunkKey, chunk)
	})
}

// TruncateRange - gets existing bitmap in db and call RemoveRange operator on it.
// starts from hot shard, stops when shard not overlap with [from-to)
// !Important: [from, to)
func TruncateRange(db ethdb.Database, bucket string, key []byte, to uint32) error {
	chunkKey := make([]byte, len(key)+4)
	copy(chunkKey, key)
	binary.BigEndian.PutUint32(chunkKey[len(chunkKey)-4:], to)
	bm, err := Get(db, bucket, key, to, math.MaxUint32)
	if err != nil {
		return err
	}

	if bm.GetCardinality() > 0 && to <= bm.Maximum() {
		bm.RemoveRange(uint64(to), uint64(bm.Maximum())+1)
	}

	if err := db.Walk(bucket, chunkKey, 0, func(k, v []byte) (bool, error) {
		if !bytes.HasPrefix(k, key) {
			return false, nil
		}
		if err := db.Delete(bucket, k, nil); err != nil {
			return false, err
		}
		return true, nil
	}); err != nil {
		return err
	}

	buf := bytes.NewBuffer(nil)
	return WalkChunkWithKeys(key, bm, ChunkLimit, func(chunkKey []byte, chunk *roaring.Bitmap) error {
		buf.Reset()
		if _, err := chunk.WriteTo(buf); err != nil {
			return err
		}
		return db.Put(bucket, chunkKey, common.CopyBytes(buf.Bytes()))
	})
}

// Get - reading as much chunks as needed to satisfy [from, to] condition
// join all chunks to 1 bitmap by Or operator
func Get(db ethdb.Getter, bucket string, key []byte, from, to uint32) (*roaring.Bitmap, error) {
	var chunks []*roaring.Bitmap

	fromKey := make([]byte, len(key)+4)
	copy(fromKey, key)
	binary.BigEndian.PutUint32(fromKey[len(fromKey)-4:], from)

	if err := db.Walk(bucket, fromKey, len(key)*8, func(k, v []byte) (bool, error) {
		bm := roaring.New()
		_, err := bm.ReadFrom(bytes.NewReader(v))
		if err != nil {
			return false, err
		}
		chunks = append(chunks, bm)
		if binary.BigEndian.Uint32(k[len(k)-4:]) >= to {
			return false, nil
		}
		return true, nil
	}); err != nil {
		return nil, err
	}

	if len(chunks) == 0 {
		return roaring.New(), nil
	}
	return roaring.FastOr(chunks...), nil
}

// SeekInBitmap - returns value in bitmap which is >= n
func SeekInBitmap(m *roaring.Bitmap, n uint32) (found uint32, ok bool) {
	i := m.Iterator()
	i.AdvanceIfNeeded(n)
	ok = i.HasNext()
	if ok {
		found = i.Next()
	}
	return found, ok
}
