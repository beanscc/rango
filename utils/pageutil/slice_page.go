package pageutil

import "fmt"

// pageWithSliceData Page slice data
type pageWithSliceData struct {
	size  int // 查询每页记录数
	count int // 总记录数
}

// NewPageWithSliceData new page with slice data
// size : page size
// count: data len
func NewPageWithSliceData(size, count int) *pageWithSliceData { // nolint: golint
	if size < 1 {
		panic(fmt.Sprintf("wrong page size: %v", size))
	}

	if count < 0 {
		panic(fmt.Sprintf("wrong data count: %v", count))
	}

	return &pageWithSliceData{size: size, count: count}
}

// TotalPageNum 获取总页数
func (p *pageWithSliceData) TotalPageNum() int {
	t := p.count / p.size
	if (p.count % p.size) > 0 {
		t++
	}

	return t
}

// PageDataOffset 根据 index 分页获取数据
func (p *pageWithSliceData) PageDataOffset(index int) (int, int, bool) {
	if index < 1 {
		panic(fmt.Sprintf("wrong page index: %v", index))
	}

	start := (index - 1) * p.size
	end := start + p.size

	// 请求页无数据
	if p.count <= start {
		return 0, 0, false
	}

	if p.count > start && p.count < end {
		return start, p.count, true
	}

	// 数据记录总数>=分页请求记录数
	return start, end, true
}

// GetPageParam 根据参数合理设置分页参数
func GetPageParam(index, size, defaultSize, defaultMaxSize int) (int, int) {
	if index < 1 {
		index = 1
	}

	if size < 1 {
		size = defaultSize
	}

	if size > defaultMaxSize {
		size = defaultMaxSize
	}

	return index, size
}
