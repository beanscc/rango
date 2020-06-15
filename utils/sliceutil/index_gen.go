// Code generated by "github.com/beanscc/rango/cmd/genericgenerator -method index -etype int8,int16,int,int32,int64,uint8,uint16,uint,uint32,uint64,float32,float64,string"; DO NOT EDIT.

package sliceutil

// Index 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1
func Index(slice interface{}, x interface{}) interface{} {
	switch slice.(type) {
	case []int8:
		return IndexInt8s(slice.([]int8), x.(int8))
	case []int16:
		return IndexInt16s(slice.([]int16), x.(int16))
	case []int:
		return IndexInts(slice.([]int), x.(int))
	case []int32:
		return IndexInt32s(slice.([]int32), x.(int32))
	case []int64:
		return IndexInt64s(slice.([]int64), x.(int64))
	case []uint8:
		return IndexUint8s(slice.([]uint8), x.(uint8))
	case []uint16:
		return IndexUint16s(slice.([]uint16), x.(uint16))
	case []uint:
		return IndexUints(slice.([]uint), x.(uint))
	case []uint32:
		return IndexUint32s(slice.([]uint32), x.(uint32))
	case []uint64:
		return IndexUint64s(slice.([]uint64), x.(uint64))
	case []float32:
		return IndexFloat32s(slice.([]float32), x.(float32))
	case []float64:
		return IndexFloat64s(slice.([]float64), x.(float64))
	case []string:
		return IndexStrings(slice.([]string), x.(string))
	default:
		return indexReflect(slice, x)
	}
}

// IndexInt8s 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1
func IndexInt8s(slice []int8, x int8) int {
	l := len(slice)
	for i := 0; i < l; i++ {
		if x == slice[i] {
			return i
		}
	}

	return -1
}

// IndexInt16s 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1
func IndexInt16s(slice []int16, x int16) int {
	l := len(slice)
	for i := 0; i < l; i++ {
		if x == slice[i] {
			return i
		}
	}

	return -1
}

// IndexInts 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1
func IndexInts(slice []int, x int) int {
	l := len(slice)
	for i := 0; i < l; i++ {
		if x == slice[i] {
			return i
		}
	}

	return -1
}

// IndexInt32s 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1
func IndexInt32s(slice []int32, x int32) int {
	l := len(slice)
	for i := 0; i < l; i++ {
		if x == slice[i] {
			return i
		}
	}

	return -1
}

// IndexInt64s 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1
func IndexInt64s(slice []int64, x int64) int {
	l := len(slice)
	for i := 0; i < l; i++ {
		if x == slice[i] {
			return i
		}
	}

	return -1
}

// IndexUint8s 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1
func IndexUint8s(slice []uint8, x uint8) int {
	l := len(slice)
	for i := 0; i < l; i++ {
		if x == slice[i] {
			return i
		}
	}

	return -1
}

// IndexUint16s 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1
func IndexUint16s(slice []uint16, x uint16) int {
	l := len(slice)
	for i := 0; i < l; i++ {
		if x == slice[i] {
			return i
		}
	}

	return -1
}

// IndexUints 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1
func IndexUints(slice []uint, x uint) int {
	l := len(slice)
	for i := 0; i < l; i++ {
		if x == slice[i] {
			return i
		}
	}

	return -1
}

// IndexUint32s 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1
func IndexUint32s(slice []uint32, x uint32) int {
	l := len(slice)
	for i := 0; i < l; i++ {
		if x == slice[i] {
			return i
		}
	}

	return -1
}

// IndexUint64s 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1
func IndexUint64s(slice []uint64, x uint64) int {
	l := len(slice)
	for i := 0; i < l; i++ {
		if x == slice[i] {
			return i
		}
	}

	return -1
}

// IndexFloat32s 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1
func IndexFloat32s(slice []float32, x float32) int {
	l := len(slice)
	for i := 0; i < l; i++ {
		if x == slice[i] {
			return i
		}
	}

	return -1
}

// IndexFloat64s 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1
func IndexFloat64s(slice []float64, x float64) int {
	l := len(slice)
	for i := 0; i < l; i++ {
		if x == slice[i] {
			return i
		}
	}

	return -1
}

// IndexStrings 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1
func IndexStrings(slice []string, x string) int {
	l := len(slice)
	for i := 0; i < l; i++ {
		if x == slice[i] {
			return i
		}
	}

	return -1
}