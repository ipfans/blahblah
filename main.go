package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/mozillazg/go-pinyin"
)

func main() {
	name := os.Args[1]
	b, _ := os.ReadFile("words.json")
	var m map[string][]map[string]string
	json.Unmarshal(b, &m)
	a := pinyin.NewArgs()
	pys := pinyin.Convert(name, &a)
	var spell string
	if len(pys) > 1 {
		slices := make([]string, len(pys))
		for k := range pys {
			slices[k] = pys[k][0]
		}
		spell = strings.Join(slices, "'")
	} else {
		spell = pys[0][0]
	}
	for _, v := range m["words"] {
		if strings.HasPrefix(v["name"], spell) || strings.HasSuffix(v["name"], spell) {
			fmt.Println(v["word"], v["pinyin"], spell)
		} else if strings.Contains(v["pinyin"], "'"+spell+"'") {
			fmt.Println(v["word"], v["pinyin"], spell)
		}
	}
}
