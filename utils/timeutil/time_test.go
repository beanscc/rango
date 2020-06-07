package timeutil

import (
	"reflect"
	"testing"
	"time"
)

func Test_DayRange(t *testing.T) {
	td := time.Date(2018, 01, 12, 15, 30, 0, 0, time.Local)

	dsWant := time.Date(2018, 01, 12, 0, 0, 0, 0, time.Local)
	deWant := dsWant.AddDate(0, 0, 1).Add(-1 * time.Second)

	ds, de := DayRange(td, time.Local)
	if dsWant != ds {
		t.Errorf("DayRange() ds failed. want: %v, got: %v", dsWant, ds)
	}

	if deWant != de {
		t.Errorf("DayRange() de failed. want: %v, got: %v", deWant, de)
	}
}

func TestRange(t *testing.T) {
	type args struct {
		start time.Time
		end   time.Time
		step  time.Duration
	}

	tests := []struct {
		name string
		args args
		want []time.Time
	}{
		{
			"t1",
			args{
				start: time.Date(2019, 07, 04, 0, 0, 0, 0, time.Local),
				end:   time.Date(2019, 07, 04, 0, 0, 0, 0, time.Local),
				step:  time.Hour,
			},
			[]time.Time{
				time.Date(2019, 07, 04, 0, 0, 0, 0, time.Local),
			},
		},
		{
			"t2",
			args{
				start: time.Date(2019, 07, 04, 0, 0, 0, 0, time.Local),
				end:   time.Date(2019, 07, 04, 1, 0, 0, 0, time.Local),
				step:  time.Hour,
			},
			[]time.Time{
				time.Date(2019, 07, 04, 0, 0, 0, 0, time.Local),
				time.Date(2019, 07, 04, 1, 0, 0, 0, time.Local),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Range(tt.args.start, tt.args.end, tt.args.step)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Range()=%+v, want=%+v", got, tt.want)
			}
		})
	}
}

func TestRangeN(t *testing.T) {
	type args struct {
		start time.Time
		step  time.Duration
		n     int
	}

	tests := []struct {
		name string
		args args
		want []time.Time
	}{
		{
			"t1",
			args{
				start: time.Date(2019, 07, 04, 0, 0, 0, 0, time.Local),
				step:  time.Hour,
				n:     2,
			},
			[]time.Time{
				time.Date(2019, 07, 04, 0, 0, 0, 0, time.Local),
				time.Date(2019, 07, 04, 1, 0, 0, 0, time.Local),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RangeN(tt.args.start, tt.args.step, tt.args.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RangeN()=%+v, want=%+v", got, tt.want)
			}
		})
	}
}
