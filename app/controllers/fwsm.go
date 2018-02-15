package controllers

import (
	"bytes"
	"encoding/csv"
	"github.com/revel/revel"
	"github.com/xaionaro-go/fwsmAPI/app"
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

func (c FWSM) GetVLANs() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}

	return c.render(app.FWSMConfig.VLANs)
}

func (c FWSM) GetDHCPs() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}

	return c.render(app.FWSMConfig.DHCPs)
}

func (c FWSM) GetSNATs() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}

	return c.render(app.FWSMConfig.SNATs)
}

func (c FWSM) GetDNATs() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}

	return c.render(app.FWSMConfig.DNATs)
}

func (c FWSM) GetRoutes() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}

	return c.render(app.FWSMConfig.Routes)
}

func (c FWSM) GetACLs() revel.Result {
	if !c.IsCanRead() {
		return c.noPerm()
	}

	return c.render(app.FWSMConfig.ACLs)
}

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
