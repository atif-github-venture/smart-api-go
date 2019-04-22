package utility

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

type Prop struct {
	MongoDBUrl string `yaml:"mongourl"`
}

func (c *Prop) ReadProperty() *Prop {
	absPath, _ := filepath.Abs("../smartapigo/property.yml")
	yamlFile, err := ioutil.ReadFile(absPath)
	if err != nil {
		print(err)
		absPath, _ = filepath.Abs("../property.yml")
		yamlFile, err = ioutil.ReadFile(absPath)
		if err != nil {
			panic(err)
		}
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		panic(err)
	}
	return c
}
