package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

type Kafka struct {
	Host     string `yaml:"host"`
	Topic    string `yaml:"topic"`
	Port     string `yaml:"port"`
	MaxBytes int    `yaml:"max_bytes"`
}

type Config struct {
	Kafka *Kafka `yaml:"kafka"`
}

// func NewConfig() *KafkaConfig {
// 	return &KafkaConfig{}
// }

func GetConf() *Config {
	conf := Config{}
	yamlFile, err := ioutil.ReadFile("config.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, &conf)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return &conf
}
