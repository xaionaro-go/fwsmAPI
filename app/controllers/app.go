package controllers

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
	"github.com/xaionaro-go/mswfAPI/app"
	"github.com/xaionaro-go/mswfAPI/app/common"
	"strings"
)

type App struct {
	Controller
}
type routeInfo struct {
	Method string
	Path   string
}

func (c App) Index() revel.Result {
	var routes []routeInfo
	for _, route := range revel.MainRouter.Routes {
		if len(route.ControllerNamespace) < 1 {
			continue
		}

		if route.ControllerNamespace[0:1] != strings.ToUpper(route.ControllerNamespace[0:1]) {
			continue
		}

		routes = append(routes, routeInfo{
			Path:   route.Path,
			Method: route.Method,
		})
	}

	return c.render(routes)
}

func (c App) sendJWT(login string) revel.Result {
	signingSecret, ok := revel.Config.String("jwt_secret")
	if !ok {
		revel.AppLog.Errorf("Shouldn't happened")
		return c.error("Internal error")
	}

	c.Session["login"] = login
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": common.UserInfo{Username: login, CanRead: true, CanWrite: true},
	})
	tokenString, err := token.SignedString([]byte(signingSecret))
	if err != nil {
		revel.AppLog.Errorf("Got error: %v %v", err.Error())
		return c.error("Internal error")
	}

	return c.render(map[string]string{"token": tokenString})
}

func (c App) AuthJWT() revel.Result {
	var jsonData map[string]interface{}
	c.Params.BindJSON(&jsonData)
	login, _ := jsonData["login"].(string)
	password, _ := jsonData["password"].(string)

	if login == "" {
		return c.error("empty \"login\" is passed")
	}

	if app.CheckLoginPass(login, password) {
		return c.sendJWT(login)
	}
	return c.error("invalid login/password")
}
