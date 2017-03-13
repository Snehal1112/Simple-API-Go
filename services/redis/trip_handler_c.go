package redis


import (
	// "gopkg.in/redis.v5"
	// "fmt"
	// "time"
	// "bytes"
	// "encoding/json"
	// "simple-api/core/models"
	// "strconv"
)








func LastInsertedIndex( index int ) error {
	
	err:= Client.Set("create_task#last_inserted#id" , index,0).Err()

	return err;
}

func GetLastInsertedIndex() (string,error) {
	
	
	val, err := Client.Get("create_task#last_inserted#id").Result();
	
	return val, err;
}