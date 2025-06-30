package login

import (
	"fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const emptyValue = 0

func ValidationError(email string, password string, app_id int32) error {

	for _, field := range []string{email, password} {
		if field == "" {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("%s is required", field))
		}
	}

	if app_id == emptyValue {
		return status.Error(codes.InvalidArgument, "app_id is required")
	}

	return nil
}
