package user

import (
	"testing"

	"github.com/hrs-o/docker-go/domain/shared"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		id        UserID
		firstName string
		lastName  string
		wantErr   bool
	}{
		{UserID("some-id"), "John", "Doe", false},
		{UserID("some-id"), "", "Doe", true},
		{UserID("some-id"), "John", "", true},
		{UserID("some-id"), "verylongfirstnameverylongfirstname", "Doe", true},
		{UserID("some-id"), "John", "verylonglastnameverylonglastname", true},
	}

	for _, test := range tests {
		user, err := NewUser(test.id, test.firstName, test.lastName)
		if (err != nil) != test.wantErr {
			t.Errorf("NewUser(%q, %q, %q) error = %v, wantErr %v", test.id, test.firstName, test.lastName, err, test.wantErr)
			continue
		}
		if err == nil && (user.ID != test.id || user.FirstName != test.firstName || user.LastName != test.lastName) {
			t.Errorf("NewUser(%q, %q, %q) = %v, want %v", test.id, test.firstName, test.lastName, user, test)
		}
	}
}

func TestGetCreatedAt(t *testing.T) {
	user := &User{
		ID:        UserID("some-id"),
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: shared.NewCreatedAt(),
	}

	got := user.GetCreatedAt()
	if got.Value().IsZero() {
		t.Errorf("GetCreatedAt() = %v, want non-zero time", got)
	}
}
