# Introduction

因为秋招裂开了，转变心态的同时大概花了一个礼拜时间写了一下苍穹外卖的大部分代码。

后面那个涉及到催单之后的由于感觉太特定了，就没写了，其他业务基本上都写完了。

# Tech Stack

技术栈喜闻乐见，用的都是比较常用的。github上面基本上都能找到文档，看不懂文档b站也基本上有那种简明教程。

`Gin`： 基本的路由包；

`GORM`：orm包，用来写mysql的；

`go-redis`：golang的redis client的包；

`gRPC`: 微服务的包，写proto生成对应文件用的；

`wire`： 依赖注入的包，一个普通的服务要调的接口还比较少，但是到了`api-gateway`之后要调的就多了，所以用这个依赖注入方便一点；

# Appendix

`langchaingo` 调ai接口的包，我给那个semeal和dish的description的更新和创建业务中，加入了ai自动整理优化内容的操作

`kafka-go` 主要整合业务log用的，因为正常来说不同service是部署在不同的机器上吗，如果出了问题还要一台一台机器，所以就利用kafka的机制针对log创建event，异步记日志，然后持久化到`mysql`做日志分析，为什么不用ES因为不会

`zap` 结构化日志的包，b站有很详细的简明教程

`DDD` 领域驱动设计，每个服务内部分`application` `interfaces` `domain` `infrastructure`



