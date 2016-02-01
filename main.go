package main

import (
	"fmt"
	//"github.com/golang/example/stringutil"
	//"gopkg.in/mgo.v2"
	//"gopkg.in/mgo.v2/bson"
	"flag"
	"gopkg.in/gcfg.v1"
	"log"
)

var config struct{
	Database struct {
		Url1 string
	 }
}

var cfg = config

func main() {
	ip := flag.String("config","config.cfg","sets the config file")

	flag.Parse()
	fmt.Printf("Using config file '%v' \n", *ip)

	err := gcfg.ReadFileInto(&cfg, *ip)
	if err != nil {
		log.Fatal("Failed to parse config data from file %v with err: %s", *ip, err)
	}
	fmt.Println(cfg.Database.Url1)
}
