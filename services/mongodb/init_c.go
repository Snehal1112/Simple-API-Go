package mongodb



import (
	"gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"fmt"
)


var (

	Session *mgo.Session
)


func init() {
	CreateSession()
	InitTaskHandlerCollections()
	// InitIndexTask()
}


func CreateSession() {
	var err error
	Session,err =mgo.Dial("127.0.0.1")
	if err != nil {
		panic (err)
	}
	fmt.Println("SESSION ACTIVE")
}

func CheckMongoDB()  *mgo.Session{
	fmt.Println("ACTIVE MAMMEN")

	return Session

	
	
}

