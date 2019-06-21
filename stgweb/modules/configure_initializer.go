package modules

import (
	"github.com/ttstringiot/golangiot/stgcommon"
	"os"
)

type ConfigureInitializer struct {
	namesrvAddr string `toml:"namesrvAddr" json:"namesrvAddr"`
}

func newConfigureInitializer() *ConfigureInitializer {
	configureInitializer := new(ConfigureInitializer)
	configureInitializer.namesrvAddr = os.Getenv(stgcommon.NAMESRV_ADDR_ENV)
	return configureInitializer
}

func (configureInitializer *ConfigureInitializer) GetNamesrvAddr() string {
	return configureInitializer.namesrvAddr
}
