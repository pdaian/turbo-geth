// +build go1.6

// Code generated by codecgen - DO NOT EDIT.

package common

import (
	"errors"
	codec1978 "github.com/ugorji/go/codec"
	"runtime"
	"strconv"
)

const (
	// ----- content types ----
	codecSelferCcUTF81 = 1
	codecSelferCcRAW1  = 255
	// ----- value types used ----
	codecSelferValueTypeArray1     = 10
	codecSelferValueTypeMap1       = 9
	codecSelferValueTypeString1    = 6
	codecSelferValueTypeInt1       = 2
	codecSelferValueTypeUint1      = 3
	codecSelferValueTypeFloat1     = 4
	codecSelferValueTypeNil1       = 1
	codecSelferBitsize1            = uint8(32 << (^uint(0) >> 63))
	codecSelferDecContainerLenNil1 = -2147483648
)

var (
	errCodecSelferOnlyMapOrArrayEncodeToStruct1 = errors.New(`only encoded map or array can be decoded into a struct`)
)

type codecSelfer1 struct{}

func codecSelfer1False() bool { return false }

func init() {
	if codec1978.GenVersion != 17 {
		_, file, _, _ := runtime.Caller(0)
		ver := strconv.FormatInt(int64(codec1978.GenVersion), 10)
		panic(errors.New("codecgen version mismatch: current: 17, need " + ver + ". Re-generate file: " + file))
	}
}

func (x *Hash) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if !z.EncBinary() {
			z.EncTextMarshal(x)
		} else {
			h.encHash((*Hash)(x), e)
		}
	}
}

func (x *Hash) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if !z.DecBinary() && z.IsJSONHandle() {
		z.DecJSONUnmarshal(x)
	} else {
		h.decHash((*Hash)(x), d)
	}
}

func (x *UnprefixedHash) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if !z.EncBinary() {
			z.EncTextMarshal(x)
		} else {
			h.encUnprefixedHash((*UnprefixedHash)(x), e)
		}
	}
}

func (x *UnprefixedHash) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if !z.DecBinary() {
		z.DecTextUnmarshal(x)
	} else {
		h.decUnprefixedHash((*UnprefixedHash)(x), d)
	}
}

func (x *Address) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if !z.EncBinary() {
			z.EncTextMarshal(x)
		} else {
			h.encAddress((*Address)(x), e)
		}
	}
}

func (x *Address) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if !z.DecBinary() && z.IsJSONHandle() {
		z.DecJSONUnmarshal(x)
	} else {
		h.decAddress((*Address)(x), d)
	}
}

func (x *UnprefixedAddress) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if !z.EncBinary() {
			z.EncTextMarshal(x)
		} else {
			h.encUnprefixedAddress((*UnprefixedAddress)(x), e)
		}
	}
}

func (x *UnprefixedAddress) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if !z.DecBinary() {
		z.DecTextUnmarshal(x)
	} else {
		h.decUnprefixedAddress((*UnprefixedAddress)(x), d)
	}
}

func (x *MixedcaseAddress) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		if !z.EncBinary() && z.IsJSONHandle() {
			z.EncJSONMarshal(x)
		} else {
			yysep2 := !z.EncBinary()
			yy2arr2 := z.EncBasicHandle().StructToArray
			_, _ = yysep2, yy2arr2
			const yyr2 bool = false // struct tag has 'toArray'
			if yyr2 || yy2arr2 {
				z.EncWriteArrayStart(0)
				z.EncWriteArrayEnd()
			} else {
				z.EncWriteMapStart(0)
				z.EncWriteMapEnd()
			}
		}
	}
}

func (x *MixedcaseAddress) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	if !z.DecBinary() && z.IsJSONHandle() {
		z.DecJSONUnmarshal(x)
	} else {
		yyct2 := r.ContainerType()
		if yyct2 == codecSelferValueTypeNil1 {
			*(x) = MixedcaseAddress{}
		} else if yyct2 == codecSelferValueTypeMap1 {
			yyl2 := z.DecReadMapStart()
			if yyl2 == 0 {
			} else {
				x.codecDecodeSelfFromMap(yyl2, d)
			}
			z.DecReadMapEnd()
		} else if yyct2 == codecSelferValueTypeArray1 {
			yyl2 := z.DecReadArrayStart()
			if yyl2 != 0 {
				x.codecDecodeSelfFromArray(yyl2, d)
			}
			z.DecReadArrayEnd()
		} else {
			panic(errCodecSelferOnlyMapOrArrayEncodeToStruct1)
		}
	}
}

func (x *MixedcaseAddress) codecDecodeSelfFromMap(l int, d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyhl3 bool = l >= 0
	for yyj3 := 0; ; yyj3++ {
		if yyhl3 {
			if yyj3 >= l {
				break
			}
		} else {
			if z.DecCheckBreak() {
				break
			}
		}
		z.DecReadMapElemKey()
		yys3 := z.StringView(r.DecodeStringAsBytes())
		z.DecReadMapElemValue()
		switch yys3 {
		default:
			z.DecStructFieldNotFound(-1, yys3)
		} // end switch yys3
	} // end for yyj3
}

func (x *MixedcaseAddress) codecDecodeSelfFromArray(l int, d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	var yyj4 int
	var yyb4 bool
	var yyhl4 bool = l >= 0
	for {
		yyj4++
		if yyhl4 {
			yyb4 = yyj4 > l
		} else {
			yyb4 = z.DecCheckBreak()
		}
		if yyb4 {
			break
		}
		z.DecReadArrayElem()
		z.DecStructFieldNotFound(yyj4-1, "")
	}
}

func (x Hashes) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		h.encHashes((Hashes)(x), e)
	} // end block: if x slice == nil
}

func (x *Hashes) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	h.decHashes((*Hashes)(x), d)
}

func (x StorageKeys) CodecEncodeSelf(e *codec1978.Encoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if x == nil {
		r.EncodeNil()
	} else {
		h.encStorageKeys((StorageKeys)(x), e)
	} // end block: if x slice == nil
}

func (x *StorageKeys) CodecDecodeSelf(d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	h.decStorageKeys((*StorageKeys)(x), d)
}

func (x codecSelfer1) encHash(v *Hash, e *codec1978.Encoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if v == nil {
		r.EncodeNil()
		return
	}
	r.EncodeStringBytesRaw(((*[32]byte)(v))[:])
}

func (x codecSelfer1) decHash(v *Hash, d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	r.DecodeBytes(((*[32]byte)(v))[:], true)
}

func (x codecSelfer1) encUnprefixedHash(v *UnprefixedHash, e *codec1978.Encoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if v == nil {
		r.EncodeNil()
		return
	}
	r.EncodeStringBytesRaw(((*[32]byte)(v))[:])
}

func (x codecSelfer1) decUnprefixedHash(v *UnprefixedHash, d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	r.DecodeBytes(((*[32]byte)(v))[:], true)
}

func (x codecSelfer1) encAddress(v *Address, e *codec1978.Encoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if v == nil {
		r.EncodeNil()
		return
	}
	r.EncodeStringBytesRaw(((*[20]byte)(v))[:])
}

func (x codecSelfer1) decAddress(v *Address, d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	r.DecodeBytes(((*[20]byte)(v))[:], true)
}

func (x codecSelfer1) encUnprefixedAddress(v *UnprefixedAddress, e *codec1978.Encoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if v == nil {
		r.EncodeNil()
		return
	}
	r.EncodeStringBytesRaw(((*[20]byte)(v))[:])
}

func (x codecSelfer1) decUnprefixedAddress(v *UnprefixedAddress, d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	r.DecodeBytes(((*[20]byte)(v))[:], true)
}

func (x codecSelfer1) encHashes(v Hashes, e *codec1978.Encoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if v == nil {
		r.EncodeNil()
		return
	}
	z.EncWriteArrayStart(len(v))
	for _, yyv1 := range v {
		z.EncWriteArrayElem()
		yy2 := &yyv1
		yy2.CodecEncodeSelf(e)
	}
	z.EncWriteArrayEnd()
}

func (x codecSelfer1) decHashes(v *Hashes, d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyh1, yyl1 := z.DecSliceHelperStart()
	var yyc1 bool
	_ = yyc1
	if yyh1.IsNil {
		if yyv1 != nil {
			yyv1 = nil
			yyc1 = true
		}
	} else if yyl1 == 0 {
		if yyv1 == nil {
			yyv1 = []Hash{}
			yyc1 = true
		} else if len(yyv1) != 0 {
			yyv1 = yyv1[:0]
			yyc1 = true
		}
	} else {
		yyhl1 := yyl1 > 0
		var yyrl1 int
		_ = yyrl1
		if yyhl1 {
			if yyl1 > cap(yyv1) {
				yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 32)
				if yyrl1 <= cap(yyv1) {
					yyv1 = yyv1[:yyrl1]
				} else {
					yyv1 = make([]Hash, yyrl1)
				}
				yyc1 = true
			} else if yyl1 != len(yyv1) {
				yyv1 = yyv1[:yyl1]
				yyc1 = true
			}
		}
		var yyj1 int
		for yyj1 = 0; (yyhl1 && yyj1 < yyl1) || !(yyhl1 || z.DecCheckBreak()); yyj1++ { // bounds-check-elimination
			if yyj1 == 0 && yyv1 == nil {
				if yyhl1 {
					yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 32)
				} else {
					yyrl1 = 8
				}
				yyv1 = make([]Hash, yyrl1)
				yyc1 = true
			}
			yyh1.ElemContainerState(yyj1)
			var yydb1 bool
			if yyj1 >= len(yyv1) {
				yyv1 = append(yyv1, Hash{})
				yyc1 = true
			}
			if yydb1 {
				z.DecSwallow()
			} else {
				yyv1[yyj1].CodecDecodeSelf(d)
			}
		}
		if yyj1 < len(yyv1) {
			yyv1 = yyv1[:yyj1]
			yyc1 = true
		} else if yyj1 == 0 && yyv1 == nil {
			yyv1 = make([]Hash, 0)
			yyc1 = true
		}
	}
	yyh1.End()
	if yyc1 {
		*v = yyv1
	}
}

func (x codecSelfer1) encStorageKeys(v StorageKeys, e *codec1978.Encoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if v == nil {
		r.EncodeNil()
		return
	}
	z.EncWriteArrayStart(len(v))
	for _, yyv1 := range v {
		z.EncWriteArrayElem()
		yy2 := &yyv1
		h.encStorageKey((*StorageKey)(yy2), e)
	}
	z.EncWriteArrayEnd()
}

func (x codecSelfer1) decStorageKeys(v *StorageKeys, d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r

	yyv1 := *v
	yyh1, yyl1 := z.DecSliceHelperStart()
	var yyc1 bool
	_ = yyc1
	if yyh1.IsNil {
		if yyv1 != nil {
			yyv1 = nil
			yyc1 = true
		}
	} else if yyl1 == 0 {
		if yyv1 == nil {
			yyv1 = []StorageKey{}
			yyc1 = true
		} else if len(yyv1) != 0 {
			yyv1 = yyv1[:0]
			yyc1 = true
		}
	} else {
		yyhl1 := yyl1 > 0
		var yyrl1 int
		_ = yyrl1
		if yyhl1 {
			if yyl1 > cap(yyv1) {
				yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 72)
				if yyrl1 <= cap(yyv1) {
					yyv1 = yyv1[:yyrl1]
				} else {
					yyv1 = make([]StorageKey, yyrl1)
				}
				yyc1 = true
			} else if yyl1 != len(yyv1) {
				yyv1 = yyv1[:yyl1]
				yyc1 = true
			}
		}
		var yyj1 int
		for yyj1 = 0; (yyhl1 && yyj1 < yyl1) || !(yyhl1 || z.DecCheckBreak()); yyj1++ { // bounds-check-elimination
			if yyj1 == 0 && yyv1 == nil {
				if yyhl1 {
					yyrl1 = z.DecInferLen(yyl1, z.DecBasicHandle().MaxInitLen, 72)
				} else {
					yyrl1 = 8
				}
				yyv1 = make([]StorageKey, yyrl1)
				yyc1 = true
			}
			yyh1.ElemContainerState(yyj1)
			var yydb1 bool
			if yyj1 >= len(yyv1) {
				yyv1 = append(yyv1, StorageKey{})
				yyc1 = true
			}
			if yydb1 {
				z.DecSwallow()
			} else {
				h.decStorageKey((*StorageKey)(&yyv1[yyj1]), d)
			}
		}
		if yyj1 < len(yyv1) {
			yyv1 = yyv1[:yyj1]
			yyc1 = true
		} else if yyj1 == 0 && yyv1 == nil {
			yyv1 = make([]StorageKey, 0)
			yyc1 = true
		}
	}
	yyh1.End()
	if yyc1 {
		*v = yyv1
	}
}

func (x codecSelfer1) encStorageKey(v *StorageKey, e *codec1978.Encoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperEncoder(e)
	_, _, _ = h, z, r
	if v == nil {
		r.EncodeNil()
		return
	}
	r.EncodeStringBytesRaw(((*[72]byte)(v))[:])
}

func (x codecSelfer1) decStorageKey(v *StorageKey, d *codec1978.Decoder) {
	var h codecSelfer1
	z, r := codec1978.GenHelperDecoder(d)
	_, _, _ = h, z, r
	r.DecodeBytes(((*[72]byte)(v))[:], true)
}
