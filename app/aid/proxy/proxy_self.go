package proxy

import (
	"errors"
	"pholcus/common/net/myhttp"
	"log"
	"regexp"
	"time"

	"pholcus/config"
)

/**
 * 自定义获取代理ip，将获取到的ip存储到allProxyIps中
 */

var (
	proxyHttp *myhttp.HttpSend
	allProxyIps       []string
)

const MaxIpSize = 7

func (self *Proxy) ProxyInfo() ([]string,error) {
	//每3秒钟请求一次
	if time.Now().Unix()%3 != 0 {
		return nil,nil
	}
	if proxyHttp == nil {
		proxyHttp = myhttp.NewHttpSend(config.PROXY)
	}
	proxyData, err := proxyHttp.Get()
	if err != nil {
		return nil,err
	}
	reg := regexp.MustCompile(`\d{1,3}([.]\d{1,3}){3}:\d{2,5}`)
	res := reg.FindString(string(proxyData))
	if res == "" {
		return nil,errors.New(string(proxyData))
	}

	if len(allProxyIps) > MaxIpSize {
		allProxyIps = allProxyIps[len(allProxyIps)-MaxIpSize : len(allProxyIps)-1]
	}
	has := false
	for _, ip := range allProxyIps {
		if ip == res {
			has = true
			break
		}
	}
	if !has {
		res = "http://" + res
		allProxyIps = append(allProxyIps, res)
	}
	log.Printf(" *     添加新的IP: %s \n", res)

	return allProxyIps,err
}
