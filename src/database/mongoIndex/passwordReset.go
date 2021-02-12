package mongoIndex

import (
	"fmt"

	"github.com/jethro91/fiber-mongo-redis/src/database/mongoDB"
	"go.mongodb.org/mongo-driver/bson"
)

const passwordResetDatabaseId = ""
const passwordResetCollectionId = "passwordReset"

func createPasswordResetIndex() <-chan error {
	chanError := make(chan error)
	go func() {
		defer close(chanError)
		indexes, err := mongoDB.GetIndexes(
			passwordResetDatabaseId,
			passwordResetCollectionId,
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
			err = passwordResetTextSearch()
			chanError <- err
			return
		}

	}()
	return chanError

}

func passwordResetTextSearch() error {
	fmt.Printf("Creating %s Indexes \n", passwordResetCollectionId)
	err := mongoDB.CreateIndex(
		passwordResetDatabaseId,
		passwordResetCollectionId,
		"text_search",
		bson.M{
			"userId": "text",
			"token":  "text",
		},
		bson.M{
			"userId": 10,
			"token":  5,
		},
	)
	if err != nil {
		fmt.Printf("Failed Create %s Indexes \n", passwordResetCollectionId)
		return err
	}
	fmt.Printf("Success Create %s Indexes \n", passwordResetCollectionId)

	return nil
}
