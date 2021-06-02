# Pholcus爬虫框架（改造版）
原官方爬虫框架[henrylee2cn/pholcus](https://github.com/henrylee2cn/pholcus)基本停止更新，由于个人学习需要，对本项目做了一些修改和完善。
有兴趣对可以看看原官方项目的介绍。
感谢：[henrylee2cn](https://github.com/henrylee2cn)
因爬虫项目在政策上的敏感性，本项目将根据原作者的项目状态而随时做出调整（如项目删除、免责声明等）
再次声明：本项目仅供学习

# 免责声明
本软件仅用于学术研究，使用者需遵守其所在地的相关法律法规，请勿用于非法用途！！ 如在中国大陆频频爆出爬虫开发者涉诉与违规的[新闻](https://github.com/HiddenStrawberry/Crawler_Illegal_Cases_In_China)。
郑重声明：因违法违规使用造成的一切后果，使用者自行承担！！

## 2019-08-17更新
1. 当历史记录中有错误信息时，用户可以在Spider中设置是否只爬取这些历史记录，若不设置，则不仅会爬取这些错误信息，同时还会继续爬取定义的规则。该功能主要是针对增量更新时，只爬取一遍网站，收集到失败记录后，只爬取失败记录即可，没必要为了爬取败记录就同时爬取整个网站网。

## 2019-07-21更新
1. 多处细节和异常更新
2. 代理模块重构，简化代理逻辑
3. views的操作，请参考 /pholcus/web/bindata_assetfs_usage，使用前请现在该目录下将views.zip解压并根据自己需要修改页面
4. 目前可在/pholcus/config/config.go中根据需要修改日志文件名称，方便区分每个网站的爬虫
5. 其余众多功能改造，就不一一细说了

## 以下历史记录，部分可供参考

## 使用
需要将项目放在$GOPATH/github.com/jason-wj/目录中

## 使用技巧
1. 每个链接的请求，最好设置一下链接`DialTimeout`和`ConnTimeout`，默认框架提供的是2分钟，这个大批量爬取时候，这个时间影响还是很大的，我控制在15秒左右。
2. 使用代理，结合上面的方式是最佳选择

## 该项目个人使用感悟
重量级爬取，性能和体验比很多开源项目要好很多，无需怀疑。

## 当前如下修改和完善
1. 为方便在爬虫规则中调用ctx.Sleep时能够动态切换爬虫间隔频率，在`pholcus/app/crawler/crawler.go`第200行处加入：
```
self.setPauseTime()
```  

2. 该框架是使用map将爬虫结果导入到mongo中的，原先爬虫上下文中的参数默认值不能为空（只能搞成""空字符串），这就回导致json解析数据时候同时解析了该参数，后面不仅浪费空间（etl时候会将该空字段再次加入到mongo，很浪费资源），而且容易造成误理解。  
为此，在`pholcus/app/downloader/request/request.go`第282行将如下代码注释：
```
if defaultValue == nil {
	panic("*Request.GetTemp()的defaultValue不能为nil，错误位置：key=" + key)
}
```  

3. (待定，当前已还原)多次发起请求时候，head会被重复利用，这样有的爬虫规则下， 会造成请求错误，始终无法继续（会误以为是ip被封），为此，注释掉`pholcus/app/downloader/surfer/param.go`第60行如下代码：
```
param.header = req.GetHeader()
```
4. 当页面请求错误，获取不到数据时候（提示`convert err xxx`），此时如果错误数量超过goroutine限制的上限，则会陷入死锁状态，为此需要在`pholcus/app/spider/context.go`第643行加入如下一行代码：
```
self.text = []byte("") //防止self.text为nil
```

5. 新增方法，用于获取请求到的页面的原始[]byte数据。原先没有提供时，如果要获取图片或者自行处理数据是很难搞的，为此在`pholcus/app/spider/context.go`第579行处加入如下代码：  
```  
// GetBytes returns plain bytes crawled.
func (self *Context) GetBytes() []byte {
	if self.text == nil {
		self.initText()
	}
	return self.text
}
```

6. 当response编码为"image/jpeg"或者没有指定编码时，不要进行转码操作（默认会转为utf-8，会影响图片等内容等展示）,需要在`pholcus/app/spider/context.go`第641行加入一项：
```  
"image/jpeg",""
```

7. 部分网站可能会发生url变化，此时继续爬取，会被识别为新的url来爬取，会造成和旧的url爬的数据重复。为了解决这个问题，需要在`pholcus/app/downloader/request/request.go`第19行加入：
```  
UrlAlias      string          //url别名，主要是为了防止网站url发生变化，影响去重。（若网站url变化，只需要在此处加入旧的url就行）
```
同时在142行加入：
```  
// 请求的唯一识别码
func (self *Request) Unique() string {
	if self.unique == "" {
		if self.UrlAlias != "" {
			block := md5.Sum([]byte(self.Spider + self.Rule + self.UrlAlias + self.Method))
			self.unique = hex.EncodeToString(block[:])
		} else {
			block := md5.Sum([]byte(self.Spider + self.Rule + self.Url + self.Method))
			self.unique = hex.EncodeToString(block[:])
		}
	}
	return self.unique
}
```
以后只要在Request中，指定`UrlAlias`的旧根地址即可

8. 为更直观展示代理使用时候的错误提示，在`pholcus/app/aid/proxy/proxy.go`第239行加入：
ps：曾犯下一个错误，代理测试始终报错，后来才知道，代理ip需要加上`http://`前缀，就是因为源码中忽略了下面的错误提示
```  
if err != nil {
	logs.Log.Informational(" *     [%v]代理测试发生错误：" + err.Error())
}
```

9. 将[henrylee2cn/teleport](https://github.com/henrylee2cn/teleport)和[henrylee2cn/goutil](https://github.com/jason-wj/pholcus/common/goutil)两个辅助源码直接放在`/pholcus/common`目录中

10. 加入爬虫规则示例包到项目根目录

11. 删除请求成功，但返回内容不正确的记录（比如请求的数据页面，返回的却是验证码页面），此时框架认为请求成功成功，而我们基于页面内容认为是错误的，那我们就可以在爬取期间将该请求链接删除，防止去重造成不再访问该网站。
就是说,业务判断认为失败的链接，想要从pholcus的历史记录中删除（如果不删除，那下次该链接更新了内容就很难判断是否要爬取该页面了）
```
ctx.GetSpider().DeleteSuccess(ctx.Request.Unique())
```

12. 可手动判断是否要某条链接作为去重处理，Request中加入参数：NeedUrlUnique，加入库中，默认不去重:

```txt
app/scheduler/matrix.go 169行
if ok && req.NeedUrlUnique {

}

```

13. mongo完善，支持用admin的username和password来加密村粗，若username为空，则认为不需要账号和密码



剩余调整将会根据后续需要来逐步调整。。。
