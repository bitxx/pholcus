# Pholcus爬虫框架（改造版）
原官方爬虫框架[henrylee2cn/pholcus](https://github.com/henrylee2cn/pholcus)基本停止更新，由于个人项目需要，对本项目做了一些修改和完善。
有兴趣对可以看看原官方项目的介绍。
感谢：[henrylee2cn](https://github.com/henrylee2cn)

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

## 后续打算（仅仅只是打算，根据业余时间来安排）
1. 将pipe部分自定义化，能够根据需要灵活实现
2. 添加新的功能，将爬虫结果输出到nsq中

## 该项目个人使用感悟
1. 能够将精力集中在规则的实现和完善中，不用专门去考虑输出结果和方式，已经给出固定模板。
2. 每只爬虫独立化，互不干扰
3. 规则简单易懂，规则格式已确定，任何网站的爬虫规则，都按照规则格式来实现就行，不同的网站也很容易理解该爬虫规则（原先使用过scrapy，随意实现规则，写到后面别人再来看会花费很多不必要的时间去理解）


## 增量更新：pholcus是根据链接是否已经爬取过来判断是否继续爬取该链接，因此只要在链接之后加入一个时间戳参数，此时pholcus就会认为这是一个新的链接

另外，该功能有一个极大的作用，就是对于一些不影响去重的末端链接（一次爬取过程的最后一次请求）是可以忽视的，这样可以为数据库节省大量空间。
比如，一个城市有数百万个户信息，我们获取到户信息后就结束了，如果要把每个户信息的url都存下来，这种资源消耗是蛮大的，如果综合业务考虑后，这些url都可以不用存储，这样会极大的减轻去重记录的压力。

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

3. 多次发起请求时候，head会被重复利用，这样有的爬虫规则下， 会造成请求错误，始终无法继续（会误以为是ip被封），为此，注释掉`pholcus/app/downloader/surfer/param.go`第60行如下代码：
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


剩余调整将会根据后续需要来逐步调整。。。
