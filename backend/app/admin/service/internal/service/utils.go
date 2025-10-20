package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"

	"kratos-admin/app/admin/service/internal/data"

	userV1 "kratos-admin/api/gen/go/user/service/v1"

	"kratos-admin/pkg/utils/name_set"
)

// InitOrganizationNameSetMap initializes the organization manager IDs into the userSet map.
func InitOrganizationNameSetMap(organizations []*userV1.Organization, userSet *name_set.UserNameSetMap) {
	for _, v := range organizations {
		if v.ManagerId != nil {
			(*userSet)[v.GetManagerId()] = nil
		}
		for _, c := range v.Children {
			InitOrganizationNameSetMap(c.Children, userSet)
		}
	}
}

// QueryUserInfoFromRepo queries user information from user repository and fills the nameSetMap.
func QueryUserInfoFromRepo(ctx context.Context, userRepo *data.UserRepo, nameSetMap *name_set.UserNameSetMap) {
	userIds := make([]uint32, 0, len(*nameSetMap))
	for userId := range *nameSetMap {
		userIds = append(userIds, userId)
	}

	users, err := userRepo.GetUsersByIds(ctx, userIds)
	if err != nil {
		log.Errorf("query users err: %v", err)
		return
	}

	for _, user := range users {
		(*nameSetMap)[user.GetId()] = &name_set.UserNameSet{
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

// InitDepartmentNameSetMap initializes the department manager IDs into the userSet map.
func InitDepartmentNameSetMap(departments []*userV1.Department, userSet *name_set.UserNameSetMap, orgSet *name_set.UserNameSetMap) {
	for _, v := range departments {
		if v.ManagerId != nil {
			(*userSet)[v.GetManagerId()] = nil
		}
		if v.OrganizationId != nil {
			(*orgSet)[v.GetOrganizationId()] = nil
		}

		for _, c := range v.Children {
			InitDepartmentNameSetMap(c.Children, userSet, orgSet)
		}
	}
}

func QueryOrganizationInfoFromRepo(ctx context.Context, organizationRepo *data.OrganizationRepo, nameSetMap *name_set.UserNameSetMap) {
	var orgIds []uint32
	for orgId := range *nameSetMap {
		orgIds = append(orgIds, orgId)
	}

	orgs, err := organizationRepo.GetOrganizationsByIds(ctx, orgIds)
	if err != nil {
		log.Errorf("query organizations err: %v", err)
		return
	}

	for _, o := range orgs {
		(*nameSetMap)[o.GetId()] = &name_set.UserNameSet{
			UserName: o.GetName(),
		}
	}
}

func QueryDepartmentInfoFromRepo(ctx context.Context, departmentRepo *data.DepartmentRepo, nameSetMap *name_set.UserNameSetMap) {
	var deptIds []uint32
	for deptId := range *nameSetMap {
		deptIds = append(deptIds, deptId)
	}

	depts, err := departmentRepo.GetDepartmentsByIds(ctx, deptIds)
	if err != nil {
		log.Errorf("query departments err: %v", err)
		return
	}

	for _, dept := range depts {
		(*nameSetMap)[dept.GetId()] = &name_set.UserNameSet{
			UserName: dept.GetName(),
		}
	}
}

func FillDepartmentUserInfo(departments []*userV1.Department, userSet *name_set.UserNameSetMap) {
	for k, v := range *userSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(departments); i++ {
			if departments[i].ManagerId != nil && departments[i].GetManagerId() == k {
				departments[i].ManagerName = &v.NickName
			}

			FillDepartmentUserInfo(departments[i].Children, userSet)
		}
	}
}

func FillDepartmentOrganizationInfo(departments []*userV1.Department, orgSet *name_set.UserNameSetMap) {
	for k, v := range *orgSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(departments); i++ {
			if departments[i].OrganizationId != nil && departments[i].GetOrganizationId() == k {
				departments[i].OrganizationName = &v.UserName
			}

			FillDepartmentOrganizationInfo(departments[i].Children, orgSet)
		}
	}
}

func InitPositionNameSetMap(positions []*userV1.Position, orgSet *name_set.UserNameSetMap, deptSet *name_set.UserNameSetMap) {
	for _, v := range positions {
		if v.OrganizationId != nil {
			(*orgSet)[v.GetOrganizationId()] = nil
		}
		if v.DepartmentId != nil {
			(*deptSet)[v.GetDepartmentId()] = nil
		}

		for _, c := range v.Children {
			InitPositionNameSetMap(c.Children, orgSet, deptSet)
		}
	}
}

func FillPositionOrganizationInfo(positions []*userV1.Position, orgSet *name_set.UserNameSetMap) {
	for k, v := range *orgSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(positions); i++ {
			if positions[i].OrganizationId != nil && positions[i].GetOrganizationId() == k {
				positions[i].OrganizationName = &v.UserName
			}

			FillPositionOrganizationInfo(positions[i].Children, orgSet)
		}
	}
}

func FillPositionDepartmentInfo(positions []*userV1.Position, deptSet *name_set.UserNameSetMap) {
	for k, v := range *deptSet {
		if v == nil {
			continue
		}

		for i := 0; i < len(positions); i++ {
			if positions[i].DepartmentId != nil && positions[i].GetDepartmentId() == k {
				positions[i].DepartmentName = &v.UserName
			}

			FillPositionDepartmentInfo(positions[i].Children, deptSet)
		}
	}
}

func InitUserNameSetMap(users []*userV1.User, orgSet *name_set.UserNameSetMap, deptSet *name_set.UserNameSetMap, posSet *name_set.UserNameSetMap, roleSet *name_set.UserNameSetMap) {
	for _, v := range users {
		if v.OrgId != nil {
			(*orgSet)[v.GetOrgId()] = nil
		}
		if v.DepartmentId != nil {
			(*deptSet)[v.GetDepartmentId()] = nil
		}
		if v.PositionId != nil {
			(*posSet)[v.GetPositionId()] = nil
		}
		for _, roleId := range v.RoleIds {
			(*roleSet)[roleId] = nil
		}
	}
}

func QueryPositionInfoFromRepo(ctx context.Context, positionRepo *data.PositionRepo, nameSetMap *name_set.UserNameSetMap) {
	var posIds []uint32
	for posId := range *nameSetMap {
		posIds = append(posIds, posId)
	}

	poss, err := positionRepo.GetPositionByIds(ctx, posIds)
	if err != nil {
		log.Errorf("query positions err: %v", err)
		return
	}

	for _, position := range poss {
		(*nameSetMap)[position.GetId()] = &name_set.UserNameSet{
			UserName: position.GetName(),
		}
	}
}

func QueryRoleInfoFromRepo(ctx context.Context, roleRepo *data.RoleRepo, nameSetMap *name_set.UserNameSetMap) {
	var roleIds []uint32
	for roleId := range *nameSetMap {
		roleIds = append(roleIds, roleId)
	}

	roles, err := roleRepo.GetRolesByRoleIds(ctx, roleIds)
	if err != nil {
		log.Errorf("query roles err: %v", err)
		return
	}

	for _, role := range roles {
		(*nameSetMap)[role.GetId()] = &name_set.UserNameSet{
			UserName: role.GetName(),
			Code:     role.GetCode(),
		}
	}
}
