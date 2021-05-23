package debuglog

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"regexp"
	"runtime"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

var (
	// default logger
	logger = log.New(os.Stdout, "DEBUG ", log.Ldate|log.Ltime)
)

// 转换8进制utf-8字符串到中文
// eg: `\346\200\241` -> 怡
func convertOctonaryUtf8(in string) string {
	s := []byte(in)
	reg := regexp.MustCompile(`\\[0-7]{3}`)

	out := reg.ReplaceAllFunc(s,
		func(b []byte) []byte {
			i, _ := strconv.ParseInt(string(b[1:]), 8, 0)
			return []byte{byte(i)}
		})
	return string(out)
}

func short(s string) string {
	short := s
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			short = s[i+1:]
			break
		}
	}
	return short
}

func logStr(v interface{}, mid string, prefix ...string) string {
	if len(prefix) > 0 {
		return fmt.Sprintf("%+v%s%+v\n", prefix, mid, v)
	}
	return fmt.Sprintf("%+v\n", v)
}

func print(s string) {
	function, file, line, _ := runtime.Caller(2)
	funcName := runtime.FuncForPC(function).Name()
	logger.Printf("%s:%d %s %s", short(file), line, short(funcName), s)
}

// 打印一个值
func Val(v interface{}, prefix ...string) {
	print(logStr(v, " = ", prefix...))
}

// Spew 打印一个值
func SpewVal(v interface{}, prefix ...string) {
	s := spew.Sdump(v)
	print(logStr(s, " = ↙↙↙\n", prefix...))
}

// Spew 打印一个包含8进制utf-8字符串的值
func OctUtf8Val(v interface{}, prefix ...string) {
	str := spew.Sdump(v)
	s := convertOctonaryUtf8(str)
	print(logStr(s, " = ↙↙↙\n", prefix...))
}

// 转换成json后打印
func ToJson(v interface{}, prefix ...string) {
	jsonData, err := json.Marshal(v)
	if err != nil {
		print(logStr(fmt.Sprintf("Print json val err: %s\n", err), " ", prefix...))
	}
	print(logStr(string(jsonData), " = ", prefix...))
}

// 转换成json后打印
func ToJsonPretty(v interface{}, prefix ...string) {
	jsonData, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		print(logStr(fmt.Sprintf("Print json val err: %s\n", err), " ", prefix...))
	}
	print(logStr(string(jsonData), " = ↙↙↙\n", prefix...))
}
