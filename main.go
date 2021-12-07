package main

import (
	"flag"
	"fmt"
	"github.com/signmem/redispool/db"
	"github.com/signmem/redispool/g"
	"github.com/signmem/redispool/http"
	"log"
	"os"
	"time"
)

func main() {
	cfg := flag.String("c", "cfg.json", "configuration file")
	version := flag.Bool("v", false, "show version")

	flag.Parse()

	if *version {
		version := g.Version
		fmt.Printf("%s", version)
		os.Exit(0)
	}

	g.ParseConfig(*cfg)

	g.InitLog()
	log.Println("[INFO] log init success.")


	go db.ResetMetric()
	go http.Start()

	if g.Config().ROLE == "client" {
		go db.WriteTest()
	}

	if g.Config().ROLE == "server" {
		go db.ServerRedisKeyMaintain()
		time.Sleep(3 * time.Second)
		go db.ServerRedisValueMaintain()
	}
	select {}

}