package db

import (
	"fmt"
	"github.com/signmem/redispool/g"
	"log"
	"strconv"
	"time"
)

var (
	REDISHOSTKEY []string
	FLUSHKEY int
)

func WriteTest() {
	client, err := NewDatabase()
	if err != nil {
		fmt.Printf("[ERROR] connection to Redis Server err:%s", err)
		return
	}
	num := 1
	endLine := g.Config().TestLine
	for {
		if num >= endLine {
			num = 1
		}
		numStr := strconv.Itoa(num)
		hostname := "falcon-agent-test-" + numStr
		_, err = client.PutHost(hostname)
		if err != nil {
			log.Printf("[ERROR] put host %s err:%s", hostname, err)
		}
		num += 1
	}
}

func ServerRedisKeyMaintain() {
	client, err := NewDatabase()

	if err != nil {
		fmt.Printf("[ERROR] connection to Redis Server err:%s", err)
		return
	}

	for {
		hosts , err := client.GetAllHosts()
		REDISHOSTKEY = hosts
		if err != nil {
			return
		}

		FLUSHKEY = 0

		if g.Config().Debug {
			log.Printf("[INFO] key len is: %d\n", len(REDISHOSTKEY))
		}

		time.Sleep(time.Second * 300)
	}

}

func ServerRedisKeyReflush() {
	client, err := NewDatabase()

	if err != nil {
		fmt.Printf("[ERROR] connection to Redis Server err:%s", err)
		return
	}

	hosts , err := client.GetAllHosts()
	REDISHOSTKEY = hosts
	if err != nil {
		return
	}

	FLUSHKEY = 0

	if g.Config().Debug {
		log.Printf("[INFO] ServerRedisKeyReflush() key len is: %d\n", len(REDISHOSTKEY))
	}
}

func ServerRedisValueMaintain() {
	client, err := NewDatabase()
	if err != nil {
		fmt.Printf("[ERROR] connection to Redis Server err:%s", err)
		return
	}

	if g.Config().Debug {
		log.Printf("[INFO] Going to get keys. length is %d\n", len(REDISHOSTKEY))
	}

	for {
		if len(REDISHOSTKEY) == 0 {
			time.Sleep(time.Second * 60)
		} else {
			client.CheckAllValues(REDISHOSTKEY)
			time.Sleep(time.Second * 60)
		}
	}
}