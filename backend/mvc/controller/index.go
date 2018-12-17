package controller

import (
	"net/http"
	"github.com/owenliang/go-bgm/backend/framework/view"
	"fmt"
)

// 首页
func IndexController(resp http.ResponseWriter, req *http.Request) {
	var (
		html []byte
		err error
	)

	if html, err = view.Render("index.html", map[string]string{"title": "gogogo"}); err != nil {
		fmt.Println(err)
		return
	}

	resp.Write(html)
}