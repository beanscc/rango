// Code generated by "github.com/beanscc/rango/cmd/genericgenerator -method unique -etype int8,int16,int,int32,int64,uint8,uint16,uint,uint32,uint64,float32,float64,string"; DO NOT EDIT.

package sliceutil

// Unique 切片去重
func Unique(slice interface{}) interface{} {
	switch slice.(type) {
	case []int8:
		return UniqueInt8s(slice.([]int8))
	case []int16:
		return UniqueInt16s(slice.([]int16))
	case []int:
		return UniqueInts(slice.([]int))
	case []int32:
		return UniqueInt32s(slice.([]int32))
	case []int64:
		return UniqueInt64s(slice.([]int64))
	case []uint8:
		return UniqueUint8s(slice.([]uint8))
	case []uint16:
		return UniqueUint16s(slice.([]uint16))
	case []uint:
		return UniqueUints(slice.([]uint))
	case []uint32:
		return UniqueUint32s(slice.([]uint32))
	case []uint64:
		return UniqueUint64s(slice.([]uint64))
	case []float32:
		return UniqueFloat32s(slice.([]float32))
	case []float64:
		return UniqueFloat64s(slice.([]float64))
	case []string:
		return UniqueStrings(slice.([]string))
	default:
		return UniqueReflect(slice, nil)
	}
}

// UniqueInt8s 切片去重
func UniqueInt8s(slice []int8) []int8 {
	l := len(slice)
	if l < 2 {
		return slice
	}

	resp := make([]int8, 0, l)
	m := make(map[int8]bool, l)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			resp = append(resp, v)
			m[v] = true
		}
	}

	return resp
}

// UniqueInt16s 切片去重
func UniqueInt16s(slice []int16) []int16 {
	l := len(slice)
	if l < 2 {
		return slice
	}

	resp := make([]int16, 0, l)
	m := make(map[int16]bool, l)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			resp = append(resp, v)
			m[v] = true
		}
	}

	return resp
}

// UniqueInts 切片去重
func UniqueInts(slice []int) []int {
	l := len(slice)
	if l < 2 {
		return slice
	}

	resp := make([]int, 0, l)
	m := make(map[int]bool, l)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			resp = append(resp, v)
			m[v] = true
		}
	}

	return resp
}

// UniqueInt32s 切片去重
func UniqueInt32s(slice []int32) []int32 {
	l := len(slice)
	if l < 2 {
		return slice
	}

	resp := make([]int32, 0, l)
	m := make(map[int32]bool, l)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			resp = append(resp, v)
			m[v] = true
		}
	}

	return resp
}

// UniqueInt64s 切片去重
func UniqueInt64s(slice []int64) []int64 {
	l := len(slice)
	if l < 2 {
		return slice
	}

	resp := make([]int64, 0, l)
	m := make(map[int64]bool, l)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			resp = append(resp, v)
			m[v] = true
		}
	}

	return resp
}

// UniqueUint8s 切片去重
func UniqueUint8s(slice []uint8) []uint8 {
	l := len(slice)
	if l < 2 {
		return slice
	}

	resp := make([]uint8, 0, l)
	m := make(map[uint8]bool, l)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			resp = append(resp, v)
			m[v] = true
		}
	}

	return resp
}

// UniqueUint16s 切片去重
func UniqueUint16s(slice []uint16) []uint16 {
	l := len(slice)
	if l < 2 {
		return slice
	}

	resp := make([]uint16, 0, l)
	m := make(map[uint16]bool, l)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			resp = append(resp, v)
			m[v] = true
		}
	}

	return resp
}

// UniqueUints 切片去重
func UniqueUints(slice []uint) []uint {
	l := len(slice)
	if l < 2 {
		return slice
	}

	resp := make([]uint, 0, l)
	m := make(map[uint]bool, l)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			resp = append(resp, v)
			m[v] = true
		}
	}

	return resp
}

// UniqueUint32s 切片去重
func UniqueUint32s(slice []uint32) []uint32 {
	l := len(slice)
	if l < 2 {
		return slice
	}

	resp := make([]uint32, 0, l)
	m := make(map[uint32]bool, l)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			resp = append(resp, v)
			m[v] = true
		}
	}

	return resp
}

// UniqueUint64s 切片去重
func UniqueUint64s(slice []uint64) []uint64 {
	l := len(slice)
	if l < 2 {
		return slice
	}

	resp := make([]uint64, 0, l)
	m := make(map[uint64]bool, l)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			resp = append(resp, v)
			m[v] = true
		}
	}

	return resp
}

// UniqueFloat32s 切片去重
func UniqueFloat32s(slice []float32) []float32 {
	l := len(slice)
	if l < 2 {
		return slice
	}

	resp := make([]float32, 0, l)
	m := make(map[float32]bool, l)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			resp = append(resp, v)
			m[v] = true
		}
	}

	return resp
}

// UniqueFloat64s 切片去重
func UniqueFloat64s(slice []float64) []float64 {
	l := len(slice)
	if l < 2 {
		return slice
	}

	resp := make([]float64, 0, l)
	m := make(map[float64]bool, l)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			resp = append(resp, v)
			m[v] = true
		}
	}

	return resp
}

// UniqueStrings 切片去重
func UniqueStrings(slice []string) []string {
	l := len(slice)
	if l < 2 {
		return slice
	}

	resp := make([]string, 0, l)
	m := make(map[string]bool, l)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			resp = append(resp, v)
			m[v] = true
		}
	}

	return resp
}