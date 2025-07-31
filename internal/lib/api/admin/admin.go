package admin

import (
	ssov1 "github.com/krawwwwy/rosatomprotos"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const emptyValue = 0

func ValidationError(req *ssov1.IsAdminRequest) error {
	userID := req.GetUserId()

	if userID == emptyValue {
		return status.Error(codes.InvalidArgument, "user_id is required")
	}
	return nil
}
