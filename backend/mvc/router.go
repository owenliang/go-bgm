package mvc

import (
	"net/http"
	"github.com/owenliang/go-bgm/backend/mvc/controller"
)

var Routes = map[string] http.HandlerFunc {
	"/": controller.IndexController,
}