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
    return c.Render(http.StatusOK, "tweet.html", nil)
}
