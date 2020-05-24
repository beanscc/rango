package ecode

import (
	"errors"
	"fmt"
	"reflect"
	"testing"
)

func TestFromError(t *testing.T) {
	type fromError struct {
		isError bool
		err     *Error
	}

	tests := []struct {
		name      string
		err       func() error
		fromError fromError
	}{
		{
			name: "test nil error",
			err: func() error {
				return nil
			},
			fromError: fromError{
				isError: true,
				err:     nil,
			},
		},
		{
			name: "test std err",
			err: func() error {
				return errors.New("not ecode.Error")
			},
			fromError: fromError{
				isError: false,
				err:     nil,
			},
		},
		{
			name: "test ecode.Error",
			err: func() error {
				return NewError(400, "invalid params")
			},
			fromError: fromError{
				isError: true,
				err:     NewError(400, "invalid params"),
			},
		},
		{
			name: "test ecode.Error with detailed err 2",
			err: func() error {
				return NewErrorWithDetail(400, "invalid params", errors.New("detail err"))
			},
			fromError: fromError{
				isError: true,
				err:     NewErrorWithDetail(400, "invalid params", errors.New("detail err")),
				// err: NewError(400, "invalid params"),
			},
		},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprintf("i:%d, name:%s", i, tt.name), func(t *testing.T) {
			e, ok := FromError(tt.err())
			if ok != tt.fromError.isError {
				t.Errorf("TestFromError failed. gotOk:%v, wantOk:%v", ok, tt.fromError.isError)
				return
			}

			if !reflect.DeepEqual(e, tt.fromError.err) {
				t.Errorf("TestFromError failed. err not equal. gotErr:%#v, wantErr:%#v", e, tt.fromError.err)
			}
		})
	}
}

func TestError_String(t *testing.T) {
	tests := []struct {
		err *Error
		str string
	}{
		{nil, `<nil>`},
		{NewError(400, "bad request"), `error code: 400, msg: bad request`},
		{NewErrorWithDetail(500, "system error", errors.New("can't connect to redis")), `error code: 500, msg: system error`},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Logf(`
e.Error() :%v
e.String():%v
`, tt.err, tt.err.String())

			if tt.err != nil {
				if tt.err.Error() != tt.str {
					t.Errorf("ecode.Error string failed. got:%s; want:%s", tt.err, tt.str)
				}
			}
		})
	}
}
