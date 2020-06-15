package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/beanscc/rango/utils/stringutil"
	"golang.org/x/tools/go/packages"
)

var (
	method    = flag.String("method", "", "要生成的泛型方法类型; eg: filter/unique")
	eTypes    = flag.String("etype", "", "对应方法的类型; eg: int,int32")
	buildTags = flag.String("tags", "", "comma-separated list of build tags to apply")
)

func main() {
	flag.Parse()

	var tags []string
	if len(*buildTags) > 0 {
		tags = strings.Split(*buildTags, ",")
	}

	// We accept either one directory or a list of files. Which do we have?
	args := flag.Args()
	if len(args) == 0 {
		// Default: process whole package in current directory.
		args = []string{"."}
	}

	// Parse the package once.
	var dir string
	// TODO(suzmue): accept other patterns for packages (directories, list of files, import paths, etc).
	if len(args) == 1 && isDirectory(args[0]) {
		dir = args[0]
	} else {
		if len(tags) != 0 {
			log.Fatal("-tags option applies only to directories, not when files are specified")
		}
		dir = filepath.Dir(args[0])
	}

	pkg := parsePackage(args, tags)
	switch *method {
	case "filter":
		src, err := filterGenerator(strings.Split(*eTypes, ","), pkg.Name, filepath.Join(dir, "filter_gen.go"))
		if err != nil {
			log.Fatalf("filterGenerator failed. method:%s, etype:%s, err:%v; src:%s", *method, *eTypes, err, src)
		}
	case "chunk":
		src, err := chunkGenerator(strings.Split(*eTypes, ","), pkg.Name, filepath.Join(dir, "chunk_gen.go"))
		if err != nil {
			log.Fatalf("chunkGenerator failed. method:%s, etype:%s, err:%v; src:%s", *method, *eTypes, err, src)
		}
	case "index":
		src, err := indexGenerator(strings.Split(*eTypes, ","), pkg.Name, filepath.Join(dir, "index_gen.go"))
		if err != nil {
			log.Fatalf("indexGenerator failed. method:%s, etype:%s, err:%v; src:%s", *method, *eTypes, err, src)
		}
	case "unique":
		src, err := uniqueGenerator(strings.Split(*eTypes, ","), pkg.Name, filepath.Join(dir, "unique_gen.go"))
		if err != nil {
			log.Fatalf("uniqueGenerator failed. method:%s, etype:%s, err:%v; src:%s", *method, *eTypes, err, src)
		}
	case "join":
		src, err := joinGenerator(strings.Split(*eTypes, ","), pkg.Name, filepath.Join(dir, "join_gen.go"))
		if err != nil {
			log.Fatalf("joinGenerator failed. method:%s, etype:%s, err:%v; src:%s", *method, *eTypes, err, src)
		}
	default:
		log.Fatalf("unknown method:%s", *method)
	}
}

// isDirectory reports whether the named file is a directory.
func isDirectory(name string) bool {
	info, err := os.Stat(name)
	if err != nil {
		log.Fatal(err)
	}
	return info.IsDir()
}

// parsePackage analyzes the single package constructed from the patterns and tags.
// parsePackage exits if there is an error.
func parsePackage(patterns []string, tags []string) *packages.Package {
	cfg := &packages.Config{
		Mode: packages.LoadSyntax,
		// TODO: Need to think about constants in test files. Maybe write type_string_test.go
		// in a separate pass? For later.
		Tests:      false,
		BuildFlags: []string{fmt.Sprintf("-tags=%s", strings.Join(tags, " "))},
	}
	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		log.Fatal(err)
	}
	if len(pkgs) != 1 {
		log.Fatalf("error: %d packages found", len(pkgs))
	}
	return pkgs[0]
}

func formatAndWriteFile(data []byte, output string) ([]byte, error) {
	src, err := format.Source(data)
	if err != nil {
		// Should never happen, but can arise when developing this code.
		// The user can compile the output to see the error.
		log.Printf("warning: internal error: invalid Go generated: %s", err)
		log.Printf("warning: compile the package to analyze the error")
		return data, err
	}

	err = ioutil.WriteFile(output, src, 0644)
	if err != nil {
		return src, err
	}

	return src, nil
}

func filterGenerator(eTypes []string, pkg string, output string) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("// Code generated by \"github.com/beanscc/rango/cmd/genericgenerator %s\"; DO NOT EDIT.\n", strings.Join(os.Args[1:], " ")))
	b.WriteString(fmt.Sprintf("\npackage %s\n", pkg))
	imports := []string{
		// `"fmt"`,
	}
	if len(imports) > 0 {
		b.WriteString("\nimport (\n\t")
		b.WriteString(strings.Join(imports, "\n\t"))
		b.WriteString("\n)\n")
	}

	b.WriteString("\n// Filter filter() == true 时，保留该项，否则丢弃改项，返回一个新的该类型的切片")
	b.WriteString("\nfunc Filter(slice interface{}, filter func(i int) bool) interface{} {\n")
	b.WriteString("\tswitch slice.(type) {")
	for _, eType := range eTypes {
		b.WriteString(fmt.Sprintf("case []%s:\n", eType))
		b.WriteString(fmt.Sprintf("\treturn Filter%ss(slice.([]%s), filter)\n", stringutil.Snake2Camel(eType, true), eType))
	}
	b.WriteString("\tdefault:\n\t\treturn filterReflect(slice, filter)\n\t}\n}\n")

	for _, eType := range eTypes {
		fnName := fmt.Sprintf("Filter%ss", stringutil.Snake2Camel(eType, true))
		b.WriteString(fmt.Sprintf(`
// %s 过滤 []%s 切片，只保留 filter(i) == true 时的索引项，并返回一个新的该类型切片
func %s(slice []%s, filter func(i int) bool) []%s {
	l := len(slice)
	resp := make([]%s, 0, l)
	for i := 0; i < l; i++ {
		if filter(i) {
			resp = append(resp, slice[i])
		}
	}

	return resp
}
`, fnName, eType,
			fnName, eType, eType,
			eType))
	}

	return formatAndWriteFile(b.Bytes(), output)
}

func chunkGenerator(eTypes []string, pkg string, output string) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("// Code generated by \"github.com/beanscc/rango/cmd/genericgenerator %s\"; DO NOT EDIT.\n", strings.Join(os.Args[1:], " ")))
	b.WriteString(fmt.Sprintf("\npackage %s\n", pkg))
	imports := []string{
		// `"fmt"`,
	}
	if len(imports) > 0 {
		b.WriteString("\nimport (\n\t")
		b.WriteString(strings.Join(imports, "\n\t"))
		b.WriteString("\n)\n")
	}

	b.WriteString("\n// Chunk 将切片按 size 大小分块")
	b.WriteString("\nfunc Chunk(slice interface{}, size int) interface{} {\n")
	b.WriteString("\tswitch slice.(type) {")
	for _, eType := range eTypes {
		b.WriteString(fmt.Sprintf("case []%s:\n", eType))
		b.WriteString(fmt.Sprintf("\treturn Chunk%ss(slice.([]%s), size)\n", stringutil.Snake2Camel(eType, true), eType))
	}
	b.WriteString("\tdefault:\n\t\treturn chunkReflect(slice, size)\n\t}\n}\n")

	for _, eType := range eTypes {
		fnName := fmt.Sprintf("Chunk%ss", stringutil.Snake2Camel(eType, true))
		b.WriteString(fmt.Sprintf(`
// %s 将 []%s 切片按 size 大小分块，返回一个新的切片[][]%s
func %s(slice []%s, size int) [][]%s {
	l := len(slice)
	chunks := make([][]%s, 0, chunkCap(l, size)) 
	for i := 0; i < l; i += size {
		end := i + size
		if end >= l {
			end = l
		}
		chunks = append(chunks, slice[i:end])
	}

	return chunks
}
`, fnName, eType, eType,
			fnName, eType, eType,
			eType))
	}

	return formatAndWriteFile(b.Bytes(), output)
}

func indexGenerator(eTypes []string, pkg string, output string) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("// Code generated by \"github.com/beanscc/rango/cmd/genericgenerator %s\"; DO NOT EDIT.\n", strings.Join(os.Args[1:], " ")))
	b.WriteString(fmt.Sprintf("\npackage %s\n", pkg))
	imports := []string{
		// `"fmt"`,
	}
	if len(imports) > 0 {
		b.WriteString("\nimport (\n\t")
		b.WriteString(strings.Join(imports, "\n\t"))
		b.WriteString("\n)\n")
	}

	b.WriteString("\n// Index 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1")
	b.WriteString("\nfunc Index(slice interface{}, x interface{}) interface{} {\n")
	b.WriteString("\tswitch slice.(type) {")
	for _, eType := range eTypes {
		b.WriteString(fmt.Sprintf("case []%s:\n", eType))
		b.WriteString(fmt.Sprintf("\treturn Index%ss(slice.([]%s), x.(%s))\n", stringutil.Snake2Camel(eType, true), eType, eType))
	}
	b.WriteString("\tdefault:\n\t\treturn indexReflect(slice, x)\n\t}\n}\n")

	for _, eType := range eTypes {
		fnName := fmt.Sprintf("Index%ss", stringutil.Snake2Camel(eType, true))
		b.WriteString(fmt.Sprintf(`
// %s 查找 x 在 slice 中第一次出现的索引位置；若不存在，则返回 -1
func %s(slice []%s, x %s) int {
	l := len(slice)
	for i := 0; i < l; i++ {
		if x == slice[i] {
			return i
		}
	}

	return -1
}
`, fnName,
			fnName, eType, eType))
	}

	return formatAndWriteFile(b.Bytes(), output)
}

func uniqueGenerator(eTypes []string, pkg string, output string) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("// Code generated by \"github.com/beanscc/rango/cmd/genericgenerator %s\"; DO NOT EDIT.\n", strings.Join(os.Args[1:], " ")))
	b.WriteString(fmt.Sprintf("\npackage %s\n", pkg))
	imports := []string{
		// `"fmt"`,
	}
	if len(imports) > 0 {
		b.WriteString("\nimport (\n\t")
		b.WriteString(strings.Join(imports, "\n\t"))
		b.WriteString("\n)\n")
	}

	cases := make([]string, 0, len(eTypes)*2)
	for _, eType := range eTypes {
		cases = append(cases, fmt.Sprintf("\tcase []%s:", eType))
		cases = append(cases, fmt.Sprintf("\t\treturn Unique%ss(slice.([]%s))", stringutil.Snake2Camel(eType, true), eType))
	}

	b.WriteString(fmt.Sprintf(`
// %s
func Unique(slice interface{}) interface{} {
	switch slice.(type) {
	%s
	default:
		return UniqueReflect(slice, nil)
	}
}
`, "Unique 切片去重", strings.Join(cases, "\n")))

	for _, eType := range eTypes {
		fnName := fmt.Sprintf("Unique%ss", stringutil.Snake2Camel(eType, true))
		b.WriteString(fmt.Sprintf(`
// %s 切片去重
func %s(slice []%s) []%s {
	l := len(slice)
	if l < 2 {
		return slice
	}
	
	resp := make([]%s, 0, l)
	m := make(map[%s]bool, l)
	for _, v := range slice {
		if _, ok := m[v]; !ok {
			resp = append(resp, v)
			m[v] = true
		}
	}

	return resp
}
`, fnName,
			fnName, eType, eType,
			eType,
			eType))
	}

	return formatAndWriteFile(b.Bytes(), output)
}

func joinGenerator(eTypes []string, pkg string, output string) ([]byte, error) {
	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("// Code generated by \"github.com/beanscc/rango/cmd/genericgenerator %s\"; DO NOT EDIT.\n", strings.Join(os.Args[1:], " ")))
	b.WriteString(fmt.Sprintf("\npackage %s\n", pkg))
	imports := []string{
		`"fmt"`,
		`"reflect"`,
		`"strconv"`,
	}
	if len(imports) > 0 {
		b.WriteString("\nimport (\n\t")
		b.WriteString(strings.Join(imports, "\n\t"))
		b.WriteString("\n)\n")
	}

	cases := make([]string, 0, len(eTypes)*2)
	for _, eType := range eTypes {
		cases = append(cases, fmt.Sprintf("\tcase []%s:", eType))
		cases = append(cases, fmt.Sprintf("\t\treturn Join%ss(slice.([]%s), sep)", stringutil.Snake2Camel(eType, true), eType))
	}

	b.WriteString(fmt.Sprintf(`
// %s
func Join(slice interface{}, sep string) string {
	switch slice.(type) {
	%s
	default:
		rv := reflect.ValueOf(slice)
		if rv.Kind() != reflect.Slice {
			panic("only support slice")
		}
		l := rv.Len()
		return JoinFunc(l, sep, func(i int) string {
			return fmt.Sprint(rv.Index(i).Interface())
		})
	}
}
`, "Join 切片元素以字符串形式拼接", strings.Join(cases, "\n")))

	for _, eType := range eTypes {
		var toStr string
		switch eType {
		case "string":
			toStr = `slice[i]`
		case "int":
			toStr = `strconv.Itoa(slice[i])`
		case "int8", "int16":
			toStr = `strconv.Itoa(int(slice[i]))`
		case "int32": // as rune
			toStr = `strconv.FormatInt(int64(slice[i]), 10)`
		case "int64":
			toStr = `strconv.FormatInt(slice[i], 10)`
		case "uint", "uint8", "uint16", "uint32":
			toStr = `strconv.FormatUint(uint64(slice[i]), 10)`
		case "uint64":
			toStr = `strconv.FormatUint(slice[i], 10)`
		case "float32":
			toStr = `strconv.FormatFloat(float64(slice[i]), 'f', -1, 64)`
		case "float64":
			toStr = `strconv.FormatFloat(slice[i], 'f', -1, 64)`
		case "bool":
			toStr = `strconv.FormatBool(slice[i])`
		default:
			toStr = `fmt.Sprint(slice[i])`
		}

		fnName := fmt.Sprintf("Join%ss", stringutil.Snake2Camel(eType, true))
		b.WriteString(fmt.Sprintf(`
// %s []%s 切片元素以字符串形式拼接
func %s(slice []%s, sep string) string {
	return JoinFunc(len(slice), sep, func(i int) string {
		return %s
	})
}
`, fnName, eType,
			fnName, eType,
			toStr))
	}

	return formatAndWriteFile(b.Bytes(), output)
}
