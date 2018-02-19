package controllers

import (
	"bytes"
	"encoding/csv"
	"github.com/revel/revel"
	"github.com/xaionaro-go/fwsmAPI/app"
	"github.com/xaionaro-go/fwsmAPI/app/helpers"
	"github.com/xaionaro-go/fwsmConfig"
	"strings"
	"os/exec"
)

type FWSM struct {
	Controller
}

func (c FWSM) GetConfiguration() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}

	return c.render(app.FWSMConfig)
}

// VLAN

func (c FWSM) getVLAN() (fwsmConfig.VLAN, bool) {
	var vlanId int
	c.Params.Bind(&vlanId, "vlan")

	return app.FWSMConfig.VLANs.Find(vlanId)
}

func (c FWSM) GetVLAN() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}

	vlan, ok := c.getVLAN()
	if ok {
		return c.render(vlan)
	}
	return c.notFound()
}

func (c FWSM) DeleteVLANs() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	var vlanIdsStr string
	c.Params.Bind(&vlanIdsStr, "vlans")
	vlanIdStrs := strings.Split(vlanIdsStr,  ",")

	vlanIdsI, err := helpers.Atoi(vlanIdStrs)
	if err != nil {
		revel.AppLog.Errorf("Got error: ", err.Error)
		return c.error("Internal error")
	}
	vlanIds := vlanIdsI.([]int)

	if len(vlanIds) == 0 {
		revel.AppLog.Errorf("no vlans selected")
		return c.invalidArgs()
	}

	err = app.FWSMConfig.VLANs.Remove(app.NetworkHosts, vlanIds...)
	if err != nil {
		revel.AppLog.Errorf("Got an error: %v", err.Error())
		return c.error("Got an error while communicating with network hosts")
	}

	return c.GetVLANs()
}

func (c FWSM) UpdateVLAN() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) CreateVLAN() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) GetVLANs() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}

	return c.render(app.FWSMConfig.VLANs)
}

// DHCP

func (c FWSM) getDHCP() (fwsmConfig.DHCP, bool) {
	return fwsmConfig.DHCP{}, false
}

func (c FWSM) GetDHCP() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) DeleteDHCPs() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) UpdateDHCP() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) CreateDHCP() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) GetDHCPs() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}

	return c.render(app.FWSMConfig.DHCPs)
}

// SNAT

func (c FWSM) getSNAT() (fwsmConfig.SNAT, bool) {
	return fwsmConfig.SNAT{}, false
}

func (c FWSM) GetSNAT() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) DeleteSNATs() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) UpdateSNAT() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) CreateSNAT() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) GetSNATs() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}

	return c.render(app.FWSMConfig.SNATs)
}

// DNAT

func (c FWSM) getDNAT() (fwsmConfig.DNAT, bool) {
	return fwsmConfig.DNAT{}, false
}

func (c FWSM) GetDNAT() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) DeleteDNATs() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) UpdateDNAT() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) CreateDNAT() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) GetDNATs() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}

	return c.render(app.FWSMConfig.DNATs)
}

// Route

func (c FWSM) getRoute() (fwsmConfig.Route, bool) {
	return fwsmConfig.Route{}, false
}

func (c FWSM) GetRoute() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) DeleteRoutes() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) UpdateRoute() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) CreateRoute() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) GetRoutes() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}

	return c.render(app.FWSMConfig.Routes)
}

// ACL

func (c FWSM) getACL() (fwsmConfig.ACL, bool) {
	return fwsmConfig.ACL{}, false
}

func (c FWSM) GetACL() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) DeleteACLs() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) UpdateACL() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) CreateACL() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	return c.notImplemented()
}

func (c FWSM) GetACLs() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}

	return c.render(app.FWSMConfig.ACLs)
}

// The rest

type bwmRow struct {
	TS              string
	Iface           string
	BytesOutPS      string
	BytesInPS       string
	BytesTotalPS    string
	BytesOutTotal   string
	BytesInTotal    string
	PacketsOutPS    string
	PacketsInPS     string
	PacketsTotalPS  string
	PacketsOutTotal string
	PacketsInTotal  string
	ErrorsOutPS     string
	ErrorsInPS      string
	ErrorsOutTotal  string
	ErrorsInTotal   string
}

type bwmRows []bwmRow

func (c FWSM) GetStatus() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}
	cmd := exec.Command("bwm-ng", "-o", "csv", "-c", "1", "-t", "5000")
	var cmdOut bytes.Buffer
	cmd.Stdout = &cmdOut
	err := cmd.Run()
	if err != nil {
		return c.error(err.Error())
	}

	r := csv.NewReader(&cmdOut)
	r.Comma = ';'

	lines, err := r.ReadAll()
	if err != nil {
		return c.error(err.Error())
	}

	bwmRows := bwmRows{}
	for _, line := range lines {
		bwmRows = append(bwmRows, bwmRow{
			TS:              line[0],
			Iface:           line[1],
			BytesOutPS:      line[2],
			BytesInPS:       line[3],
			BytesTotalPS:    line[4],
			BytesOutTotal:   line[5],
			BytesInTotal:    line[6],
			PacketsOutPS:    line[7],
			PacketsInPS:     line[8],
			PacketsTotalPS:  line[9],
			PacketsOutTotal: line[10],
			PacketsInTotal:  line[11],
			ErrorsOutPS:     line[12],
			ErrorsInPS:      line[13],
			ErrorsOutTotal:  line[14],
			ErrorsInTotal:   line[15],
		})
	}

	return c.render(bwmRows)
}

func (c FWSM) Apply() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	err := app.FWSMConfig.Apply(app.NetworkHosts)
	if err != nil {
		revel.AppLog.Errorf("Got an error: %v", err.Error())
		return c.render("Internal error")
	}
	return c.render(app.FWSMConfig)
}

func (c FWSM) Revert() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	err := app.FWSMConfig.Revert(app.NetworkHosts)
	if err != nil {
		revel.AppLog.Errorf("Got an error: %v", err.Error())
		return c.render("Internal error")
	}
	return c.render(app.FWSMConfig)
}

func (c FWSM) Save() revel.Result {
	if !c.IsCanWrite() {
		return c.noPerm()
	}

	err := app.FWSMConfig.Save(app.NetworkHosts, "/root/fwsm-config/dynamic")
	if err != nil {
		revel.AppLog.Errorf("Got an error: %v", err.Error())
		return c.render("Internal error")
	}
	return c.render(app.FWSMConfig)
}


