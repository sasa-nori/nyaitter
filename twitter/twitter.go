package twitter

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
    "time"

    "github.com/ChimeraCoder/anaconda"
    "github.com/garyburd/go-oauth/oauth"
    session "github.com/ipfans/echo-session"

    "github.com/labstack/echo"
)

const callback = "https://cat.newstyleservice.net/callback"
const test = "http://localhost:3022/callback"

// AuthTwitter ツイッターの認証開始
func AuthTwitter(c echo.Context) error {
    api := connectAPI()
    var url = callback
    hostname, _ := os.Hostname()
    if strings.Contains(hostname, "local") {
        url = test
    }
    uri, _, error := api.AuthorizationURL(url)
    if error != nil {
        fmt.Println(error)
        return error
    }

    return c.Redirect(http.StatusFound, uri)
}

// Callback ログイン後のコールバックから認証まで
func Callback(c echo.Context) error {
    token := c.QueryParam("oauth_token")
    secret := c.QueryParam("oauth_verifier")
    api := connectAPI()

    cred, _, error := api.GetCredentials(&oauth.Credentials{
        Token: token,
    }, secret)
    if error != nil {
        fmt.Println(error)
        return error
    }
    api = anaconda.NewTwitterApi(cred.Token, cred.Secret)

    sess := session.Default(c)
    sess.Set("token", cred.Token)
    sess.Set("secret", cred.Secret)
    sess.Save()

    return c.Redirect(http.StatusFound, "./tweet")
}

// PostTwitterAPI ツイッター投稿
func PostTwitterAPI(c echo.Context) error {
    input := c.FormValue("input")
    writeCookie(c, "message", input)
    sess := session.Default(c)

    token := sess.Get("token")
    secret := sess.Get("secret")
    if token == nil || secret == nil {
        c.Redirect(http.StatusFound, "./tweet")
    }
    api := anaconda.NewTwitterApi(token.(string), secret.(string))

    message := c.FormValue("message")
    tweet, error := api.PostTweet(message, nil)
    if error != nil {
        fmt.Println(error)
        return c.JSON(http.StatusAccepted, "redirect")
    }
    link := "https://twitter.com/" + tweet.User.IdStr + "/status/" + tweet.IdStr
    clearCookie(c, "message")
    return c.JSON(http.StatusOK, link)
}

// HasCookie クッキーあるかどうか確認
func HasCookie(c echo.Context) error {
    session := session.Default(c)
    token := session.Get("token")
    if token == nil {
        return c.JSON(http.StatusNoContent, "no")
    }

    return c.JSON(http.StatusOK, "has data")
}

func connectAPI() *anaconda.TwitterApi {
    // Json読み込み
    raw, error := ioutil.ReadFile("./path/to/twitterAccount.json")
    if error != nil {
        fmt.Println(error.Error())
        return nil
    }

    var twitterAccount Account
    // 構造体にセット
    json.Unmarshal(raw, &twitterAccount)

    anaconda.SetConsumerKey(twitterAccount.ConsumerKey)
    anaconda.SetConsumerSecret(twitterAccount.ConsumerSecret)

    // 認証
    return anaconda.NewTwitterApi("", "")
}

func writeCookie(c echo.Context, name string, value string) {
    cookie := new(http.Cookie)
    cookie.Name = name
    cookie.Value = value
    cookie.Expires = time.Now().Add(24 * time.Hour)
    c.SetCookie(cookie)
}

func clearCookie(c echo.Context, name string) {
    cookie := new(http.Cookie)
    cookie.Name = name
    cookie.Value = ""
    c.SetCookie(cookie)
}

// Account はTwitterの認証用の情報
type Account struct {
    AccessToken       string `json:"accessToken"`
    AccessTokenSecret string `json:"accessTokenSecret"`
    ConsumerKey       string `json:"consumerKey"`
    ConsumerSecret    string `json:"consumerSecret"`
}
