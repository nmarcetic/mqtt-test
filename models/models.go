package models

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config represents yaml config
type Config struct {
	BrokerURL         string `yaml:"brokerURL"`
	MfxAccessToken    string `yaml:"mfxAccessToken"`
	MfxAccessUsername string `yaml:"mfxAccessUsername"`
	ClientsCount      int    `yaml:"clientsCount"`
	MsgPerClientCount int    `yaml:"msgPerClientCount"`
	QosLevel          byte   `yaml:"qosLevel"`
	ChannelID         int    `yaml:"channelID"`
	SenMLPayload      string `yaml:"senMLPayload"`
	TLSCertPath       string `yaml:"tlsCertPath"`
}

// LoadConfig loads config from yaml and bind it to structure
func (c *Config) LoadConfig() (s *Config, err error) {
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		return nil, err
	}

	return c, nil
}
