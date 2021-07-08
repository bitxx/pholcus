package proxy

import (
	"github.com/jason-wj/pholcus/config"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"time"
)

//author: 代理模块基本重构 wj
type Proxy struct {
	proxyIPTypeRegexp *regexp.Regexp
	allProxyIps       []string
	ticker            *time.Ticker
	tickSecond        int64
}

func New() *Proxy {
	p := &Proxy{
		proxyIPTypeRegexp: regexp.MustCompile(`https?://([\w]*:[\w]*@)?[0-9]+\.[0-9]+\.[0-9]+\.[0-9]+:[0-9]+`),
	}
	go p.Update()
	return p
}

// 代理IP数量
func (self *Proxy) Count() int32 {
	return int32(len(self.allProxyIps))
}

// 更新代理IP列表
/*func (self *Proxy) Update() *Proxy {
	allProxyIps,err := self.ProxyInfo()
	if err != nil {
		log.Printf("代理读取错误：%s", err.Error())
		return self
	}

	self.allProxyIps = allProxyIps
	if len(self.allProxyIps)<=0{
		return self
	}
	log.Printf(" *     读取代理IP: %v 条\n", len(self.allProxyIps))
	return self
}*/

// Update 更新代理IP列表
func (self *Proxy) Update() *Proxy {
	f, err := os.Open(config.PROXY)
	if err != nil {
		// logs.Log.Error("Error: %v\n", err)
		return self
	}
	b, _ := ioutil.ReadAll(f)
	f.Close()

	proxysIPType := self.proxyIPTypeRegexp.FindAllString(string(b), -1)
	for _, proxy := range proxysIPType {
		self.allProxyIps = append(self.allProxyIps, proxy)
	}
	log.Printf(" *     读取代理IP: %v 条\n", len(self.allProxyIps))
	return self
}

// 更新继时器
func (self *Proxy) UpdateTicker(tickSecond int64) {
	self.tickSecond = tickSecond
	self.ticker = time.NewTicker(time.Duration(self.tickSecond) * time.Second)
}

// 获取本次循环中未使用的代理IP及其响应时长
var count = 0

func (self *Proxy) GetOne(u string) (curProxy string) {

	select {
	case <-self.ticker.C:
		self.allProxyIps = []string{}
		self.Update()
	default:
	}

	if len(self.allProxyIps) <= 0 {
		return "" //没有代理则使用本机ip
	}
	if count >= len(self.allProxyIps) {
		count = 0
	}
	//TODO 可自身根据需要设计算法随机获取代理
	s := self.allProxyIps[count]
	count++

	return s
}
