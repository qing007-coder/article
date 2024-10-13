package errors

import "errors"

var DBInitFailed = errors.New("db init failed")
var AutoMigrateFailed = errors.New("auto migrate failed")
