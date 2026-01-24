package types

import "fmt"

type User struct {
	ID               string   `json:"id"`
	PrimaryEmail     string   `json:"primaryEmail"`
	Name             UserName `json:"name"`
	IsAdmin          bool     `json:"isAdmin,omitempty"`
	IsDelegatedAdmin bool     `json:"isDelegatedAdmin,omitempty"`
	Suspended        bool     `json:"suspended,omitempty"`
	CreationTime     string   `json:"creationTime,omitempty"`
	LastLoginTime    string   `json:"lastLoginTime,omitempty"`
}

type UserName struct {
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
	FullName   string `json:"fullName"`
}

type CreateUserRequest struct {
	Email      string `json:"email"`
	GivenName  string `json:"givenName"`
	FamilyName string `json:"familyName"`
	Password   string `json:"password"`
}

type UpdateUserRequest struct {
	GivenName   *string `json:"givenName,omitempty"`
	FamilyName  *string `json:"familyName,omitempty"`
	Suspended   *bool   `json:"suspended,omitempty"`
	OrgUnitPath *string `json:"orgUnitPath,omitempty"`
}

type UsersListResponse struct {
	Users         []User `json:"users"`
	NextPageToken string `json:"nextPageToken,omitempty"`
}

func (r *UsersListResponse) Headers() []string {
	return []string{"Email", "Name", "Admin", "Suspended"}
}

func (r *UsersListResponse) Rows() [][]string {
	rows := make([][]string, len(r.Users))
	for i, user := range r.Users {
		adminStatus := "No"
		if user.IsAdmin {
			adminStatus = "Yes"
		}
		suspendedStatus := "No"
		if user.Suspended {
			suspendedStatus = "Yes"
		}
		rows[i] = []string{
			user.PrimaryEmail,
			user.Name.FullName,
			adminStatus,
			suspendedStatus,
		}
	}
	return rows
}

func (r *UsersListResponse) EmptyMessage() string {
	return "No users found"
}

type Group struct {
	ID                 string `json:"id"`
	Email              string `json:"email"`
	Name               string `json:"name"`
	Description        string `json:"description,omitempty"`
	AdminCreated       bool   `json:"adminCreated,omitempty"`
	DirectMembersCount int    `json:"directMembersCount,omitempty"`
}

type CreateGroupRequest struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
}

type UpdateGroupRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
}

type GroupsListResponse struct {
	Groups        []Group `json:"groups"`
	NextPageToken string  `json:"nextPageToken,omitempty"`
}

func (r *GroupsListResponse) Headers() []string {
	return []string{"Email", "Name", "Description", "Members"}
}

func (r *GroupsListResponse) Rows() [][]string {
	rows := make([][]string, len(r.Groups))
	for i, group := range r.Groups {
		rows[i] = []string{
			group.Email,
			group.Name,
			truncateAdminText(group.Description, 50),
			fmt.Sprintf("%d", group.DirectMembersCount),
		}
	}
	return rows
}

func (r *GroupsListResponse) EmptyMessage() string {
	return "No groups found"
}

func truncateAdminText(s string, max int) string {
	if len(s) <= max {
		return s
	}
	if max <= 3 {
		return s[:max]
	}
	return s[:max-3] + "..."
}

type Member struct {
	ID     string `json:"id"`
	Email  string `json:"email"`
	Role   string `json:"role"`
	Type   string `json:"type"`
	Status string `json:"status,omitempty"`
}

type MembersListResponse struct {
	Members       []Member `json:"members"`
	NextPageToken string   `json:"nextPageToken,omitempty"`
}

func (r *MembersListResponse) Headers() []string {
	return []string{"Email", "Role", "Type", "Status"}
}

func (r *MembersListResponse) Rows() [][]string {
	rows := make([][]string, len(r.Members))
	for i, member := range r.Members {
		rows[i] = []string{
			member.Email,
			member.Role,
			member.Type,
			member.Status,
		}
	}
	return rows
}

func (r *MembersListResponse) EmptyMessage() string {
	return "No members found"
}

type AddMemberRequest struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}
