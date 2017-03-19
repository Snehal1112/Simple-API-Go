package models



import (
	"time"

	// elastic "gopkg.in/olivere/elastic.v5"
	"gopkg.in/nullbio/null.v6"
	// "gopkg.in/gin-gonic/gin.v1"
	"simple-api/core/models"
	"fmt"
	// "strconv"
	// "context"
	// "encoding/json"
	// "bytes"
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
)



type User struct{
	Id      	bson.ObjectId 	`json:"id" bson:"_id,omitempty"`
 	Username	string		  	`json:"username"`	
 	Password	string			`json:"password"`
 	Email		string			`json:"email"`
 	CreatedAt	time.Time		`json:"created_at"`
 	UpdatedAt	time.Time		`json:"updated_at"`
}


var (

	Collection *mgo.Collection
)

func InitUserCollection() {
	Collection =  Session.DB("Taskref").C("user")	
}

func DropUserCollections() {
	Collection.DropCollection()
}

func (U User)InsertOne() error{

	_, err := Collection.Insert(tobeInserted)

	return err;
}

func (U User)Upsert() error{

	_, err := Collection.Upsert(tobeInserted)

	return err;
}
func InsertMany(arr []User) error{
	
}