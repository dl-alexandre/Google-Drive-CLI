package admin

import (
	"testing"

	adminapi "google.golang.org/api/admin/directory/v1"
)

func TestConvertUsers(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		if got := convertUsers(nil); len(got) != 0 {
			t.Fatalf("expected empty slice, got %d", len(got))
		}
	})

	t.Run("empty users", func(t *testing.T) {
		if got := convertUsers(&adminapi.Users{Users: []*adminapi.User{}}); len(got) != 0 {
			t.Fatalf("expected empty slice, got %d", len(got))
		}
	})

	t.Run("missing name", func(t *testing.T) {
		users := &adminapi.Users{Users: []*adminapi.User{{Id: "1", PrimaryEmail: "a@example.com"}}}
		got := convertUsers(users)
		if got[0].Name.FullName != "" || got[0].Name.GivenName != "" || got[0].Name.FamilyName != "" {
			t.Fatalf("expected empty name fields")
		}
	})

	t.Run("full fields", func(t *testing.T) {
		users := &adminapi.Users{Users: []*adminapi.User{{
			Id:           "1",
			PrimaryEmail: "a@example.com",
			Name:         &adminapi.UserName{FullName: "A B", GivenName: "A", FamilyName: "B"},
			IsAdmin:      true,
			Suspended:    true,
		}}}
		got := convertUsers(users)
		if got[0].PrimaryEmail != "a@example.com" {
			t.Fatalf("expected email to match")
		}
		if got[0].Name.FullName != "A B" {
			t.Fatalf("expected full name to match")
		}
		if !got[0].IsAdmin || !got[0].Suspended {
			t.Fatalf("expected admin and suspended to be true")
		}
	})
}

func TestConvertGroups(t *testing.T) {
	t.Run("nil input", func(t *testing.T) {
		if got := convertGroups(nil); len(got) != 0 {
			t.Fatalf("expected empty slice, got %d", len(got))
		}
	})

	t.Run("valid member count", func(t *testing.T) {
		groups := &adminapi.Groups{Groups: []*adminapi.Group{{
			Id:                 "1",
			Email:              "g@example.com",
			Name:               "Group",
			DirectMembersCount: 42,
		}}}
		got := convertGroups(groups)
		if got[0].DirectMembersCount != 42 {
			t.Fatalf("expected count 42, got %d", got[0].DirectMembersCount)
		}
	})
}

func TestConvertUser(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		got := convertUser(nil)
		if got.ID != "" || got.PrimaryEmail != "" {
			t.Fatalf("expected empty user")
		}
	})

	t.Run("with name", func(t *testing.T) {
		user := &adminapi.User{
			Id:           "1",
			PrimaryEmail: "a@example.com",
			Name:         &adminapi.UserName{FullName: "A B", GivenName: "A", FamilyName: "B"},
		}
		got := convertUser(user)
		if got.Name.FullName != "A B" || got.Name.GivenName != "A" || got.Name.FamilyName != "B" {
			t.Fatalf("expected name fields")
		}
	})
}

func TestConvertGroup(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		got := convertGroup(nil)
		if got.ID != "" || got.Email != "" {
			t.Fatalf("expected empty group")
		}
	})

	t.Run("with count", func(t *testing.T) {
		group := &adminapi.Group{
			Id:                 "1",
			Email:              "g@example.com",
			Name:               "Group",
			DirectMembersCount: 3,
		}
		got := convertGroup(group)
		if got.DirectMembersCount != 3 {
			t.Fatalf("expected count 3")
		}
	})
}
