package shared

import "errors"

var ErrResourceNotFound = errors.New("resource not found")
var ErrPermissionDenied = errors.New("permission denied")
var ErrTransientFailure = errors.New("transient failure")
var ErrInvalidRequest = errors.New("invalid request")
var ErrConflict = errors.New("conflict")
