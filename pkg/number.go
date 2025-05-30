package pkg

type Number interface {
	UnsignedNumber | SignedNumber | FloatNumber
}

type UnsignedNumber interface {
	uint | uint8 | uint16 | uint32 | uint64 | uintptr
}

type SignedNumber interface {
	int | int8 | int16 | int32 | int64
}

type FloatNumber interface {
	float32 | float64
}
