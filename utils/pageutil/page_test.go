package pageutil

import (
	"reflect"
	"testing"
)

func TestPage(t *testing.T) {
	tests := []struct {
		name          string
		page          *Page
		defaultSize   int
		wantErr       bool
		wantValidPage *Page
		wantOffset    int
		remark        string
	}{
		{name: "t1", page: New(0, 0), defaultSize: 10, wantErr: false, wantValidPage: New(1, 10), wantOffset: 0, remark: "零值测试"},
		{name: "t2", page: New(-1, 0), defaultSize: 10, wantErr: true, wantValidPage: New(1, 10), wantOffset: 0, remark: "无效的页码测试"},
		{name: "t3", page: New(0, -1), defaultSize: 10, wantErr: true, wantValidPage: New(1, 10), wantOffset: 0, remark: "无效的分页大小测试"},
		{name: "t4", page: New(-1, -1), defaultSize: 10, wantErr: true, wantValidPage: New(1, 10), wantOffset: 0, remark: "无效的页码和分页大小测试"},
		{name: "t5", page: New(2, 5), defaultSize: 10, wantErr: false, wantValidPage: New(2, 5), wantOffset: 5, remark: "正常的页码和分页大小测试"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.page.Check(); err != nil && !tt.wantErr {
				t.Errorf("TestPage Check failed. err:%v", err)
			}

			tt.page.SetValidVal(tt.defaultSize)
			if !reflect.DeepEqual(tt.page, tt.wantValidPage) {
				t.Errorf("TestPage SetValidVal failed. got:%+v, want:%+v", tt.page, tt.wantValidPage)
				return
			}
			offset := tt.page.Offset()
			if tt.wantOffset != offset {
				t.Errorf("TestPage Offset() failed. got:%v, want:%v", offset, tt.wantOffset)
				return
			}
		})
	}
}

func TestSlicePage(t *testing.T) {
	type args struct {
		total     int
		pageIndex int
		pageSize  int
	}
	type want struct {
		start int
		end   int
		exist bool
	}
	tests := []struct {
		name   string
		remark string
		args   args
		want   want
	}{
		{name: "t1", remark: "测试分页数据全包含在切片数据中", args: args{total: 10, pageIndex: 1, pageSize: 5}, want: want{start: 0, end: 5, exist: true}},
		{name: "t2", remark: "测试分页数据半包含在切片数据中", args: args{total: 10, pageIndex: 2, pageSize: 6}, want: want{start: 6, end: 10, exist: true}},
		{name: "t3", remark: "测试分页数据不包含在切片数据中", args: args{total: 10, pageIndex: 3, pageSize: 5}, want: want{start: 0, end: 0, exist: false}},
		{name: "t4", remark: "测试分页数据不包含在切片数据中", args: args{total: 10, pageIndex: 3, pageSize: 6}, want: want{start: 0, end: 0, exist: false}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			start, end, exist := SlicePage(tt.args.total, tt.args.pageIndex, tt.args.pageSize)
			if exist != tt.want.exist ||
				start != tt.want.start ||
				end != tt.want.end {
				t.Errorf("TestSlicePage failed. got(start:%v, end:%v, exist:%v), want(start:%v, end:%v, exist:%v)", start, end, exist, tt.want.start, tt.want.end, tt.want.exist)
			}
		})
	}
}
