package pipeline

import (
	"sort"

	"github.com/bitxx/pholcus/app/pipeline/collector"
	"github.com/bitxx/pholcus/common/kafka"
	"github.com/bitxx/pholcus/common/mgo"
	"github.com/bitxx/pholcus/common/mysql"
	"github.com/bitxx/pholcus/runtime/cache"
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
