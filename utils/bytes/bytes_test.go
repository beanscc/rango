package bytes

import "testing"

func Test_countPrefix(t *testing.T) {
	type args struct {
		s   []byte
		sep []byte
	}

	type want struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{"t1", args{s: []byte("# a"), sep: []byte("#")}, want{n: 1}},
		{"t2", args{s: []byte(" # a"), sep: []byte("#")}, want{n: 0}},
		{"t3", args{s: []byte("## ab"), sep: []byte("#")}, want{n: 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			args := tt.args
			got := CountPrefix(args.s, args.sep)
			if got != tt.want.n {
				t.Errorf("CountPrefix() failed. want: %v, got: %v", tt.want.n, got)
			}
		})
	}
}
