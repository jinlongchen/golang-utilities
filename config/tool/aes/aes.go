package main

import (
	"flag"
	"io/ioutil"
	"github.com/jinlongchen/golang-utilities/log"
	"github.com/jinlongchen/golang-utilities/config"
)

func main() {
	confFile := flag.String("conf", "./conf-file.toml", "")
	dataFile := flag.String("data", "./data.txt", "")

	flag.Parse()

	data, err := ioutil.ReadFile(*dataFile)
	if err != nil {
		log.Fatalf("cannot read data file: %s", *dataFile)
	}

	encrypted := config.NewConfig(*confFile).EncryptString(string(data))

	log.Infof("encrypted: %s", encrypted)
}
