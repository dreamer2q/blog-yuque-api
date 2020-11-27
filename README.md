# yuque api for blog

![go-report](https://goreportcard.com/badge/github.com/dreamer2q/blog-yuque-api)

代理语雀开放API, 提供给个人博客使用

由于语雀的API并不支持跨域调用，因此这里做了一个语雀API的转发，用于自己创新实践博客小项目

 - API地址 `https://api.dreamer2q.wang`
 - 博客地址 `https://blo.dreamer2q.wang`    
 - 语雀文档 `https://www.yuque.com/dreamer2q/blog` 
 
## 语雀配置

 - 为了获取文章列表的时候`description`不为空，需要设置首页风格为专栏

![配置语雀](https://i.imgur.com/h6cb8Xk.png)

 - 配置token
 
在账户设置里面生成一个token即可，同时获取命名空间，这样语雀的配置就完成了。
 
## 获取文章列表

 - GET `api/articles` 
 - Response 
 
```json5

``` 

## 获取单个文章详情

 - GET `api/article/:slug`
 - Response
 
```json5

``` 

# 部署

由Git Action提供自动部署

# License

[mit](LICENSE)
