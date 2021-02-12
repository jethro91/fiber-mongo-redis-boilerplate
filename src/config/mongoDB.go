//nolint
package config

import (
	"fmt"
	"os"
)

var MONGO_USERNAME = os.Getenv("MONGO_USERNAME")
var MONGO_PASSWORD = os.Getenv("MONGO_PASSWORD")

var MONGO_HOST = os.Getenv("MONGO_HOST")
var MONGO_PORT = os.Getenv("MONGO_PORT")
var MONGO_DATABASE = os.Getenv("MONGO_DATABASE")

var MONGO_URI = fmt.Sprintf("mongodb://%s:%s@%s:%s/%s", MONGO_USERNAME, MONGO_PASSWORD, MONGO_HOST, MONGO_PORT, MONGO_DATABASE)
