package elastic

import (
 	elastic "gopkg.in/olivere/elastic.v5"
 	// "gopkg.in/inconshreveable/log15.v2"
 	"fmt"
 	"log"
 	"os"
 	"context"
)

var (Client  *elastic.Client)

func init() {
	/**
	*
	* Un comment for using ElasticSearch
	* Client = InitElastic() 
	**/				
}


func InitElastic() (*elastic.Client) {

 var errorlog = log.New(os.Stdout, "APP ", log.LstdFlags)

 	// elastic.SetTraceLog(errorlog)
	client,err := elastic.NewClient(elastic.SetErrorLog(errorlog)) // use for error only
	// client,err := elastic.NewClient(elastic.SetTraceLog(errorlog)) // used for tracing debugging

	if err != nil {
		panic(err)
	}


	return client
}


func CheckElastic() {

	info, code, err := Client.Ping("http://127.0.0.1:9200").Do(context.Background())
	if err != nil {
    	// Handle error
    	panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

}