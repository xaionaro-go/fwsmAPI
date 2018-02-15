package app

import (
	"bufio"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/revel/revel"
	"github.com/xaionaro-go/fwsmAPI/app/common"
	"github.com/xaionaro-go/fwsmConfig"
	"strings"
	"os"
)

const (
	FWSM_CONFIG_PATH = "/root/fwsm-config/dynamic"
)

var (
	// AppVersion revel app version (ldflags)
	AppVersion string

	// BuildTime revel app build-time (ldflags)
	BuildTime string

	FWSMConfig fwsmConfig.FwsmConfig
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadConfig() {
	file, err := os.Open(FWSM_CONFIG_PATH)
	checkErr(err)
	defer file.Close()
	cfgReader := bufio.NewReader(file)
	FWSMConfig, err = fwsmConfig.Parse(cfgReader)
	checkErr(err)
}

func init() {
	// Filters is the default set of global filters.
	revel.Filters = []revel.Filter{
		revel.PanicFilter,             // Recover from panics and display an error page instead.
		revel.RouterFilter,            // Use the routing table to select the right Action
		revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
		revel.ParamsFilter,            // Parse parameters into Controller.Params.
		revel.SessionFilter,           // Restore and write the session cookie.
		revel.FlashFilter,             // Restore and write the flash cookie.
		revel.ValidationFilter,        // Restore kept validation errors and save new ones from cookie.
		revel.I18nFilter,              // Resolve the requested language
		HeaderFilter,                  // Add some security based headers
		revel.InterceptorFilter,       // Run interceptors around the action.
		revel.CompressFilter,          // Compress the result.
		ActionInvoker,                 // Invoke the action.
	}

	// Register startup functions with OnAppStart
	// revel.DevMode and revel.RunMode only work inside of OnAppStart. See Example Startup Script
	// ( order dependent )
	// revel.OnAppStart(ExampleStartupScript)
	// revel.OnAppStart(InitDB)
	revel.OnAppStart(ReadConfig)
}

// HeaderFilter adds common security headers
// There is a full implementation of a CSRF filter in
// https://github.com/revel/modules/tree/master/csrf
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
	c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
	c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
	c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")
	c.Response.Out.Header().Add("Referrer-Policy", "strict-origin-when-cross-origin")

	fc[0](c, fc[1:]) // Execute the next filter stage.
}

func claimsUserToUserInfo(claimsUser map[string]interface{}) common.UserInfo {
	return common.UserInfoFromClaimsUser(claimsUser)
}

func tryParseJWT(c *revel.Controller) {
	var tokenString string
	c.Params.Bind(&tokenString, "token")
	if tokenString == "" {
		tokenString = c.Request.GetHttpHeader("Authorization")
		revel.AppLog.Debugf("tokenString from the header == <%v>", tokenString)
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			return // No authorization information is passed
		}
		tokenString = strings.Split(tokenString, " ")[1]

	}
	if tokenString == "" {
		revel.AppLog.Errorf("tokenString is empty")
		return // this shouldn't happened, see above
	}
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		jwtSecret, ok := revel.Config.String("jwt_secret")
		if !ok {
			revel.AppLog.Errorf("Shouldn't happened")
			panic("Internal error")
		}
		return []byte(jwtSecret), nil
	})

	if err != nil {
		revel.AppLog.Errorf("Got error: %v; token:<%v>", err.Error(), tokenString)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		c.ViewArgs["me"] = claimsUserToUserInfo(claims["user"].(map[string]interface{}))
	}

	return
}

var ActionInvoker = func(c *revel.Controller, f []revel.Filter) {
	c.ViewArgs["me"] = common.UserInfo{}
	tryParseJWT(c)

	revel.ActionInvoker(c, f)
}

//func ExampleStartupScript() {
//	// revel.DevMod and revel.RunMode work here
//	// Use this script to check for dev mode and set dev/prod startup scripts here!
//	if revel.DevMode == true {
//		// Dev mode
//	}
//}
