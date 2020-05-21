package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var (
	connDSN    = flag.String("dsn", "", "mysql connect dsn; eg: user:password@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local")
	outputPath = flag.String("output", "scheme", "output path, the last base path is the package name; eg: ${project path}/repo/scheme")
)

func Usage() {
	fmt.Fprintf(os.Stderr, "Usage of autoschme:\n")
	fmt.Fprintf(os.Stderr, "\t autochme -dsn $db_dsn -output $your/path/packagename\n")
	flag.PrintDefaults()
}

func main() {
	flag.Usage = Usage
	flag.Parse()

	if *connDSN == "" {
		flag.Usage()
		os.Exit(2)
	}

	start := time.Now()
	log.Println("start to connect to database ...")

	// 连接 mysql
	if err := initDB(*connDSN); err != nil {
		panic(err)
	}
	connTime := time.Now().Sub(start)
	log.Printf("connect to database takes: %s", connTime)

	tables, err := getTables()
	if err != nil {
		log.Fatalf("get tables failed. err: %v", err)
	}

	dir, err := filepath.Abs(*outputPath)
	if err != nil {
		log.Fatalf("get output abs path failed. err: %v", err)
	}
	log.Printf("output dir: %v", dir)

	pkgName := filepath.Base(dir)
	log.Printf("output package: %v", pkgName)

	// 检查目录
	_, err = os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			if err := os.MkdirAll(dir, 0755); err != nil {
				log.Fatalf("mkdir %s failed. err: %v", dir, err)
			}
		} else {
			log.Fatalf("check dir failed. err: %v", err)
		}
	}

	for _, v := range tables {
		b, err := buildTableStruct(pkgName, v)
		if err != nil {
			log.Fatalf("gen table scheme failed. err: %v", err)
		}

		// write to file
		file := filepath.Join(dir, v+".go")
		if err := ioutil.WriteFile(file, []byte(b), 0644); err != nil {
			log.Fatalf("gen table scheme file failed. err: %v", err)
		}
		log.Printf("created file: %s", file)
	}

	end := time.Now()
	log.Printf("total takes %s, conn takes: %s", end.Sub(start), connTime)
}
