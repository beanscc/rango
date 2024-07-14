package numberutil

// SignedInteger 有符号整型
type SignedInteger interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

// UnsignedInteger 无符号整型
type UnsignedInteger interface {
	~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

// Integer 整型
type Integer interface {
	SignedInteger | UnsignedInteger
}

// Float 浮点型
type Float interface {
	~float32 | ~float64
}
