package types

import "testing"

func TestMembersListRows(t *testing.T) {
	resp := &MembersListResponse{
		Members: []Member{
			{Email: "a@example.com", Role: "MEMBER", Type: "USER", Status: "ACTIVE"},
		},
	}
	rows := resp.Rows()
	if len(rows) != 1 {
		t.Fatalf("expected 1 row")
	}
	if rows[0][0] != "a@example.com" || rows[0][1] != "MEMBER" {
		t.Fatalf("unexpected row: %#v", rows[0])
	}
}
