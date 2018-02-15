package controllers

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
	"github.com/xaionaro-go/fwsmAPI/app/common"
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
	sha1HashBytes := sha1.Sum([]byte(password))
	sha1Hash := hex.EncodeToString(sha1HashBytes[:])
	for i := 0; true; i++ {
		cfgKey := fmt.Sprintf("user%v.login", i)
		loginCheck, ok := revel.Config.String(cfgKey)
		if !ok {
			revel.AppLog.Debug("there's no configuration option", cfgKey)
			break
		}
		if login != loginCheck {
			revel.AppLog.Debug("login check: ", login, "<>", loginCheck)
			continue
		}
		sha1HashCheck, ok := revel.Config.String(fmt.Sprintf("user%v.password_sha1", i))
		if !ok {
			revel.AppLog.Errorf("Shouldn't happened")
			continue
		}
		if sha1Hash != sha1HashCheck {
			revel.AppLog.Debugf("sha1 check failed: %v: %v %v %v %v", login, len(sha1Hash), len(sha1HashCheck), sha1Hash, sha1HashCheck)
			continue
		}

		revel.AppLog.Infof("Authed as %v", login)
		return c.sendJWT(login)
	}

	return c.error("invalid login/password")
}
