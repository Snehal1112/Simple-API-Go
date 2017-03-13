package redis

import (
	"gopkg.in/redis.v5"
	"fmt"
)


var (

    Client  *redis.Client
)

func init() {
    
    var err error
    Client , err = InitRedisClient();
    if err != nil {
        panic(err)
    }

}

func InitRedisClient() (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    pong, err := client.Ping().Result()
    
    fmt.Println(pong, err)

    if err != nil {
    	return nil , err
    }

    return client,nil;
}