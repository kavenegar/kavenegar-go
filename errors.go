package kavenegar

import (
	"fmt"
)

type APIError struct {
	Status  int
	Message string
	Err     error
}
type HTTPError struct {
	Status  int
	Message string
	Err     error
}

func (e *APIError) Error() string {
	return fmt.Sprintf("APIError[%d] : %s", e.Status, e.Message)
}

func (e *HTTPError) Error() string {
	return fmt.Sprintf("HTTPError[%d] : %s", e.Status, e.Message)
}
