package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	bytes, err := ioutil.ReadFile("test.json")
	if err != nil {
		panic("打开文件错误")
	}
	m := make(map[string]interface{})
	json.Unmarshal(bytes, &m)
	fmt.Println(m)

	// re := regexp.MustCompile(`(^\d+\.\d*)|(^\d*\.\d+)`)
	rei := regexp.MustCompile(`^[_a-zA-Z][_a-zA-Z\d]*`)
	fmt.Println(rei.MatchString("_s"))
}
