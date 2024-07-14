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
				err:     New(CodeUnknown, errors.New("not ecode.Error").Error()),
			},
		},
		{
			name: "test ecode.Error",
			err: func() error {
				return New(400, "invalid params")
			},
			fromError: fromError{
				isError: true,
				err:     New(400, "invalid params"),
			},
		},
		{
			name: "test ecode.Error with detailed err 2",
			err: func() error {
				return New(400, "invalid params").WithDetails(errors.New("detail err"))
			},
			fromError: fromError{
				isError: true,
				err:     New(400, "invalid params").WithDetails(errors.New("detail err")),
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

func TestError(t *testing.T) {
	tests := []struct {
		err *Error
		str string
	}{
		{nil, `<nil>`},
		{New(400, "bad request"), `error code: 400, message: bad request`},
		{New(500, "system error").WithDetails(errors.New("can't connect to redis")), `error code: 500, message: system error, details: can't connect to redis`},
	}

	for i, tt := range tests {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Logf(`e.Error():%v`, tt.err)
			if tt.err != nil {
				if tt.err.Error() != tt.str {
					t.Errorf("ecode.Error string failed. got:%s; want:%s", tt.err, tt.str)
				}
			}
		})
	}
}
