package proxy

/**
 * 自定义获取代理ip，将获取到的ip存储到allProxyIps中
 */

/*var (
	proxyHttp   *myhttp.HttpSend
	allProxyIps []string
)
const MaxIpSize = 5*/
/*
//全网代理IP，http://www.goubanjia.com/
//业务体验不好，暂时不用

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
}*/

// ProxyInfo
/** ProxyInfo
 * @Description: 快代理调用策略，快代理官网：https://www.kuaidaili.com/
 * @receiver self
 * @return []string
 * @return error
 */
/*func (self *Proxy) ProxyInfo() ([]string, error) {
	if proxyHttp == nil {
		proxyHttp = myhttp.NewHttpSend(config.PROXY)
	}
	proxyData, err := proxyHttp.Get()
	if err != nil {
		return nil, err
	}
	reg := regexp.MustCompile(`\d{1,3}([.]\d{1,3}){3}:\d{2,5}`)
	ips := reg.FindAllString(string(proxyData), -1)
	if len(ips) <= 0 {
		return nil, errors.New(string(proxyData))
	}

	proxyTest := myhttp.NewHttpSend("http://www.baidu.com/")
	for _, ip := range ips {
		_,err := proxyTest.GetWithProxy("http",ip)
		if err!=nil{
			continue
		}
		ip = "http://" + ip

		if len(allProxyIps) > MaxIpSize {
			allProxyIps = allProxyIps[len(allProxyIps)-MaxIpSize : len(allProxyIps)-1]
		}
		has := false
		for _, ipTmp := range allProxyIps {
			if ipTmp == ip {
				has = true
				break
			}
		}
		if !has {
			allProxyIps = append(allProxyIps, ip)
		}
	}
	log.Println(" *     当前ip代理数量: ", len(allProxyIps))

	return allProxyIps, err
}*/
