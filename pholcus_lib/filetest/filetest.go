package pholcus_lib

// 基础包
import (
	// "pholcus/common/goquery"                          //DOM解析
	"pholcus/app/downloader/request" //必需
	. "pholcus/app/spider"           //必需
	// . "pholcus/app/spider/common" //选用
	// "pholcus/logs"
	// net包
	// "net/http" //设置http.Header
	// "net/url"
	// 编码包
	// "encoding/xml"
	//"encoding/json"
	// 字符串处理包
	//"regexp"
	// "strconv"
	//	"strings"
	// 其他包
	// "fmt"
	// "math"
	// "time"
)

func init() {
	FileTest.Register()
}

var FileTest = &Spider{
	Name:        "文件下载测试",
	Description: "文件下载测试",
	// Pausetime: 300,
	// Keyin:   KEYIN,
	// Limit:        LIMIT,
	EnableCookie: false,
	RuleTree: &RuleTree{
		Root: func(ctx *Context) {
			ctx.AddQueue(&request.Request{
				Url:          "https://www.baidu.com/img/bd_logo1.png",
				Rule:         "百度图片",
				ConnTimeout:  -1,
				DownloaderID: 0, //图片等多媒体文件必须使用0（surfer surf go原生下载器）
			})
			ctx.AddQueue(&request.Request{
				Url:          "https://pholcus",
				Rule:         "Pholcus页面",
				ConnTimeout:  -1,
				DownloaderID: 0, //文本文件可使用0或者1（0：surfer surf go原生下载器；1：surfer plantomjs内核）
			})
		},

		Trunk: map[string]*Rule{

			"百度图片": {
				ParseFunc: func(ctx *Context) {
					ctx.FileOutput("baidu") // 等价于ctx.AddFile("baidu")
				},
			},
			"Pholcus页面": {
				ParseFunc: func(ctx *Context) {
					ctx.FileOutput() // 等价于ctx.AddFile()
				},
			},
		},
	},
}
