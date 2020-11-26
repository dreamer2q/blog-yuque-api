package main

type Config struct {
	Addr      string
	Token     string //语雀 Token
	Namespace string //语雀 Namespace
}

var (
	config = &Config{
		Addr:      ":8080",
		Token:     "8yx2GugoI9kLsKAQ9pDR1f2S371Vwh4D8EHzx1OW",
		Namespace: "dreamer2q/blog",
	}
)
