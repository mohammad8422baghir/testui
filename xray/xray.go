
package xrayproxy

import (
	"github.com/xtls/xray-core/core"
	_ "github.com/xtls/xray-core/main/distro/all"
)

func StartXray(configJSON string) error {
	config, err := core.LoadConfig("json", configJSON)
	if err != nil {
		return err
	}
	server, err := core.New(config)
	if err != nil {
		return err
	}
	return server.Start()
}
