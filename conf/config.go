package conf

import (
    "encoding/json"
    "fmt"
)

const (
    //EASHOST = "https://a1.easemob.com"
    EASHOST = ""
    APPNAME = ""
    ORGNAME = ""
    CLIENTID = ""
    CLIENTSECRET = ""
    
)

var App *Apps

func init(){
	App = NewApp(APPNAME, ORGNAME, CLIENTID, CLIENTSECRET)
}

type Apps struct {
    ClientId, ClientSecret, OrgName, AppName string
}

func NewApp(appName, orgName, clientId, clientSecret string) *Apps{
    app := new(Apps)
    app.ClientId = clientId
    app.AppName = appName
    app.OrgName = orgName
    app.ClientSecret = clientSecret
    return app
}

func (a *Apps) GetPreQuery() string {
    return fmt.Sprintf(EASHOST+"/%s/%s/", a.OrgName, a.AppName)
}

func (a *Apps) GetTokenParams() string {
    data := make(map[string]string)
    data["grant_type"] = "client_credentials"
    data["client_id"] = a.ClientId
    data["client_secret"] = a.ClientSecret
    rest, _  := json.Marshal(data)
    return string(rest)
}

