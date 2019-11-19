package page

import (
    "net/http"

    "github.com/labstack/echo"
)

// Index index.html
func Index(c echo.Context) error {
    return c.Render(http.StatusOK, "index.html", nil)
}

// Tweet tweet.html
func Tweet(c echo.Context) error {
    return c.Render(http.StatusOK, "tweet.html", nil)
}
