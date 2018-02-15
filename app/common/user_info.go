package common

import (
	"github.com/mitchellh/mapstructure"
)

type UserInfo struct {
	Username string
	CanRead  bool
	CanWrite bool
}


func UserInfoFromClaimsUser(claimsUser map[string]interface{}) (result UserInfo) {
	mapstructure.Decode(claimsUser, &result)
	return
}
