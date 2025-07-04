package login

import (
	"fmt"

	ssov1 "github.com/krawwwwy/protos/gen/go/sso"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const emptyValue = 0

func ValidationError(req *ssov1.LoginRequest) error {
	email := req.GetEmail()
	password := req.GetPassword()
	app_id := req.GetAppid()

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
