package main

import (
	"github.com/go-ego/gse"
)

var (
	text   = "旧金山湾金门大桥"
	new, _ = gse.New("zh,../../testdata/zh/test_dict.txt")

	seg gse.Segmenter
)

func main() {
	cut()

	// loadDict()
	loadDictEmbed()
	// loadDictMap()
	segment()
}

// loadDictEmbed supported from go1.16
func loadDictEmbed() { _ = "STUB: not implemented"; return }

func loadDict() {
	_ = "STUB: not implemented"
	// var seg gse.Segmenter
	return
}

func loadDictMap() { _ = "STUB: not implemented"; return }

func cut() { _ = "STUB: not implemented"; return }

func segment() { _ = "STUB: not implemented"; return }

// 金山/nr 旧金山/ns 湾/zg 旧金山湾/ns 金门/n 大桥/ns 金门大桥/nz

// fmt.Println(gse.ToString(segments, false))

// "旧金山湾/n 金门大桥/nz "

// 搜索模式主要用于给搜索引擎提供尽可能多的关键字

// "金山/nr 旧金山/ns 湾/zg 旧金山湾/ns 金门/n 大桥/ns 金门大桥/nz "
