package server

import (
	"github.com/labstack/echo"
	"github.com/noriyuki-sasagawa/nyaitter_api/nyaitter"
	"github.com/noriyuki-sasagawa/nyaitter_api/twitter"
)

// RunAPIServer APIサーバー実行
func RunAPIServer() {
	e := echo.New()
	e.Router()
	e.GET("/auth", twitter.AuthTwitter)
	e.GET("/callback", twitter.Callback)
	e.POST("/post", twitter.PostTwitterAPI)
	e.POST("/replace", nyaitter.ReplaceMessge)
	e.Logger.Fatal(e.Start(":3022"))
}
