package mongoIndex

import (
	"fmt"

	"github.com/jethro91/fiber-mongo-redis-boilerplate/src/database/mongoDB"
	"go.mongodb.org/mongo-driver/bson"
)

const roleDatabaseId = ""
const roleCollectionId = "role"

func createRoleIndex() <-chan error {
	chanError := make(chan error)
	go func() {
		defer close(chanError)
		indexes, err := mongoDB.GetIndexes(
			roleDatabaseId,
			roleCollectionId,
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
			err = roleTextSearch()
			chanError <- err
			return
		}
	}()
	return chanError

}

func roleTextSearch() error {
	fmt.Printf("Creating %s Indexes \n", roleCollectionId)
	err := mongoDB.CreateIndex(
		roleDatabaseId,
		roleCollectionId,
		"text_search",
		bson.M{
			"name": "text",
		},
		bson.M{
			"name": 10,
		},
	)
	if err != nil {
		fmt.Printf("Failed Create %s Indexes \n", roleCollectionId)
		return err
	}
	fmt.Printf("Success Create %s Indexes \n", roleCollectionId)

	return nil
}
