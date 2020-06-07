package main

import (
	"testing"
)

// func TestMain(m *testing.M) {
// 	dsn := `root:P4m@bpet@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local`
// 	if err := initDB(dsn); err != nil {
// 		panic(err)
// 	}
// 	m.Run()
// }

func testInit() {
	dsn := `root:P4m@bpet@tcp(127.0.0.1:3306)/test?charset=utf8mb4&parseTime=True&loc=Local`
	if err := initDB(dsn); err != nil {
		panic(err)
	}
}

func Test_getTables(t *testing.T) {
	testInit()
	got, err := getTables()
	if err != nil {
		t.Errorf("Test_getTables failed. err:%v", err)
		return
	}
	for i, v := range got {
		t.Logf("Test_getTables got i:%v, table:%v", i, v)
	}
}

func Test_getTableFullColumns(t *testing.T) {
	testInit()
	tables, err := getTables()
	if err != nil {
		t.Errorf("Test_getTableFullColumns failed. get tables err:%v", err)
		return
	}
	for _, v := range tables {
		columns, err := getTableFullColumns(v)
		if err != nil {
			t.Errorf("Test_getTableFullColumns failed. get table column err:%v", err)
		}

		for i, vv := range columns {
			t.Logf("Test_getTableFullColumns table:%v, column i:%v, column:%+v", v, i, vv)
		}
	}
}

func Test_buildTableStruct(t *testing.T) {
	testInit()
	got, err := buildTableStruct("model", "user")
	if err != nil {
		t.Errorf("Test_buildTableStruct failed. err:%v", err)
		return
	}
	t.Logf("Test_buildTableStruct got:\n%s", got)
}
