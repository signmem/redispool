package db

import (
	"context"
	"errors"
	"github.com/go-redis/redis"
	"github.com/signmem/redispool/g"
	"log"
	"strconv"
	"time"
)

var (
	WRITE = 0
	READ = 0
)

type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

func NewDatabase() (*Database, error) {

	address := g.Config().Redis.Server
	client := redis.NewClient(&redis.Options{
		Addr: address,
		Password: "",
		DB: 0,
	})

	if err := client.Ping(Ctx).Err(); err != nil {
		return nil, err
	}

	return &Database{ Client: client, }, nil
}

func (db *Database) PutHost(hostname  string) ( status bool, err error ) {
	timeNow := time.Now().Unix()

	ctx := context.TODO()
	err = db.Client.Set(ctx, hostname, timeNow, 0).Err()
	if err != nil {
		return false, err
	}
	WRITE += 1
	return true, nil
}

func (db *Database) GetAllHosts() (hosts []string, err error) {
	var cursor uint64
	counter :=  int64( g.Config().TestLine )
	ctx := context.TODO()
	keys, cursor, err := db.Client.Scan(ctx, cursor, "*", counter).Result()

	if err != nil {
		log.Printf("[ERROR] GetMetrics error:%s", err)
		return hosts, err
	}

	for _, host := range keys {
		hosts = append(hosts, host)
	}


	if g.Config().Debug {
		log.Printf("[INFO] GetAllHosts() keys length is %d, hosts length is %d\n", len(keys), len(hosts))
	}
	return hosts, nil
}

func (db *Database) CheckAllValues(hosts []string) {
	timeNow := time.Now().Unix()
	timeNowStr := strconv.FormatInt(timeNow, 10)
	ctx := context.TODO()
	for _, host := range hosts {
		value, err := db.Client.Get(ctx, host).Result()
		if err != nil {
			log.Printf("[ERROR] get host %s, value err:%s \n", host, err )
			continue
		}
		valueInt64, err := strconv.ParseInt(value, 10, 64)
		if  timeNow > valueInt64 && ( timeNow - valueInt64 ) > 300  {
			// do someing for alarm
			// testing for delete key
			err := db.DeleteHostKey(host)
			if err != nil {
				log.Printf("[WARNING] now:%s, host %s, value %s\n", timeNowStr, host, value)
				log.Printf("[ERROR] delete key %s from redis err:%s", host, err)
			}

			if g.Config().Debug {
				log.Printf("[INFO] delete key %s from redis ", host)
			}
		}
	}

	if FLUSHKEY == 1 {
		ServerRedisKeyReflush()
	}
}

func (db *Database) DeleteHostKey(hostname string) (err error) {
	ctx := context.TODO()
	err = db.Client.Del(ctx, hostname).Err()
	if err != nil {
		return err
	}
	FLUSHKEY = 1
	return nil
}



