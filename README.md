# yuque api for blog

![go-report](https://goreportcard.com/badge/github.com/dreamer2q/blog-yuque-api)

代理语雀开放API, 提供给个人博客使用

由于语雀的API并不支持跨域调用，因此这里做了一个语雀API的转发，用于自己创新实践博客小项目

 - API地址 `https://api.dreamer2q.wang`
 - 博客地址 `https://blo.dreamer2q.wang`    
 - 语雀文档 `https://www.yuque.com/dreamer2q/blog` 
 
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
