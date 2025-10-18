package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-admin/app/admin/service/internal/data"

	userV1 "kratos-admin/api/gen/go/user/service/v1"

	"kratos-admin/pkg/utils/name_set"
)

// InitOrganizationManagerId initializes the organization manager IDs into the userSet map.
func InitOrganizationManagerId(organizations []*userV1.Organization, userSet *name_set.UserNameSetMap) {
	for _, v := range organizations {
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

func FileOrganizationInfo(organizations []*userV1.Organization, userSet *name_set.UserNameSetMap) {
	for k, v := range *userSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(organizations); i++ {
			if organizations[i].ManagerId != nil && organizations[i].GetManagerId() == k {
				organizations[i].ManagerName = &v.NickName
			}

			FileOrganizationInfo(organizations[i].Children, userSet)
		}
	}
}

// InitDepartmentManagerId initializes the department manager IDs into the userSet map.
func InitDepartmentManagerId(departments []*userV1.Department, userSet *name_set.UserNameSetMap, orgSet *name_set.UserNameSetMap) {
	for _, v := range departments {
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

func QueryDepartmentInfoFromRepo(ctx context.Context, departmentRepo *data.DepartmentRepo, nameSetMap *name_set.UserNameSetMap) {
	var err error
	var dept *userV1.Department
	for deptId := range *nameSetMap {
		dept, err = departmentRepo.Get(ctx, &userV1.GetDepartmentRequest{Id: deptId})
		if err != nil {
			log.Errorf("query department err: %v", err)
			continue
		}

		(*nameSetMap)[deptId] = &name_set.UserNameSet{
			UserName: dept.GetName(),
		}
	}
}

func FileDepartmentUserInfo(departments []*userV1.Department, userSet *name_set.UserNameSetMap) {
	for k, v := range *userSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(departments); i++ {
			if departments[i].ManagerId != nil && departments[i].GetManagerId() == k {
				departments[i].ManagerName = &v.NickName
			}

			FileDepartmentUserInfo(departments[i].Children, userSet)
		}
	}
}

func FileDepartmentOrganizationInfo(departments []*userV1.Department, orgSet *name_set.UserNameSetMap) {
	for k, v := range *orgSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(departments); i++ {
			if departments[i].OrganizationId != nil && departments[i].GetOrganizationId() == k {
				departments[i].OrganizationName = &v.UserName
			}

			FileDepartmentOrganizationInfo(departments[i].Children, orgSet)
		}
	}
}

func InitPositionOrgId(positions []*userV1.Position, orgSet *name_set.UserNameSetMap, deptSet *name_set.UserNameSetMap) {
	for _, v := range positions {
		if v.OrganizationId != nil {
			(*orgSet)[v.GetOrganizationId()] = nil
		}
		if v.DepartmentId != nil {
			(*deptSet)[v.GetDepartmentId()] = nil
		}

		for _, c := range v.Children {
			InitPositionOrgId(c.Children, orgSet, deptSet)
		}
	}
}

func FilePositionOrganizationInfo(positions []*userV1.Position, orgSet *name_set.UserNameSetMap) {
	for k, v := range *orgSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(positions); i++ {
			if positions[i].OrganizationId != nil && positions[i].GetOrganizationId() == k {
				positions[i].OrganizationName = &v.UserName
			}

			FilePositionOrganizationInfo(positions[i].Children, orgSet)
		}
	}
}

func FilePositionDepartmentInfo(positions []*userV1.Position, deptSet *name_set.UserNameSetMap) {
	for k, v := range *deptSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(positions); i++ {
			if positions[i].DepartmentId != nil && positions[i].GetDepartmentId() == k {
				positions[i].DepartmentName = &v.UserName
			}

			FilePositionDepartmentInfo(positions[i].Children, deptSet)
		}
	}
}
