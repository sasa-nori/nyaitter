package twitter

import (
	"encoding/json"
	"strings"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/garyburd/go-oauth/oauth"
	"github.com/labstack/echo"
)

const callback = "https://cat.newstyleservice.net/callback"
const test = "http://localhost:3022/callback"

// AuthTwitter ツイッターの認証開始
func AuthTwitter(c echo.Context) error {
	api := connectAPI()
	uri, _, error := api.AuthorizationURL(callback)
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

	cookie := new(http.Cookie)
	cookie.Name = "twitter"
	cookie.Value = cred.Token + "," + cred.Secret
	cookie.Expires = time.Now().Add(24 * 7 * time.Hour)
	c.SetCookie(cookie)

    // TODO: 2019/11/15 投稿フォームのページへリダイレクト
	return c.JSON(http.StatusOK, "success")
}

// PostTwitterAPI ツイッター投稿
func PostTwitterAPI(c echo.Context) error{
    cookie, _ := c.Cookie("twitter")
    if cookie == nil {
        return AuthTwitter(c)
    }
    token := strings.Split(cookie.Value, ",")[0]
    secret := strings.Split(cookie.Value, ",")[1]
    
    api := anaconda.NewTwitterApi(token, secret)

    message := c.QueryParam("message") + "\n #にゃイッター"
    tweet, error := api.PostTweet(message ,nil)
    if error != nil {
        fmt.Println(error)
        return error
    }

    return c.JSON(http.StatusOK, tweet.FullText)
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

// Account はTwitterの認証用の情報
type Account struct {
	AccessToken       string `json:"accessToken"`
	AccessTokenSecret string `json:"accessTokenSecret"`
	ConsumerKey       string `json:"consumerKey"`
	ConsumerSecret    string `json:"consumerSecret"`
}
