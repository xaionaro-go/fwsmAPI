package controllers

// This helper is a hack to bypass this issue: https://github.com/revel/revel/issues/1239

import (
	"fmt"
	"github.com/revel/revel"
	"github.com/xaionaro-go/fwsmAPI/app/common"
)

type Controller struct {
	*revel.Controller
}

func (c Controller) error(v interface{}) revel.Result {
	if revel.DevMode {
		revel.AppLog.Errorf("Got error: %v", v)
	}

	c.ViewArgs["status"] = "ERROR"
	c.ViewArgs["error_description"] = fmt.Sprintf("%v", v)

	return c.RenderJSON(c.ViewArgs)
}
func (c Controller) noPerm() revel.Result {
	c.ViewArgs["status"] = "ERROR"
	c.ViewArgs["error_description"] = "Permission denied"

	return c.RenderJSON(c.ViewArgs)
}
func (c Controller) notFound() revel.Result {
	c.ViewArgs["status"] = "ERROR"
	c.ViewArgs["error_description"] = "Not found"

	return c.RenderJSON(c.ViewArgs)
}
func (c Controller) notImplemented() revel.Result {
	c.ViewArgs["status"] = "ERROR"
	c.ViewArgs["error_description"] = "Not implemented, yet"

	return c.RenderJSON(c.ViewArgs)
}
func (c Controller) invalidArgs() revel.Result {
	c.ViewArgs["status"] = "ERROR"
	c.ViewArgs["error_description"] = "Invalid arguments"

	return c.RenderJSON(c.ViewArgs)
}

func (c *Controller) Redirect(url string) revel.Result {
	return c.Controller.Redirect(url)
}

func (c *Controller) render(arg interface{}) revel.Result {
	c.ViewArgs["status"] = "OK"
	c.ViewArgs["result"] = arg

	return c.RenderJSON(c.ViewArgs)
}

func (c Controller) GetMe() common.UserInfo {
	return c.ViewArgs["me"].(common.UserInfo)
}
func (c Controller) IsCanRead() bool {
	return c.GetMe().CanRead
}
func (c Controller) IsCanWrite() bool {
	return c.GetMe().CanWrite
}
