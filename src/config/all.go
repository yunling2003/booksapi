package config


import (	
	"os"
	"fmt"
	"io/ioutil"	
	"gopkg.in/yaml.v2"
)

var (
	All map[string]string
)

func init() {
	configFile := getConfigFile()
	bs, err := ioutil.ReadFile(configFile)

	if err != nil {
		panic(fmt.Sprintf("could not find config file"))
	}

	if err := yaml.Unmarshal(bs, &All); err != nil {
		panic(err.Error())
	}
}

func getConfigFile() string {
	configFile := "./config/setting.yaml"	
	return configFile
}

func fileExists(fileName string) bool {
	if _, err := os.Stat(fileName); err == nil {
		return true
	}
	return false
}