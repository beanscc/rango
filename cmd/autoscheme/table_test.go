package main

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	dsn := `root:P4m@bpet@tcp(127.0.0.1:3306)/crawler?charset=utf8mb4&parseTime=True&loc=Local`
	if err := initDB(dsn); err != nil {
		panic(err)
	}
	m.Run()
}

func Test_getTables(t *testing.T) {
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

func Test_buildTableStructBuffer(t *testing.T) {
	got, err := buildTableStructBuffer("model", "crawler_aliexpress_category")
	if err != nil {
		t.Errorf("Test_buildTableStructBuffer failed. err:%v", err)
		return
	}
	fmt.Println(got)
}
