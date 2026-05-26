/*

gse 分词服务器同时提供了两种模式：

	"/"	分词演示网页
	"/add" 把词语添加到字典中
		输入 text, freq, pos 参数
		输出 JSON :
			{"code":200,"text":"ok"}
	"/json"	JSON 格式的 RPC 服务
		输入：
			POST 或 GET 模式输入 text 参数
		输出 JSON 格式：
			{
				segments:[
					{"text":"服务器", "pos":"n"},
					{"text":"指令", "pos":"n"},
					...
				]
			}


测试服务器

*/

package main

import (
	"flag"
	"fmt"
	"log"
	"runtime"

	"net/http"

	"github.com/go-ego/gse"
)

var (
	seg = gse.Segmenter{}

	host = flag.String("host", "", "HTTP服务器主机名")
	port = flag.Int("port", 8080, "HTTP服务器端口")

	hmm          = flag.Bool("hmm", false, "use hmm")
	dict         = flag.String("dict", "../data/dict/dictionary.txt", "词典文件")
	staticFolder = flag.String("static_folder", "static", "静态页面存放的目录")
)

// JsonResponse []*segments json response
type JsonResponse struct {
	Segments []*Segment `json:"segments"`
	Err      error      `json:"err"`
}

// Segment segment json struct
type Segment struct {
	Text string `json:"text"`
	Pos  string `json:"pos"`
}

// JsonResp json response return []string
type JsonResp struct {
	Seg []string
}

// Resp http response
type Resp struct {
	Code int    `json:"code"`
	Text string `json:"text"`
}

// JsonRpcServer start json rpc server
func JsonRpcServer(w http.ResponseWriter, req *http.Request) {
	_ = "STUB: not implemented"
	// 得到要分词的文本
	return
}

// 分词

// 整理为输出格式

func addToken(w http.ResponseWriter, req *http.Request) { _ = "STUB: not implemented"; return }

func main() {
	flag.Parse()

	// 将线程数设置为 CPU数
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 初始化分词器
	seg.LoadDict(*dict)

	http.HandleFunc("/add", addToken)
	http.HandleFunc("/json", JsonRpcServer)
	http.Handle("/", http.FileServer(http.Dir(*staticFolder)))

	log.Printf("%s %s:%d \n", "Server listen: ", *host, *port)
	http.ListenAndServe(fmt.Sprintf("%s:%d", *host, *port), nil)
}
