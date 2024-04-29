package pagination

const DefaultPageSize = 10

type Pagination struct {
	// 页码
	Index int `json:"page_index" form:"page_index"`
	// 每页记录数
	Size int `json:"page_size" form:"page_size"`
	// 总记录数
	TotalCount int `json:"total_count" form:"total_count"`
}

func New(index int, size int, defaultSize int) *Pagination {
	if index < 1 {
		index = 1
	}

	if size < 1 {
		if defaultSize < 1 {
			size = DefaultPageSize
		} else {
			size = defaultSize
		}
	}
	return &Pagination{
		Index: index,
		Size:  size,
	}
}

// Limit limit of data
func (p *Pagination) Limit() int {
	return p.Size
}

// Offset offset of data
func (p *Pagination) Offset() int {
	return (p.Index - 1) * p.Size
}

// SetTotalCount 设置总记录数
func (p *Pagination) SetTotalCount(total int) {
	p.TotalCount = total
}

// PageCount 总页数
func (p *Pagination) PageCount() int {
	if p.TotalCount <= p.Size {
		return 1
	}

	d := p.TotalCount / p.Size
	m := p.TotalCount % p.Size
	if m == 0 {
		return d
	}

	return d + 1
}
