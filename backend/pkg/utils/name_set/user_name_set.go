package name_set

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	userV1 "go-wind-admin/api/gen/go/user/service/v1"
)

type UserNameSet struct {
	UserName string
	RealName string
	NickName string
	Code     string
}

type UserNameSetMap map[uint32]*UserNameSet

// FillUserInfoFromUserServiceClient fills user information into the UserNameSetMap
func FillUserInfoFromUserServiceClient(ctx context.Context, userClient userV1.UserServiceClient, nameSetMap *UserNameSetMap) {
	var err error
	var user *userV1.User
	for userId := range *nameSetMap {
		user, err = userClient.Get(ctx, &userV1.GetUserRequest{
			QueryBy: &userV1.GetUserRequest_Id{
				Id: userId,
			},
		})
		if err != nil {
			log.Errorf("query user err: %v", err)
			continue
		}

		(*nameSetMap)[userId] = &UserNameSet{
			UserName: user.GetUsername(),
			RealName: user.GetRealname(),
			NickName: user.GetNickname(),
		}
	}
}
