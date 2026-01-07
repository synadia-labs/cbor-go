package cbor

import "math"

const (
	byteValueCount = math.MaxUint8 + 1

	float16ExpBits  = 5
	float16MantBits = 10

	float32ExpBits  = 8
	float32MantBits = 23

	float16SignShift        = float16ExpBits + float16MantBits
	float16ExpShift         = float16MantBits
	float16ExpMask   uint16 = math.MaxUint16 >> (16 - float16ExpBits)
	float16MantMask  uint16 = math.MaxUint16 >> (16 - float16MantBits)
	float16ExpBias          = int(float16ExpMask >> 1)

	float32SignShift        = float32ExpBits + float32MantBits
	float32ExpShift         = float32MantBits
	float32ExpMask   uint32 = math.MaxUint8
	float32MantMask  uint32 = math.MaxUint32 >> (32 - float32MantBits)
	float32ExpBias          = int(float32ExpMask >> 1)
	float32HiddenBit uint32 = float32MantMask + 1

	float32ToFloat16MantShift  = float32MantBits - float16MantBits
	float32ToFloat16RoundShift = float32ToFloat16MantShift - 1
)
