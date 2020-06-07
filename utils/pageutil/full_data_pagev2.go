package pageutil

// PageWithFullData Page
type PageWithFullData struct {
	size           int // 查询每页记录数
	count          int // 总记录数
	defaultMinSize int // 默认查询每页最小记录数
	defaultMaxSize int // 默认查询每页最大记录数
}

const (
	defaultMinSize = 10
	defaultMaxSize = 50
)

func New(size, count int) *PageWithFullData {
	p := &PageWithFullData{
		count:          count,
		defaultMinSize: defaultMinSize,
		defaultMaxSize: defaultMaxSize,
	}

	// 根据设置的minSize 和 maxSize 重新调整 size
	p.SetPageSize(size)

	return p
}

// SetDefaultMinPageSize 设置每页最小查询记录数
func (p *PageWithFullData) SetDefaultMinPageSize(v int) {
	if v > 0 {
		p.defaultMinSize = v
	}
}

// SetDefaultMaxPageSize 设置每页最大查询记录数
func (p *PageWithFullData) SetDefaultMaxPageSize(v int) {
	if v > 0 && v > p.defaultMinSize {
		p.defaultMaxSize = v
	}
}

// DefaultMinPageSize DefaultMinPageSize
func (p *PageWithFullData) DefaultMinPageSize() int {
	return p.defaultMinSize
}

// DefaultMaxPageSize DefaultMaxPageSize
func (p *PageWithFullData) DefaultMaxPageSize() int {
	return p.defaultMaxSize
}

// TotalPage 获取总页数
func (p *PageWithFullData) TotalPage() int {
	t := p.count / p.size
	if (p.count % p.size) > 0 {
		t++
	}

	return t
}

func (p *PageWithFullData) initPageIndex(index int) int {
	if index < 1 {
		index = 1
	}

	return index
}

// SetPageSize 根据设置的minSize 和 maxSize 重新调整 size
func (p *PageWithFullData) SetPageSize(size int) {
	if size < p.defaultMinSize {
		size = p.defaultMinSize
	}

	if size > p.defaultMaxSize {
		size = p.defaultMaxSize
	}

	p.size = size
}

// PageDataOffset 根据
func (p *PageWithFullData) PageDataOffset(index int) (int, int, bool) {
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
