package main

type Config struct {
	Addr      string //本地监听地址
	Token     string //语雀 Token
	Namespace string //语雀 Namespace
}

var (
	config = &Config{
		Addr:      ":8080",
		// 创建Token的时候不要选择write权限，仅read即可
		Token:     "8yx2GugoI9kLsKAQ9pDR1f2S371Vwh4D8EHzx1OW",
		Namespace: "dreamer2q/blog",
	}
)
