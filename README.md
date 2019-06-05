# Pholcus爬虫框架（改造版）
原官方爬虫框架[henrylee2cn/pholcus](https://github.com/jason-wj/pholcus)基本停止更新，由于个人项目需要，对本项目做了一些修改和完善。
有兴趣对可以看看原官方项目的介绍。
感谢：[henrylee2cn](https://github.com/jason-wj)

## 使用
需要将项目放在$GOPATH/github.com/jason-wj/目录中

## 后续打算（仅仅只是打算，根据业余时间来安排）
1. 将pipe部分自定义化，能够根据需要灵活实现
2. 添加新的功能，将爬虫结果输出到nsq中

## 该项目个人使用感悟
1. 能够将精力集中在规则的实现和完善中，不用专门去考虑输出结果和方式，已经给出固定模板。
2. 每只爬虫独立化，互不干扰
3. 规则简单易懂，规则格式已确定，任何网站的爬虫规则，都按照规则格式来实现就行，不同的网站也很容易理解该爬虫规则（原先使用过scrapy，随意实现规则，写到后面别人再来看会花费很多不必要的时间去理解）

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

5. 将[henrylee2cn/teleport](https://github.com/henrylee2cn/teleport)和[henrylee2cn/goutil](https://github.com/henrylee2cn/goutil)两个辅助源码直接放在`/pholcus/common`目录中

剩余调整将会根据后续需要来逐步调整。。。
