package mongodb





import (
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


type TaskMongoModel struct{
	IDTask     int         `boil:"id_task" json:"id_task" toml:"id_task" yaml:"id_task"`
	TitleTask  string      `boil:"title_task" json:"title_task" toml:"title_task" yaml:"title_task"`
	Urgency    string      `boil:"urgency" json:"urgency" toml:"urgency" yaml:"urgency"`
	Difficulty string      `boil:"difficulty" json:"difficulty" toml:"difficulty" yaml:"difficulty"`
	DescTask   null.String `boil:"desc_task" json:"desc_task,omitempty" toml:"desc_task" yaml:"desc_task,omitempty"`
}




var (

	Collection *mgo.Collection
)

func InitTaskHandlerCollections() {
	Collection =  Session.DB("Taskref").C("tasks")

}




func DeleteTaskRefCollections() {
	// Collection= 
	Collection.DropCollection()
}

func InsertDataTask(Task models.Oktask) error{

	 var err error
	 // var tobeInserted TaskMongoModel
	 
	 

	 
	 // err = Collection.Insert(tobeInserted)

	 // info, err := Collection.Upsert(bson.M{"_id": Task.IDRef}, tobeInserted)

	 fmt.Println("Okeee")
	 return err

}



