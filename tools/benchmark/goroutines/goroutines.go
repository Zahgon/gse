// 测试 gse 并行分词速度

package main

import (
	"log"
	"runtime"
	"time"

	"github.com/go-ego/gse"
)

var (
	seg        = gse.Segmenter{}
	numThreads = runtime.NumCPU()
	task       = make(chan []byte, numThreads*40)
	done       = make(chan bool, numThreads)
	numRuns    = 50
)

func worker() { _ = "STUB: not implemented"; return }

func openBook() (int, [][]byte) {
	_ = "STUB: not implemented"
	// 打开将要分词的文件
	return 0, nil
}

// 逐行读入

func main() {
	// 将线程数设置为CPU数
	runtime.GOMAXPROCS(numThreads)

	// 载入词典
	seg.LoadDict("../../data/dict/dictionary.txt")
	size, lines := openBook()

	// 启动工作线程
	for i := 0; i < numThreads; i++ {
		go worker()
	}
	log.Print("开始分词")

	// 记录时间
	t0 := time.Now()

	// 并行分词
	for i := 0; i < numRuns; i++ {
		for _, l := range lines {
			task <- l
		}
	}
	close(task)

	// 确保分词完成
	for i := 0; i < numThreads; i++ {
		<-done
	}

	// 记录时间并计算分词速度
	t1 := time.Now()
	log.Printf("分词花费时间 %v", t1.Sub(t0))

	ts := float64(size*numRuns) / t1.Sub(t0).Seconds() / (1024 * 1024)
	log.Printf("分词速度 %f MB/s", ts)
}
