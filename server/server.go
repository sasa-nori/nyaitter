package server

import (
    "context"
    "html/template"
    "io"
    "os"
    "os/signal"

    "time"

    session "github.com/ipfans/echo-session"
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
    e.File("/twitter-card.png", "public/views/twitter-card.png")
    e.Renderer = t
    //セッションを設定
    store := session.NewCookieStore([]byte("secret-key"))
    //セッション保持時間
    store.MaxAge(86400)
    e.Use(session.Sessions("ESESSION", store))

    e.GET("/", page.Index)
    e.GET("/tweet", page.Tweet)
    e.GET("/auth", twitter.AuthTwitter)
    e.GET("/callback", twitter.Callback)
    e.POST("/check", twitter.HasSessionData)
    e.POST("/post", twitter.PostTwitterAPI)
    e.POST("/replace", nyaitter.ReplaceMessge)
    // サーバーを開始
    go func() {
        if err := e.Start(":3022"); err != nil {
            e.Logger.Info("shutting down the server")
        }
    }()

    // (Graceful Shutdown)
    quit := make(chan os.Signal)
    signal.Notify(quit, os.Interrupt)
    <-quit
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

    defer cancel()
    store.MaxAge(-1)
    if err := e.Shutdown(ctx); err != nil {
        e.Logger.Fatal(err)
    }
}

// Render テンプレートレンダリング
func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
    return t.templates.ExecuteTemplate(w, name, data)
}

// Template テンプレート
type Template struct {
    templates *template.Template
}
