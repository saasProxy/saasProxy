package saasProxy

import (
	"fmt"
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
)

func LoadConfiguration(config Configuration) Configuration {
	// Provide the path to your TOML file
	// TODO: load from ~/.saasProxy/saasProxy.toml
	filePath := "./internal/pkg/saasProxy/config.toml"
	_, err := toml.DecodeFile(filePath, &config)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Warn(fmt.Sprintf("saasProxy error loading %s", filePath))
		config = Configuration{}
	}
	log.WithFields(log.Fields{
		"config": config,
	}).Info("saasProxy configuration loaded!")
	return config
}
