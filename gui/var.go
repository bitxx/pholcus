package gui

import (
	"github.com/lxn/walk"
	"github.com/lxn/walk/declarative"

	. "github.com/jason-wj/pholcus/gui/model"
	"github.com/jason-wj/pholcus/runtime/cache"
	"github.com/jason-wj/pholcus/runtime/status"
)

// GUI输入
type Inputor struct {
	Spiders []*GUISpider
	*cache.AppConf
	Pausetime   int64
	ProxySecond int64
}

var (
	runStopBtn      *walk.PushButton
	pauseRecoverBtn *walk.PushButton
	setting         *walk.Composite
	mw              *walk.MainWindow
	runMode         *walk.GroupBox
	db              *walk.DataBinder
	ep              walk.ErrorPresenter
	mode            *walk.GroupBox
	host            *walk.Splitter
	spiderMenu      *SpiderMenu
)

var Input = &Inputor{
	AppConf:     cache.Task,
	Pausetime:   cache.Task.Pausetime,
	ProxySecond: cache.Task.ProxySecond,
}

//****************************************GUI内容显示配置*******************************************\\

// 输出选项
var outputList []declarative.RadioButton

// 下拉菜单辅助结构体
type KV struct {
	Key   string
	Int   int
	Int64 int64
}

// 暂停时间选项及运行模式选项
var GuiOpt = struct {
	Mode        []*KV
	Pausetime   []*KV
	ProxySecond []*KV
}{
	Mode: []*KV{
		{Key: "单机", Int: status.OFFLINE},
		{Key: "服务器", Int: status.SERVER},
		{Key: "客户端", Int: status.CLIENT},
	},
	Pausetime: []*KV{
		{Key: "无暂停", Int64: 0},
		{Key: "0.1 秒", Int64: 100},
		{Key: "0.3 秒", Int64: 300},
		{Key: "0.5 秒", Int64: 500},
		{Key: "1 秒", Int64: 1000},
		{Key: "3 秒", Int64: 3000},
		{Key: "5 秒", Int64: 5000},
		{Key: "10 秒", Int64: 10000},
		{Key: "15 秒", Int64: 15000},
		{Key: "20 秒", Int64: 20000},
		{Key: "30 秒", Int64: 30000},
		{Key: "60 秒", Int64: 60000},
	},
	ProxySecond: []*KV{
		{Key: "不使用代理", Int64: 0},
		{Key: "1 秒钟", Int64: 1},
		{Key: "3 秒钟", Int64: 3},
		{Key: "5 秒钟", Int64: 5},
		{Key: "10 秒钟", Int64: 10},
		{Key: "15 秒钟", Int64: 15},
		{Key: "20 秒钟", Int64: 20},
		{Key: "30 秒钟", Int64: 30},
		{Key: "45 秒钟", Int64: 45},
		{Key: "60 秒钟", Int64: 60},
		{Key: "120 秒钟", Int64: 120},
		{Key: "180 秒钟", Int64: 180},
	},
}
