package view

import (
	"path/filepath"
	"os"
	"html/template"
	"bytes"
	"github.com/owenliang/go-bgm/backend/framework/config"
)

var (
	tplSet *template.Template
)

// 加载所有模板文件
func loadViews(dir string) (tplSet *template.Template, err error) {
	var (
		viewTpls = make([]string, 0)
		walkFunc filepath.WalkFunc
	)

	walkFunc = func(path string, info os.FileInfo, err error) (retErr error) {
		if err == nil && !info.IsDir() {
			ext := filepath.Ext(path)
			if ext == ".html" {
				viewTpls = append(viewTpls, path)
			}
		}
		return
	}

	if err = filepath.Walk(dir, walkFunc); err != nil {
		return
	}

	if tplSet, err = template.ParseFiles(viewTpls...); err != nil {
		return
	}
	return
}

// 初始化
func InitView() (err error) {
	if tplSet, err = loadViews(config.G_Conf.TemplateDir); err != nil {
		return
	}
	return
}

// 渲染模板
func Render(name string, data interface{}) (content []byte, err error) {
	var (
		buf *bytes.Buffer
	)

	buf = bytes.NewBuffer(nil)
	if err = tplSet.ExecuteTemplate(buf, name, data); err != nil {
		return
	}
	content = buf.Bytes()
	return
}
