package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-admin/app/admin/service/internal/data"

	userV1 "kratos-admin/api/gen/go/user/service/v1"

	"kratos-admin/pkg/utils/name_set"
)

// InitOrganizationManagerId initializes the organization manager IDs into the userSet map.
func InitOrganizationManagerId(orgs []*userV1.Organization, userSet *name_set.UserNameSetMap) {
	for _, v := range orgs {
		if v.ManagerId != nil {
			(*userSet)[v.GetManagerId()] = nil
		}
		for _, c := range v.Children {
			InitOrganizationManagerId(c.Children, userSet)
		}
	}
}

// QueryUserInfoFromRepo queries user information from user repository and fills the nameSetMap.
func QueryUserInfoFromRepo(ctx context.Context, userRepo *data.UserRepo, nameSetMap *name_set.UserNameSetMap) {
	var err error
	var user *userV1.User
	for userId := range *nameSetMap {
		user, err = userRepo.Get(ctx, userId)
		if err != nil {
			log.Errorf("query user err: %v", err)
			continue
		}

		(*nameSetMap)[userId] = &name_set.UserNameSet{
			UserName: user.GetUsername(),
			RealName: user.GetRealname(),
			NickName: user.GetNickname(),
		}
	}
}

func FileOrganizationInfo(orgs []*userV1.Organization, userSet *name_set.UserNameSetMap) {
	for k, v := range *userSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(orgs); i++ {
			if orgs[i].ManagerId != nil && orgs[i].GetManagerId() == k {
				orgs[i].ManagerName = &v.NickName
			}

			FileOrganizationInfo(orgs[i].Children, userSet)
		}
	}
}

// InitDepartmentManagerId initializes the department manager IDs into the userSet map.
func InitDepartmentManagerId(depts []*userV1.Department, userSet *name_set.UserNameSetMap, orgSet *name_set.UserNameSetMap) {
	for _, v := range depts {
		if v.ManagerId != nil {
			(*userSet)[v.GetManagerId()] = nil
		}
		if v.OrganizationId != nil {
			(*orgSet)[v.GetOrganizationId()] = nil
		}

		for _, c := range v.Children {
			InitDepartmentManagerId(c.Children, userSet, orgSet)
		}
	}
}

func QueryOrganizationInfoFromRepo(ctx context.Context, organizationRepo *data.OrganizationRepo, nameSetMap *name_set.UserNameSetMap) {
	var err error
	var org *userV1.Organization
	for orgId := range *nameSetMap {
		org, err = organizationRepo.Get(ctx, &userV1.GetOrganizationRequest{Id: orgId})
		if err != nil {
			log.Errorf("query organization err: %v", err)
			continue
		}

		(*nameSetMap)[orgId] = &name_set.UserNameSet{
			UserName: org.GetName(),
		}
	}
}

func FileDepartmentUserInfo(dpts []*userV1.Department, userSet *name_set.UserNameSetMap) {
	for k, v := range *userSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(dpts); i++ {
			if dpts[i].ManagerId != nil && dpts[i].GetManagerId() == k {
				dpts[i].ManagerName = &v.NickName
			}

			FileDepartmentUserInfo(dpts[i].Children, userSet)
		}
	}
}

func FileDepartmentOrganizationInfo(dpts []*userV1.Department, orgSet *name_set.UserNameSetMap) {
	for k, v := range *orgSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(dpts); i++ {
			if dpts[i].OrganizationId != nil && dpts[i].GetOrganizationId() == k {
				dpts[i].OrganizationName = &v.UserName
			}

			FileDepartmentOrganizationInfo(dpts[i].Children, orgSet)
		}
	}
}
