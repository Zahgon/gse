package main

import (
	"flag"

	"github.com/go-ego/gse"
	"github.com/go-ego/gse/hmm/pos"
)

var (
	seg    gse.Segmenter
	posSeg pos.Segmenter

	text  = "《复仇者联盟3：无限战争》是全片使用IMAX摄影机拍摄制作的的科幻片."
	text1 = flag.String("text", text, "要分词的文本")

	text2 = "西雅图地标建筑, Seattle Space Needle, 西雅图太空针. Sky tree."
)

func main() {
	flag.Parse()
	// keep the origin capital letter
	// gse.ToLower = false

	// Loading the default dictionary
	seg.LoadDict()
	// Loading the default dictionary with embed
	// seg.LoadDictEmbed()
	//
	// Loading the simple chinese dictionary
	// seg.LoadDict("zh_s")
	// seg.LoadDictEmbed("zh_s")
	//
	// Loading the traditional chinese dictionary
	// seg.LoadDict("zh_t")
	//
	// Loading the japanese dictionary
	// seg.LoadDict("jp")
	//
	// seg.LoadDict("../data/dict/dictionary.txt")
	//
	// Loading the custom dictionary
	// seg.LoadDict("zh,../../testdata/zh/test_dict.txt,../../testdata/zh/test_dict1.txt")

	addToken()

	cut()
	//
	cutPos()
	segCut()

	extAndRank(seg)
}

func addToken() { _ = "STUB: not implemented"; return }

// seg.AddTokenForce("上海东方明珠广播电视塔", 100, "n")
//

// seg.CalcToken()

// 使用 DAG 或 HMM 模式分词
func cut() {
	_ = "STUB: not implemented"
	// "《复仇者联盟3：无限战争》是全片使用IMAX摄影机拍摄制作的的科幻片."
	return
}

// use DAG and HMM

// cut use hmm:  [《复仇者联盟3：无限战争》 是 全片 使用 imax 摄影机 拍摄 制作 的 的 科幻片 .]

// cut:  [《 复仇者 联盟 3 ： 无限 战争 》 是 全片 使用 imax 摄影机 拍摄 制作 的 的 科幻片 .]

// cut search use hmm:  [复仇 仇者 联盟 无限 战争 复仇者 《复仇者联盟3：无限战争》 是 全片 使用 imax 摄影 摄影机 拍摄 制作 的 的 科幻 科幻片 .]

// cut search:  [《 复仇 者 复仇者 联盟 3 ： 无限 战争 》 是 全片 使用 imax 摄影 机 摄影机 拍摄 制作 的 的 科幻 片 科幻片 .]

// cut all:  [《复仇者联盟3：无限战争》 复仇 复仇者 仇者 联盟 3 ： 无限 战争 》 是 全片 使用 i m a x 摄影 摄影机 拍摄 摄制 制作 的 的 科幻 科幻片 .]

// cut all to string:  《复仇者联盟3：无限战争》, 复仇, 复仇者, 仇者, 联盟, 3, ：, 无限, 战争, 》, 是, 全片, 使用, i, m, a, x, 摄影, 摄影机, 拍摄, 摄制, 制作, 的, 的, 科幻, 科幻片, .

func analyzeAndTrim(cut []string) { _ = "STUB: not implemented"; return }

// analyze the segment:

// cut all:  [复仇者联盟3无限战争 复仇 复仇者 仇者 联盟 3 无限 战争 是 全片 使用 i m a x 摄影 摄影机 拍摄 摄制 制作 的 的 科幻 科幻片]

// 西雅图/nr 地标/n 建筑/n ,/x  /x seattle/x  /x space needle/n ,/x  /x 西雅图太空针/n ./x  /x sky/x  /x tree/x ./x

// [西雅图 地标 建筑 ,   seattle   space needle ,   西雅图太空针 .   sky   tree .]

func cutPos() {
	_ = "STUB: not implemented"
	// "西雅图地标建筑, Seattle Space Needle, 西雅图太空针. Sky tree."
	return
}

// pos:  [{西雅图 nr} {地标 n} {建筑 n} {, x} {  x} {seattle x} {  x} {space needle n} {, x} {  x} {西雅图太空针 n} {. x} {  x} {sky x} {  x} {tree x} {. x}]

// trim pos:  [{西雅图 nr} {地标 n} {建筑 n} {, x} {  x} {seattle x} {  x} {space needle n} {, x} {  x} {西雅图太空针 n} {. x} {  x} {sky x} {  x} {tree x} {. x}]

// pos:  [{《 x} {复仇 v} {者 k} {联盟 j} {3 x} {： x} {无限 v} {战争 n} {》 x} {是 v} {全片 n} {使用 v} {imax eng} {摄影 n} {机 n} {拍摄 v} {制作 vn} {的的 u} {科幻 n} {片 q} {. m}]

// trim pos:  [{《 x} {复仇 v} {者 k} {联盟 j} {3 x} {： x} {无限 v} {战争 n} {》 x} {是 v} {全片 n} {使用 v} {imax eng} {摄影 n} {机 n} {拍摄 v} {制作 vn} {的的 u} {科幻 n} {片 q} {. m}]

// 使用最短路径和动态规划分词
func segCut() { _ = "STUB: not implemented"; return }

// 《/x 复仇/v 者/k 复仇者/n 联盟/j 3/x ：/x 无限/v 战争/n 》/x 是/v 全片/n 使用/v imax/x 摄影/n 机/n 摄影机/n 拍摄/v 制作/vn 的/uj 的/uj 科幻/n 片/q 科幻片/n ./x

// log.Println(gse.ToString(segs, false))

// 西雅图/nr 地标/n 建筑/n ,/x  /x seattle/x  /x space needle/n ,/x  /x 西雅图太空针/n ./x  /x sky/x  /x tree/x ./x

// 搜索模式主要用于给搜索引擎提供尽可能多的关键字
// segs := seg.ModeSegment(text2, true)

// 搜索模式:  西雅图/nr 地标/n 建筑/n ,/x  /x seattle/x  /x space needle/n ,/x  /x 西雅图太空针/n ./x  /x sky/x  /x tree/x ./x

// to slice [西雅图 地标 建筑 ,   seattle   space needle ,   西雅图太空针 .   sky   tree .]

func extAndRank(segs gse.Segmenter) { _ = "STUB: not implemented"; return }

// segments:  5 [{科幻片 1.6002581704125} {全片 1.449761569875} {摄影机 1.2764747747375} {拍摄 0.9690261695075} {制作 0.8246043033375}]

// results:  [{机 1} {全片 0.9931964427972227} {摄影 0.984870660504368} {使用 0.9769826633059524} {是 0.8489363954683677}]
