package rules

import "errors"

var INSUFFICIENTPERMISSIONS = errors.New("insufficient permissions")
var OTHERERROR = errors.New("other error")
var NOMATCHINGPOLICY = errors.New("no matching policy")
