package login

import (
	"fmt"

	ssov1 "github.com/krawwwwy/rosatomprotos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const emptyValue = 0

func ValidationError(req *ssov1.LoginRequest) error {
	email := req.GetEmail()
	password := req.GetPassword()
	app_id := req.GetAppid()

	list := []string{email, password}
	dict := []string{"email", "password"}

	for key, field := range list {
		if field == "" {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("%s is required", dict[key]))
		}
	}

	if app_id == emptyValue {
		return status.Error(codes.InvalidArgument, "app_id is required")
	}

	return nil
}
