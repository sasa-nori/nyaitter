package twitter

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/labstack/echo"
)

// Timeline タイムライン取得
func Timeline(c echo.Context) error {
    api := connectAuth()
    api.EnableThrottling(1*time.Second, 4)
    v := url.Values{}
    v.Set("count", "100")
    result, _ := api.GetSearch(`#にゃイッター`, v)
    return c.JSON(http.StatusOK, result)
}

func GetTimeline() (sr anaconda.SearchResponse) {
    api := connectAuth()
    v := url.Values{}
    v.Set("count", "10")
    // 検索
    sr, _ = api.GetSearch(`#にゃイッター`, v)
    return sr
}

func connectAuth() *anaconda.TwitterApi {
    // Json読み込み
    raw, error := ioutil.ReadFile("./path/to/twitterAccount.json")
    if error != nil {
        fmt.Println(error.Error())
        return nil
    }

    var twitterAccount Account
    // 構造体にセット
    json.Unmarshal(raw, &twitterAccount)

    // 認証
    return anaconda.NewTwitterApiWithCredentials(twitterAccount.AccessToken, twitterAccount.AccessTokenSecret, twitterAccount.ConsumerKey, twitterAccount.ConsumerSecret)
}
