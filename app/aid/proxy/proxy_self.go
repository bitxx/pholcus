package proxy

import (
	"errors"
	"github.com/jason-wj/pholcus/common/net/myhttp"
	"log"
	"regexp"
	"time"

	"github.com/jason-wj/pholcus/config"
)

/**
 * 自定义获取代理ip，将获取到的ip存储到self.allProxyIps中
 */

var (
	proxyHttp *myhttp.HttpSend
)

const MaxIpSize = 7

func (self *Proxy) ProxyInfo() error {
	//每3秒钟请求一次
	if time.Now().Unix()%3 != 0 {
		return nil
	}
	if proxyHttp == nil {
		proxyHttp = myhttp.NewHttpSend(config.PROXY)
	}
	proxyData, err := proxyHttp.Get()
	if err != nil {
		return err
	}
	reg := regexp.MustCompile(`\d{1,3}([.]\d{1,3}){3}:\d{2,5}`)
	res := reg.FindString(string(proxyData))
	if res == "" {
		return errors.New(string(proxyData))
	}

	if len(self.allProxyIps) > MaxIpSize {
		self.allProxyIps = self.allProxyIps[len(self.allProxyIps)-MaxIpSize : len(self.allProxyIps)-1]
	}
	has := false
	for _, ip := range self.allProxyIps {
		if ip == res {
			has = true
			break
		}
	}
	if !has {
		self.allProxyIps = append(self.allProxyIps, res)
	}
	log.Printf(" *     添加新的IP: %s 条\n", res)

	return err
}
