package common

import (
	"github.com/elastic/go-elasticsearch/v6"
	"log"
	"net"
	"os"
	"strings"
	"time"

	mapset "github.com/deckarep/golang-set"
	"github.com/go-ego/gse"
	"github.com/go-redis/redis"
	"github.com/paulbellamy/ratecounter"
	mgo "gopkg.in/mgo.v2"
)

func InitServer() {

	mongoAddr = os.Getenv("mongoAddr")
	dataBase = os.Getenv("mongoDatabase")
	collection = os.Getenv("mongoCollection")
	mongoUsername = os.Getenv("mongoUsername")
	mongoPassWord = os.Getenv("mongoPassWord")
	verifyPassord = os.Getenv("verifyPassord")
	tmp := os.Getenv("wkNodes")
	wkNodes = strings.Split(tmp, ",")
	for _, j := range wkNodes {
		if _, _, err := net.SplitHostPort(j); err != nil {
			panic("wkNodes set error")
		}
	}
	banList = os.Getenv("banList")
	esURL = os.Getenv("esURL")
	esUrlBase = os.Getenv("esUrlBase")
	esUsername = os.Getenv("esUsername")
	esPassWord = os.Getenv("esPassWord")
}

//NewSniffer :NewSniffer
func NewSniffer() *Server {
	for {
		time.Sleep(time.Second)
		log.Println("trying to connect es")
		var err error
		ES, err = elasticsearch.NewClient(elasticsearch.Config{
			Addresses: []string{esUrlBase},
			Username:  esUsername,
			Password:  esPassWord,
		})
		if err != nil {
			log.Printf("Error creating the client: %s", err)
			continue
		}
		res, err := ES.Info()
		if err != nil {
			log.Printf("Error getting response: %s", err)
			continue
		}
		res.Body.Close()
		log.Println("connect es suss")
		log.Println(res)
		break
	}
	dialInfo := &mgo.DialInfo{
		Addrs:  []string{mongoAddr},
		Direct: false,
		//Timeout: time.Second * 1,
		Database:  dataBase,
		Source:    collection,
		Username:  mongoUsername,
		Password:  mongoPassWord,
		PoolLimit: 4096, // Session.SetPoolLimit
	}

	session, err := mgo.DialWithInfo(dialInfo)
	session.SetPoolLimit(mongoConnectLimitNum)
	session.SetMode(mgo.Monotonic, true)
	if err != nil {
		panic(err.Error)
	}
	var redisClient *redis.Client
	if redisEnable {
		redisClient = redis.NewClient(&redis.Options{
			Addr:     redisAddr,
			Password: redisPassword,
			DB:       redisDB,
		})
		_, err := redisClient.Ping().Result()
		if err != nil {
			panic(err.Error())
		}
	}

	var segmenter gse.Segmenter
	segmenter.LoadDict("config/dictionary.txt")

	return &Server{
		segmenter:   segmenter, //分词
		printChan:   make(chan string, 5),
		hashList:    mapset.NewSet(),
		Tool:        *NewTool(),
		Nodes:       wkNodes,
		Mon:         session,
		RedisClient: redisClient,
		mongoLimit:  make(chan bool, mongoConnectLimitNum),
		blackList:   loadBlackList(),
		revNum:      ratecounter.NewRateCounter(1 * time.Second),
		dropSpeed:   ratecounter.NewRateCounter(1 * time.Second),
		sussNum:     ratecounter.NewRateCounter(1 * time.Second),
		notFoundNum: ratecounter.NewRateCounter(1 * time.Second),
	}
}
