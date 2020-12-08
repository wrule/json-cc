package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, err := ioutil.ReadFile("test.json")
	if err != nil {
		panic("打开文件错误")
	}
	m := make(map[string]interface{})
	json.Unmarshal(bytes, &m)
	fmt.Println(m)
}
