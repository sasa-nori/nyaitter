package server

import (
    "html/template"
    "io"

    "github.com/labstack/echo"
    "github.com/noriyuki-sasagawa/nyaitter_api/nyaitter"
    "github.com/noriyuki-sasagawa/nyaitter_api/page"
    "github.com/noriyuki-sasagawa/nyaitter_api/twitter"
)

// RunAPIServer APIサーバー実行
func RunAPIServer() {
    t := &Template{
        templates: template.Must(template.ParseGlob("./public/views/*.html")),
    }
    e := echo.New()
    e.Static("/css", "public/views/css")
    e.Static("/js", "public/views/js")
    e.File("/header.png", "public/views/header.png")
    e.Renderer = t
    e.GET("/", page.Index)
    e.GET("/tweet", page.Tweet)
    e.GET("/auth", twitter.AuthTwitter)
    e.GET("/callback", twitter.Callback)
    e.POST("/check", twitter.HasCookie)
    e.POST("/post", twitter.PostTwitterAPI)
    e.POST("/replace", nyaitter.ReplaceMessge)
    e.Logger.Fatal(e.Start(":3022"))
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

// Template テンプレート
type Template struct {
    templates *template.Template
}
