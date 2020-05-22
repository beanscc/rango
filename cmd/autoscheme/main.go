package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/beanscc/rango/utils"
	_ "github.com/go-sql-driver/mysql"
)

var (
	version      = flag.Bool("v", false, "version prints the build information")
	help         = flag.Bool("h", false, "help prints the usage")
	connDSN      = flag.String("dsn", "", "mysql connect dsn; eg: user:password@tcp(127.0.0.1:3306)/database?charset=utf8mb4&parseTime=True&loc=Local")
	outputPath   = flag.String("output", "scheme", "output path, the last base path is the package name; eg: ${project path}/repo/scheme")
	wantTables   = flag.String("tables", "", "tables that you want to generate, separate multiple table names with commas; eg: table1,table2")
	prefixTables = flag.String("prefix", "", "若指定该选项，则只处理表名中含有此前缀的表结构")
	suffixTables = flag.String("suffix", "", "若指定该选项，则只处理表名中含有此后缀的表结构")
)

func Usage() {
	fmt.Fprintf(os.Stderr, `用于自动生成 mysql 表结构对应 go struct 的小工具

Usage:
  autoscheme -dsn $dsn [other options]

available options:
`)
	flag.PrintDefaults()
}

func handleFlag() {
	flag.Usage = Usage
	flag.Parse()
	if *version {
		fmt.Println(utils.Version())
		os.Exit(0)
	}

	if *help {
		flag.Usage()
		os.Exit(0)
	}
}

func main() {
	handleFlag()

	if *connDSN == "" {
		flag.Usage()
		os.Exit(2)
	}

	start := time.Now()
	log.Println("start to connect to database ...")

	var err error
	// 连接 mysql
	if err := initDB(*connDSN); err != nil {
		panic(err)
	}
	connTime := time.Now().Sub(start)
	log.Printf("connect to database takes: %s", connTime)

	tables, err := getValidTables()
	if err != nil {
		log.Fatalf("get valid tables failed. err: %v", err)
	}

	if len(tables) == 0 {
		log.Printf("no tables to build scheme")
		os.Exit(0)
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

func getValidTables() ([]string, error) {
	var (
		err    error
		tables []string
	)

	if *wantTables != "" {
		tables = strings.Split(*wantTables, ",")
	} else {
		tables, err = getTables()
		if err != nil {
			return nil, fmt.Errorf("get tables failed. err: %v", err)
		}
	}

	if *prefixTables != "" {
		hasPrefixTables := make([]string, 0, len(tables))
		for _, v := range tables {
			if strings.HasPrefix(v, *prefixTables) {
				hasPrefixTables = append(hasPrefixTables, v)
			}
		}

		tables = hasPrefixTables
	}

	if *suffixTables != "" {
		hasSuffixTables := make([]string, 0, len(tables))
		for _, v := range tables {
			if strings.HasSuffix(v, *suffixTables) {
				hasSuffixTables = append(hasSuffixTables, v)
			}
		}

		tables = hasSuffixTables
	}

	return tables, nil
}
