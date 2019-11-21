package nyaitter

import (
    "net/http"
    "strings"

    "github.com/labstack/echo"
    "golang.org/x/exp/utf8string"
)

var keyword = map[string]string{
    "oh":     "ｵｵﾝ",
    "oh...":  "ｵｵｵｵｵｵﾝ",
    "ﾊﾞｷｭｰﾝ": "ﾆｬｵｰﾝ",
    "わおーん":   "ﾆｬｵｰﾝ",
    "うま言う":   "ちょw誰が上手いこと言えって言ったにゃww",
    "まだ":     "まだにゃ",
    "した":     "したにゃ",
    "った":     "ったにゃ",
    "です":     "ですにゃ",
    "よう":     "ようにゃ",
    "IT藤原猫":  "は゛ぁ゛あ゛ぁ゛ぁ゛ぁ゛猫゛か゛わ゛い゛い゛に゛ゃ゛ぁ゛\n#IT藤原猫\n",
    "猫":      "にゃーん (=･ω･=)",
    "な":      "にゃ",
}

// ReplaceMessge 文字列置換
func ReplaceMessge(c echo.Context) error {
    message := c.FormValue("message")
    utf8Message := utf8string.NewString(message)
    size := utf8Message.RuneCount()
    if size > 120 {
        return c.JSON(http.StatusBadRequest, "string length over")
    }
    for key, value := range keyword {
        message = strings.ReplaceAll(message, key, value)
    }

    if strings.HasSuffix(message, "にゃ") {
        message += "ん"
    }

    message += "\n#にゃイッター"

    if utf8string.NewString(message).RuneCount() >= 140 {
        return c.JSON(http.StatusBadRequest, "string length over")
    }

    return c.JSON(http.StatusOK, message)
}
