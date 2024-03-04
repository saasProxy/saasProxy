package saasProxy

import (
	"fmt"
	"github.com/BurntSushi/toml"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

func LoadConfiguration(filename string, config Configuration) (Configuration, error) {
	// If no filename is provided, use the default path
	if filename == "" {
		filename = "./internal/pkg/saasProxy/config.toml"
	}

	// Get the absolute path for better reliability
	absPath, err := filepath.Abs(filename)
	if err != nil {
		return Configuration{}, fmt.Errorf("error getting absolute path: %w", err)
	}

	// Check if the file exists
	_, err = os.Stat(absPath)
	if os.IsNotExist(err) {
		return Configuration{}, fmt.Errorf("configuration file not found: %s", absPath)
	}

	// Load configuration from the specified file
	_, err = toml.DecodeFile(absPath, &config)
	if err != nil {
		log.WithFields(log.Fields{
			"err": err,
		}).Warn(fmt.Sprintf("saasProxy error loading %s", absPath))
		config = Configuration{}
	}

	log.Info("saasProxy configuration loaded!")
	return config, nil
}

//import (
//	"fmt"
//	"github.com/BurntSushi/toml"
//	log "github.com/sirupsen/logrus"
//)
//
//func LoadConfiguration(config Configuration) Configuration {
//	// Provide the path to your TOML file
//	// TODO: load from ~/.saasProxy/saasProxy.toml
//	filePath := "./internal/pkg/saasProxy/config.toml"
//	_, err := toml.DecodeFile(filePath, &config)
//	if err != nil {
//		log.WithFields(log.Fields{
//			"err": err,
//		}).Warn(fmt.Sprintf("saasProxy error loading %s", filePath))
//		config = Configuration{}
//	}
//	log.Info("saasProxy configuration loaded!")
//	return config
//}
