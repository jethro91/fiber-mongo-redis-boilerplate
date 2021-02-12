//nolint
package config

import (
	"fmt"
	"os"
)

var APP_SECRET = os.Getenv("APP_SECRET")

var APP_HOST = os.Getenv("APP_HOST")
var APP_PORT = os.Getenv("APP_PORT")
var APP_LISTEN = fmt.Sprintf("%s:%s", APP_HOST, APP_PORT)

var APP_CLIENT_HOST = os.Getenv("APP_CLIENT_HOST")
var APP_CLIENT_PORT = os.Getenv("APP_CLIENT_PORT")
var APP_CLIENT_URL = fmt.Sprintf("http://%s:%s", APP_CLIENT_HOST, APP_CLIENT_PORT)
