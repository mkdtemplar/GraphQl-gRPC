package graph

import (
	"graphqhhowto/gRPC/proto"
	"graphqhhowto/graph/model"
)

func createUserInDb(c proto.UserServiceClient) *model.User {
	return &model.User{}
}
