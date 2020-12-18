package main

import (
	_ "go-gf-blog/boot"
	_ "go-gf-blog/router"

	"github.com/gogf/gf/frame/g"
)

// @title       go-gf-blog
// @version     1.0
// @description `GoFrame`基础开发框架搭建的博客
// @schemes     http
func main() {
	g.Server().Run()
}
