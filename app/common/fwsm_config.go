package common

import (
	"bufio"
	"github.com/xaionaro-go/fwsmConfig"
	"os"
)

const (
	FWSM_CONFIG_PATH = "/root/fwsm-config/dynamic"
)

var (
	FWSMConfig fwsmConfig.FwsmConfig
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ReadConfig() error {
	file, err := os.Open(FWSM_CONFIG_PATH)
	checkErr(err)
	defer file.Close()
	cfgReader := bufio.NewReader(file)
	FWSMConfig, err = fwsmConfig.Parse(cfgReader)
	return err
}

