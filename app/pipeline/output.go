package pipeline

import (
	"sort"

	"github.com/jason-wj/pholcus/app/pipeline/collector"
	"github.com/jason-wj/pholcus/common/kafka"
	"github.com/jason-wj/pholcus/common/mgo"
	"github.com/jason-wj/pholcus/common/mysql"
	"github.com/jason-wj/pholcus/runtime/cache"
)

// 初始化输出方式列表collector.DataOutputLib
func init() {
	for out, _ := range collector.DataOutput {
		collector.DataOutputLib = append(collector.DataOutputLib, out)
	}
	sort.Strings(collector.DataOutputLib)
}

// 刷新输出方式的状态
func RefreshOutput() {
	switch cache.Task.OutType {
	case "mgo":
		mgo.Refresh()
	case "mysql":
		mysql.Refresh()
	case "kafka":
		kafka.Refresh()
	}
}
