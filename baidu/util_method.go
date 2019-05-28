package baidu

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"

	"github.com/tidwall/gjson"
)

type Tieba struct {
	Id           string //贴吧ID
	Name         string //贴吧名称
	FavoType     string //贴吧？？
	LevelId      int64  //贴吧等级
	LevelName    string //贴吧等级名称
	CurScore     int64  //贴吧当前经验
	LevelupScore int64  //贴吧当前升级需要的经验
	Avatar       string //贴吧图像
	Slogan       string //贴吧描述

	AddScore  int64  //签到时增加的经验
	ErrorCode int64  //签到时的错误码
	ErrorMsg  string //签到时的错误信息
	Bbuss     string

	SignTotal int64 //累计签到天数
	SignKeep  int64 //连续签到天数
	Rank      int64 //今日本吧签到排名
}

func GetDoc(urlStr string) (string, error) {
	start := time.Now().Unix()
	request, _ := http.NewRequest("GET", urlStr, nil)
	request.Header.Add("Accept", "text/*, application/xml")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0")
	httpClient := *http.DefaultClient
	resp, err := httpClient.Do(request)
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("请求方式：Get，请求链接：%s，耗时:%d秒\n", urlStr, (time.Now().Unix() - start))
	return string(body), err
}
func GetDocWithCookies(urlStr string, cookies map[string]string) (string, error) {
	start := time.Now().Unix()
	request, _ := http.NewRequest("GET", urlStr, nil)
	request.Header.Add("Accept", "text/*, application/xml")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:47.0) Gecko/20100101 Firefox/47.0")

	cookieJar, _ := cookiejar.New(nil)
	_cookies := make([]*http.Cookie, 0)

	for k := range cookies {
		cookie := &http.Cookie{
			Name:   k,
			Value:  cookies[k],
			Path:   "/",
			Domain: ".baidu.com",
		}
		request.AddCookie(cookie)
		_cookies = append(_cookies, cookie)
	}
	URL, _ := url.Parse("http://baidu.com")
	cookieJar.SetCookies(URL, _cookies)
	httpClient := &http.Client{
		Jar: cookieJar,
	}
	resp, err := httpClient.Do(request)
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("请求方式：Get，请求链接：%s，耗时:%d秒\n", urlStr, (time.Now().Unix() - start))
	return string(body), err
}
func GetDocWithCookiesByPost(urlStr string, data, cookies map[string]string) (string, error) {
	start := time.Now().Unix()
	_data := url.Values{}
	for k := range data {
		key := k
		val := data[k]
		_data.Add(key, val)
	}
	postData := strings.NewReader(_data.Encode())
	request, _ := http.NewRequest("POST", urlStr, postData)
	request.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8")
	request.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Add("Accept-Language", "zh-CN,zh;q=0.8,en-US;q=0.5,en;q=0.3")
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 9_2 like Mac OS X) AppleWebKit/601.1 (KHTML, like Gecko) CriOS/47.0.2526.70 Mobile/13C71 Safari/601.1.46")
	request.Header.Add("Content-Length", strconv.Itoa(len(_data.Encode())))
	request.URL.RawQuery = _data.Encode()
	cookieJar, _ := cookiejar.New(nil)
	_cookies := make([]*http.Cookie, 0)

	for k := range cookies {
		cookie := &http.Cookie{
			Name:   k,
			Value:  cookies[k],
			Path:   "/",
			Domain: ".baidu.com",
		}
		request.AddCookie(cookie)
		_cookies = append(_cookies, cookie)
	}
	URL, _ := url.Parse("http://baidu.com")
	cookieJar.SetCookies(URL, _cookies)
	httpClient := &http.Client{
		Jar: cookieJar,
	}
	resp, err := httpClient.Do(request)
	body, err := ioutil.ReadAll(resp.Body)
	log.Printf("请求方式：Post，请求链接：%s，耗时:%d秒\n", urlStr, (time.Now().Unix() - start))
	return string(body), err
}

//获取userid
func GetUserID(userName string) string {
	urlStr := "http://tieba.baidu.com/home/get/panel?ie=utf-8&un=" + userName
	body, _ := GetDoc(urlStr)
	data := gjson.Get(body, "data")
	id := data.Get("id")
	return id.String()
}
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
func GetUserName(bduss string) string {

	baiduId := Md5(strconv.FormatInt(time.Now().Unix(), 10))
	cookies := map[string]string{}
	cookies["BAIDUID"] = baiduId
	cookies["BDUSS"] = bduss
	body, _ := GetDocWithCookies("http://wapp.baidu.com/", cookies)
	doc, _ := goquery.NewDocumentFromReader(bytes.NewReader([]byte(body)))
	b := doc.Find(".b")
	a := b.Find("a").Last()
	href, _ := a.Attr("href")
	strs := strings.Split(href, "un=")
	userName := strs[1]
	userName, _ = url.QueryUnescape(userName)
	return userName
}
func GetTiebaList(userId, bduss string) []Tieba {
	urlStr := "http://c.tieba.baidu.com/c/f/forum/like"
	data := map[string]string{}
	data["_client_id"] = fmt.Sprintf("wappc_%s_258", strconv.FormatInt(time.Now().Unix(), 10))
	data["_client_type"] = "2"
	data["_client_version"] = "6.5.8"
	data["_phone_imei"] = "357143042411618"
	data["from"] = "baidu_appstore"
	data["is_guest"] = "1"
	data["model"] = "H60-L01"
	data["page_no"] = "0"
	data["page_size"] = "200"
	data["timestamp"] = fmt.Sprintf("%s903", strconv.FormatInt(time.Now().Unix(), 10))
	data["uid"] = userId

	keys := []string{"_client_id", "_client_type", "_client_version", "_phone_imei", "from", "is_guest", "model", "page_no", "page_size", "timestamp", "uid"}

	sb := ""
	for _, k := range keys {
		sb += k + "=" + data[k]
	}
	sign := Md5(sb + "tiebaclient!!!")
	data["sign"] = sign

	cookies := map[string]string{}
	cookies["BDUSS"] = bduss

	body, _ := GetDocWithCookiesByPost(urlStr, data, cookies)

	tiebaList := []Tieba{}

	forum_list := gjson.Get(body, "forum_list")
	non_gconforum := forum_list.Get("non-gconforum")

	for _, r := range non_gconforum.Array() {

		tiebaList = append(tiebaList, Tieba{
			Id:           r.Get("id").String(),
			Name:         r.Get("name").String(),
			FavoType:     r.Get("favo_type").String(),
			LevelId:      r.Get("level_id").Int(),
			LevelName:    r.Get("level_name").String(),
			CurScore:     r.Get("cur_score").Int(),
			LevelupScore: r.Get("levelup_score").Int(),
			Avatar:       r.Get("avatar").String(),
			Slogan:       r.Get("slogan").String(),
		})
	}
	gconforum := forum_list.Get("gconforum")

	for _, r := range gconforum.Array() {

		tiebaList = append(tiebaList, Tieba{
			Id:           r.Get("id").String(),
			Name:         r.Get("name").String(),
			FavoType:     r.Get("favo_type").String(),
			LevelId:      r.Get("level_id").Int(),
			LevelName:    r.Get("level_name").String(),
			CurScore:     r.Get("cur_score").Int(),
			LevelupScore: r.Get("levelup_score").Int(),
			Avatar:       r.Get("avatar").String(),
			Slogan:       r.Get("slogan").String(),
		})
	}

	return tiebaList
}
func GetTBS(bduss string) string {
	cookies := map[string]string{}
	cookies["BDUSS"] = bduss
	body, _ := GetDocWithCookies("http://tieba.baidu.com/dc/common/tbs", cookies)
	tbs := gjson.Get(body, "tbs")
	return tbs.String()
}
func Sign(tieba Tieba, bduss string) Tieba {
	urlStr := "http://c.tieba.baidu.com/c/c/forum/sign"
	cookies := map[string]string{}
	cookies["BDUSS"] = bduss

	data := map[string]string{}
	data["BDUSS"] = bduss
	data["_client_id"] = "03-00-DA-59-05-00-72-96-06-00-01-00-04-00-4C-43-01-00-34-F4-02-00-BC-25-09-00-4E-36"
	data["_client_type"] = "4"
	data["_client_version"] = "1.2.1.17"
	data["_phone_imei"] = "540b43b59d21b7a4824e1fd31b08e9a6"
	data["fid"] = "1"
	data["kw"] = tieba.Name
	data["net_type"] = "3"
	data["tbs"] = GetTBS(bduss)

	keys := []string{"BDUSS", "_client_id", "_client_type", "_client_version", "_phone_imei", "fid", "kw", "net_type", "tbs"}

	sb := ""
	for _, k := range keys {
		sb += k + "=" + data[k]
	}
	sign := Md5(sb + "tiebaclient!!!")
	data["sign"] = sign

	body, _ := GetDocWithCookiesByPost(urlStr, data, cookies)

	tieba.ErrorCode = gjson.Get(body, "error_code").Int()
	tieba.ErrorMsg = gjson.Get(body, "error_msg").String()
	tieba.AddScore = gjson.Get(body, "user_info").Get("sign_bonus_point").Int()

	return tieba
}
func AfterSign(tieba Tieba, bduss string) Tieba {
	urlStr := "http://tieba.baidu.com/sign/loadmonth?kw=" + tieba.Name + "&ie=utf-8&t=0.8314427338417005"
	cookies := map[string]string{}
	cookies["BDUSS"] = bduss

	body, _ := GetDocWithCookies(urlStr, cookies)
	tieba.Rank = gjson.Get(body, "data").Get("sign_user_info").Get("rank").Int()
	tieba.SignTotal = gjson.Get(body, "data").Get("sign_user_info").Get("sign_total").Int()
	tieba.SignKeep = gjson.Get(body, "data").Get("sign_user_info").Get("sign_keep").Int()

	return tieba
}
