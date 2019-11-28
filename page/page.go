package page

import (
    "net/http"

    session "github.com/ipfans/echo-session"
    "github.com/labstack/echo"
)

// Index index.html
func Index(c echo.Context) error {
    return c.Render(http.StatusOK, "index.html", nil)
}

// Tweet tweet.html
func Tweet(c echo.Context) error {
    session := session.Default(c)
    token := session.Get("token")
    if token == nil {
        return c.Redirect(http.StatusFound, "./")
    }
    preData := new(PreData)
    preData.Tweet = readCookie(c, "message")
    preData.Reply = readCookie(c, "reply")

    return c.Render(http.StatusOK, "tweet.html", preData)
}

func readCookie(c echo.Context, name string) string {
    cookie, error := c.Cookie(name)
    if error != nil {
        return ""
    }
    return cookie.Value
}

// PreData 過去の情報
type PreData struct {
    Tweet string
    Reply string
}
