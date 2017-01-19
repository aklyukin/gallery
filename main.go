package main

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Struct for postgresql config
type instanceConfig struct {
	Pgsql struct {
		Host     string
		Port     string
		Username string
		Password string
		Database string
	}
}

func main() {
	log.Println("Hello world!!!")
	configFile, err := ioutil.ReadFile("./cfg/config.yml")
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	var config instanceConfig

	error := yaml.Unmarshal(configFile, &config)
	if error != nil {
		log.Fatal(err)
	}

	// initDB(config)
	reforminitDB(config)

	testUser := &User{
		Email:    RandStringBytes(10),
		Password: RandStringBytes(10),
	}

	SaveUser(testUser)
}
