package kbEasemob

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"kbEasemob/conf"
	"net/http"
	"strings"
	"sync"
)

var registerStore sync.Map

type ImUser struct {
	UserName    string      `json:"username"`
	Password    string      `json:"password"`
}

func NewImUser(userName, password string) *ImUser {
	iu := new(ImUser)
	iu.UserName = userName
	iu.Password = password
	return iu
}

// 判断是否已经注册
func (iu *ImUser) Exist() bool {
	 _, ok := registerStore.Load(iu.UserName)
	 if !ok {
	 	if _, err := iu.GetUser(iu.UserName); err != nil {
	 		// not exist
	 		return false
	    } else {
	    	// exist: reg mem
	    	registerStore.Store(iu.UserName, true)
	    	return true
	    }
	 }
	 return ok
}

func (iu *ImUser) getQueryUrl() string {
	return conf.App.GetPreQuery() + "users"
}

// 注册单个用户
func (iu *ImUser) Register() error {
	if iu.Exist() {
	    return nil
	}
    jsonData, _ := json.Marshal(iu)
    data := string(jsonData)
    client := http.Client{}
    req, _ := http.NewRequest("POST", iu.getQueryUrl(), strings.NewReader(data))
    req.Header.Add("Content-Type", "application/json")
	tokenStr := "Bearer " + GetT().AccessToken
	req.Header.Add("Authorization", tokenStr)
    response, _ := client.Do(req)
    defer response.Body.Close()
	body, _ := ioutil.ReadAll(response.Body)
    resCode := response.StatusCode
    if resCode != 200 {
    	return errors.New(string(body))
    }
    // reg mem
    registerStore.Store(iu.UserName, true)
    return nil
}

// 获取用户信息
func (iu *ImUser) GetUser(userName string) (map[string]interface{}, error) {
    url := iu.getQueryUrl() + "/" + userName
    client := http.Client{}
    req, _ := http.NewRequest("GET", url, nil)
    tokenStr := "Bearer " + GetT().AccessToken
	req.Header.Add("Authorization",tokenStr)
    response, _ := client.Do(req)
    defer response.Body.Close()
    body, _ := ioutil.ReadAll(response.Body)
    resCode := response.StatusCode
    if resCode != 200 {
    	return nil, errors.New(string(body))
    }
    result := make(map[string]interface{})
    json.Unmarshal(body, result)
    return result, nil
}
