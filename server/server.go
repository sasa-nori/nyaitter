package server

import (
    "html/template"
    "io"

    "time"
    "github.com/tylerb/graceful"
    session "github.com/ipfans/echo-session"
    "github.com/labstack/echo"
    "github.com/sasa-nori/nyaitter/nyaitter"
    "github.com/sasa-nori/nyaitter/page"
    "github.com/sasa-nori/nyaitter/twitter"
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
    e.File("/twitter-card.png", "public/views/twitter-card.png")
    e.Renderer = t
    //セッションを設定
    store := session.NewCookieStore([]byte("secret-key"))
    //セッション保持時間 1ヶ月 2592000, 1日 86400
    store.MaxAge(2592000)
    e.Use(session.Sessions("ESESSION", store))

    e.GET("/", page.Index)
    e.GET("/tweet", page.Tweet)
    e.GET("/auth", twitter.AuthTwitter)
    e.GET("/timeline", twitter.Timeline)
    e.GET("/callback", twitter.Callback)
    e.POST("/check", twitter.HasSessionData)
    e.POST("/post", twitter.PostTwitterAPI)
    e.POST("/replace", nyaitter.ReplaceMessge)
    e.GET("/logout", page.Logout)
    // サーバーを開始
    e.Server.Addr = ":2222"

    // Serve it like a boss
    graceful.ListenAndServe(e.Server, 5*time.Second)
}

// Render テンプレートレンダリング
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

// Template テンプレート
type Template struct {
    templates *template.Template
}
