package admin

import (
	"context"

	"github.com/dl-alexandre/gdrive/internal/api"
	"github.com/dl-alexandre/gdrive/internal/types"
	adminapi "google.golang.org/api/admin/directory/v1"
	"google.golang.org/api/googleapi"
)

type Manager struct {
	client  *api.Client
	service *adminapi.Service
}

func NewManager(client *api.Client, service *adminapi.Service) *Manager {
	return &Manager{
		client:  client,
		service: service,
	}
}

type ListUsersOptions struct {
	Customer   string
	Domain     string
	Query      string
	MaxResults int64
	PageToken  string
	OrderBy    string
	Fields     string
	Paginate   bool
}

func (m *Manager) ListUsers(ctx context.Context, reqCtx *types.RequestContext, opts *ListUsersOptions) (*types.UsersListResponse, error) {
	var allUsers []types.User
	pageToken := opts.PageToken
	for {
		call := m.service.Users.List()
		if opts.Domain != "" {
			call = call.Domain(opts.Domain)
		}
		if opts.Customer != "" {
			call = call.Customer(opts.Customer)
		}
		if opts.Query != "" {
			call = call.Query(opts.Query)
		}
		if opts.MaxResults > 0 {
			call = call.MaxResults(opts.MaxResults)
		}
		if pageToken != "" {
			call = call.PageToken(pageToken)
		}
		if opts.OrderBy != "" {
			call = call.OrderBy(opts.OrderBy)
		}
		if opts.Fields != "" {
			call = call.Fields(googleapi.Field(opts.Fields))
		}

		result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*adminapi.Users, error) {
			return call.Do()
		})
		if err != nil {
			return nil, err
		}

		allUsers = append(allUsers, convertUsers(result)...)
		if !opts.Paginate || result.NextPageToken == "" {
			return &types.UsersListResponse{
				Users:         allUsers,
				NextPageToken: result.NextPageToken,
			}, nil
		}
		pageToken = result.NextPageToken
	}
}

func (m *Manager) GetUser(ctx context.Context, reqCtx *types.RequestContext, userKey string, fields string) (*types.User, error) {
	call := m.service.Users.Get(userKey)
	if fields != "" {
		call = call.Fields(googleapi.Field(fields))
	}
	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*adminapi.User, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}
	converted := convertUser(result)
	return &converted, nil
}

func (m *Manager) CreateUser(ctx context.Context, reqCtx *types.RequestContext, req *types.CreateUserRequest) (*types.User, error) {
	user := &adminapi.User{
		PrimaryEmail: req.Email,
		Name: &adminapi.UserName{
			GivenName:  req.GivenName,
			FamilyName: req.FamilyName,
		},
		Password: req.Password,
	}
	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*adminapi.User, error) {
		return m.service.Users.Insert(user).Do()
	})
	if err != nil {
		return nil, err
	}
	converted := convertUser(result)
	return &converted, nil
}

func (m *Manager) DeleteUser(ctx context.Context, reqCtx *types.RequestContext, userKey string) error {
	_, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (struct{}, error) {
		return struct{}{}, m.service.Users.Delete(userKey).Do()
	})
	return err
}

func (m *Manager) UpdateUser(ctx context.Context, reqCtx *types.RequestContext, userKey string, req *types.UpdateUserRequest) (*types.User, error) {
	user := &adminapi.User{}
	fieldMask := []string{}

	if req.GivenName != nil || req.FamilyName != nil {
		user.Name = &adminapi.UserName{}
		if req.GivenName != nil {
			user.Name.GivenName = *req.GivenName
			fieldMask = append(fieldMask, "name.givenName")
		}
		if req.FamilyName != nil {
			user.Name.FamilyName = *req.FamilyName
			fieldMask = append(fieldMask, "name.familyName")
		}
	}
	if req.Suspended != nil {
		user.Suspended = *req.Suspended
		fieldMask = append(fieldMask, "suspended")
	}
	if req.OrgUnitPath != nil {
		user.OrgUnitPath = *req.OrgUnitPath
		fieldMask = append(fieldMask, "orgUnitPath")
	}
	if len(fieldMask) == 0 {
		return m.GetUser(ctx, reqCtx, userKey, "")
	}

	call := m.service.Users.Patch(userKey, user)

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*adminapi.User, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}
	converted := convertUser(result)
	return &converted, nil
}

func (m *Manager) SuspendUser(ctx context.Context, reqCtx *types.RequestContext, userKey string) (*types.User, error) {
	return m.UpdateUser(ctx, reqCtx, userKey, &types.UpdateUserRequest{
		Suspended: boolPtr(true),
	})
}

func (m *Manager) UnsuspendUser(ctx context.Context, reqCtx *types.RequestContext, userKey string) (*types.User, error) {
	return m.UpdateUser(ctx, reqCtx, userKey, &types.UpdateUserRequest{
		Suspended: boolPtr(false),
	})
}

type ListGroupsOptions struct {
	Customer   string
	Domain     string
	Query      string
	MaxResults int64
	PageToken  string
	OrderBy    string
	Fields     string
	Paginate   bool
}

func (m *Manager) ListGroups(ctx context.Context, reqCtx *types.RequestContext, opts *ListGroupsOptions) (*types.GroupsListResponse, error) {
	var allGroups []types.Group
	pageToken := opts.PageToken
	for {
		call := m.service.Groups.List()
		if opts.Domain != "" {
			call = call.Domain(opts.Domain)
		}
		if opts.Customer != "" {
			call = call.Customer(opts.Customer)
		}
		if opts.Query != "" {
			call = call.Query(opts.Query)
		}
		if opts.MaxResults > 0 {
			call = call.MaxResults(opts.MaxResults)
		}
		if pageToken != "" {
			call = call.PageToken(pageToken)
		}
		if opts.OrderBy != "" {
			call = call.OrderBy(opts.OrderBy)
		}
		if opts.Fields != "" {
			call = call.Fields(googleapi.Field(opts.Fields))
		}

		result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*adminapi.Groups, error) {
			return call.Do()
		})
		if err != nil {
			return nil, err
		}

		allGroups = append(allGroups, convertGroups(result)...)
		if !opts.Paginate || result.NextPageToken == "" {
			return &types.GroupsListResponse{
				Groups:        allGroups,
				NextPageToken: result.NextPageToken,
			}, nil
		}
		pageToken = result.NextPageToken
	}
}

type GetGroupOptions struct {
	Fields string
}

func (m *Manager) GetGroup(ctx context.Context, reqCtx *types.RequestContext, groupKey string, opts *GetGroupOptions) (*types.Group, error) {
	call := m.service.Groups.Get(groupKey)
	if opts != nil && opts.Fields != "" {
		call = call.Fields(googleapi.Field(opts.Fields))
	}
	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*adminapi.Group, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}
	converted := convertGroup(result)
	return &converted, nil
}

func (m *Manager) CreateGroup(ctx context.Context, reqCtx *types.RequestContext, req *types.CreateGroupRequest) (*types.Group, error) {
	group := &adminapi.Group{
		Email:       req.Email,
		Name:        req.Name,
		Description: req.Description,
	}
	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*adminapi.Group, error) {
		return m.service.Groups.Insert(group).Do()
	})
	if err != nil {
		return nil, err
	}
	converted := convertGroup(result)
	return &converted, nil
}

func (m *Manager) DeleteGroup(ctx context.Context, reqCtx *types.RequestContext, groupKey string) error {
	_, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (struct{}, error) {
		return struct{}{}, m.service.Groups.Delete(groupKey).Do()
	})
	return err
}

func (m *Manager) UpdateGroup(ctx context.Context, reqCtx *types.RequestContext, groupKey string, req *types.UpdateGroupRequest) (*types.Group, error) {
	group := &adminapi.Group{}
	fieldMask := []string{}

	if req.Name != nil {
		group.Name = *req.Name
		fieldMask = append(fieldMask, "name")
	}
	if req.Description != nil {
		group.Description = *req.Description
		fieldMask = append(fieldMask, "description")
	}
	if len(fieldMask) == 0 {
		return m.GetGroup(ctx, reqCtx, groupKey, nil)
	}

	call := m.service.Groups.Patch(groupKey, group)

	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*adminapi.Group, error) {
		return call.Do()
	})
	if err != nil {
		return nil, err
	}
	converted := convertGroup(result)
	return &converted, nil
}

type ListMembersOptions struct {
	MaxResults int64
	PageToken  string
	Roles      string
	Fields     string
	Paginate   bool
}

func (m *Manager) ListMembers(ctx context.Context, reqCtx *types.RequestContext, groupKey string, opts *ListMembersOptions) (*types.MembersListResponse, error) {
	var allMembers []types.Member
	pageToken := opts.PageToken

	for {
		call := m.service.Members.List(groupKey)
		if opts.MaxResults > 0 {
			call = call.MaxResults(opts.MaxResults)
		}
		if pageToken != "" {
			call = call.PageToken(pageToken)
		}
		if opts.Roles != "" {
			call = call.Roles(opts.Roles)
		}
		if opts.Fields != "" {
			call = call.Fields(googleapi.Field(opts.Fields))
		}

		result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*adminapi.Members, error) {
			return call.Do()
		})
		if err != nil {
			return nil, err
		}

		allMembers = append(allMembers, convertMembers(result)...)
		if !opts.Paginate || result.NextPageToken == "" {
			return &types.MembersListResponse{
				Members:       allMembers,
				NextPageToken: result.NextPageToken,
			}, nil
		}
		pageToken = result.NextPageToken
	}
}

func (m *Manager) AddMember(ctx context.Context, reqCtx *types.RequestContext, groupKey string, req *types.AddMemberRequest) (*types.Member, error) {
	member := &adminapi.Member{
		Email: req.Email,
		Role:  req.Role,
	}
	if member.Role == "" {
		member.Role = "MEMBER"
	}
	result, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (*adminapi.Member, error) {
		return m.service.Members.Insert(groupKey, member).Do()
	})
	if err != nil {
		return nil, err
	}
	converted := convertMember(result)
	return &converted, nil
}

func (m *Manager) RemoveMember(ctx context.Context, reqCtx *types.RequestContext, groupKey, memberKey string) error {
	_, err := api.ExecuteWithRetry(ctx, m.client, reqCtx, func() (struct{}, error) {
		return struct{}{}, m.service.Members.Delete(groupKey, memberKey).Do()
	})
	return err
}

func convertUsers(users *adminapi.Users) []types.User {
	if users == nil || users.Users == nil {
		return []types.User{}
	}
	result := make([]types.User, len(users.Users))
	for i, user := range users.Users {
		result[i] = convertUser(user)
	}
	return result
}

func convertGroups(groups *adminapi.Groups) []types.Group {
	if groups == nil || groups.Groups == nil {
		return []types.Group{}
	}
	result := make([]types.Group, len(groups.Groups))
	for i, group := range groups.Groups {
		result[i] = convertGroup(group)
	}
	return result
}

func convertUser(user *adminapi.User) types.User {
	if user == nil {
		return types.User{}
	}
	fullName := ""
	given := ""
	family := ""
	if user.Name != nil {
		fullName = user.Name.FullName
		given = user.Name.GivenName
		family = user.Name.FamilyName
	}
	return types.User{
		ID:               user.Id,
		PrimaryEmail:     user.PrimaryEmail,
		Name:             types.UserName{GivenName: given, FamilyName: family, FullName: fullName},
		IsAdmin:          user.IsAdmin,
		IsDelegatedAdmin: user.IsDelegatedAdmin,
		Suspended:        user.Suspended,
		CreationTime:     user.CreationTime,
		LastLoginTime:    user.LastLoginTime,
	}
}

func convertGroup(group *adminapi.Group) types.Group {
	if group == nil {
		return types.Group{}
	}
	count := 0
	if group.DirectMembersCount != 0 {
		count = int(group.DirectMembersCount)
	}
	return types.Group{
		ID:                 group.Id,
		Email:              group.Email,
		Name:               group.Name,
		Description:        group.Description,
		AdminCreated:       group.AdminCreated,
		DirectMembersCount: count,
	}
}

func boolPtr(value bool) *bool {
	return &value
}

func convertMembers(members *adminapi.Members) []types.Member {
	if members == nil || members.Members == nil {
		return []types.Member{}
	}
	result := make([]types.Member, len(members.Members))
	for i, member := range members.Members {
		result[i] = convertMember(member)
	}
	return result
}

func convertMember(member *adminapi.Member) types.Member {
	if member == nil {
		return types.Member{}
	}
	return types.Member{
		ID:     member.Id,
		Email:  member.Email,
		Role:   member.Role,
		Type:   member.Type,
		Status: member.Status,
	}
}
