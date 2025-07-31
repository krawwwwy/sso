package register

import (
	"fmt"

	ssov1 "github.com/krawwwwy/rosatomprotos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func ValidationError(req *ssov1.RegisterRequest) error {
	email := req.GetEmail()
	password := req.GetPassword()

	for _, field := range []string{email, password} {
		if field == "" {
			return status.Error(codes.InvalidArgument, fmt.Sprintf("%s is required", field))
		}
	}
	return nil
}
