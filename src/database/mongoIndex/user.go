package mongoIndex

import (
	"fmt"

	"github.com/jethro91/fiber-mongo-redis/src/database/mongoDB"
	"go.mongodb.org/mongo-driver/bson"
)

const userDatabaseId = ""
const userCollectionId = "user"

func createUserIndex() <-chan error {
	chanError := make(chan error)
	go func() {
		defer close(chanError)
		indexes, err := mongoDB.GetIndexes(
			userDatabaseId,
			userCollectionId,
		)
		if err != nil {
			chanError <- err
			return
		}

		isTextSearchExist := false
		for _, data := range indexes {
			if data["name"] == "text_search" {
				isTextSearchExist = true
			}
		}

		if !isTextSearchExist {
			err = userTextSearch()
			chanError <- err
			return
		}
	}()
	return chanError

}

func userTextSearch() error {
	fmt.Printf("Creating %s Indexes \n", userCollectionId)
	err := mongoDB.CreateIndex(
		userDatabaseId,
		userCollectionId,
		"text_search",
		bson.M{
			"name":  "text",
			"email": "text",
		},
		bson.M{
			"name":  10,
			"email": 1,
		},
	)
	if err != nil {
		fmt.Printf("Failed Create %s Indexes \n", userCollectionId)
		return err
	}
	fmt.Printf("Success Create %s Indexes \n", userCollectionId)

	return nil
}
