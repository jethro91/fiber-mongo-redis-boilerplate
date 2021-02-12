//nolint
package config

import "os"

var SMTP_HOST = os.Getenv("SMTP_HOST")
var SMTP_PORT = os.Getenv("SMTP_PORT")
var SMTP_USERNAME = os.Getenv("SMTP_USERNAME")
var SMTP_PASSWORD = os.Getenv("SMTP_PASSWORD")

var SMTP_FROM = os.Getenv("SMTP_FROM")
var SMTP_BCC = os.Getenv("SMTP_BCC")
var SMTP_REPLY_TO = os.Getenv("SMTP_REPLY_TO")
