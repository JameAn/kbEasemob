package kbEasemob

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"kbEasemob/conf"
	"net/http"
	"strings"
	"sync"
	"time"
)

type Token struct {
	AccessToken string    `json:"access_token"`
	ExpiresIn int         `json:"expires_in"`
}

var T *Token

var mutex sync.RWMutex


func init() {
	if T == nil {
		T = new(Token)
	}
}

func (t *Token) GetTokenUrl() string {
    return conf.App.GetPreQuery() + "token"
}

func (t *Token) InitT(){
	mutex.Lock()
	defer mutex.Unlock()
	token := httpRequestToken(t.GetTokenUrl(), conf.App.GetTokenParams())
	t.AccessToken = token.AccessToken
	t.ExpiresIn = token.ExpiresIn
}

func (t *Token) Empty() bool {
	mutex.RLock()
	defer mutex.RUnlock()
	if t.AccessToken == "" && t.ExpiresIn == 0 {
		return true
	}
	return false
}

// token是否过期
func (t *Token) IsExpire() bool {
	mutex.RLock()
	defer mutex.RUnlock()
	if time.Now().After(time.Unix(int64(t.ExpiresIn), 0)) {
		return true
	}
	return false
}

// 更新token
func (t *Token) Update() {
	mutex.Lock()
	defer mutex.Unlock()
	token := httpRequestToken(t.GetTokenUrl(), conf.App.GetTokenParams())
	t.AccessToken = token.AccessToken
	t.ExpiresIn = token.ExpiresIn
}


func GetT() *Token {
	if T.Empty() {
		T.InitT()
	} else if T.IsExpire() {
		T.Update()
	}
	return T
}

func httpRequestToken(url string, data interface{}) *Token {
	result := new(Token)
	contentType := "application/json"
	resp, err := http.Post(url, contentType, strings.NewReader(data.(string)))
    defer resp.Body.Close()
	if err != nil {
	    fmt.Println(err)
	}
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(url, "\n", string(body), "\n")
	json.Unmarshal(body, result)
	return result
}

