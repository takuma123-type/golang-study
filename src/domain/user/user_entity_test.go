package user

import (
	"testing"

	"github.com/google/uuid"
	"github.com/hrs-o/docker-go/domain/shared"
	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		firstName string
		lastName  string
		wantErr   bool
	}{
		{"John", "Doe", false},
		{"", "Doe", true},
		{"John", "", true},
		{"verylongfirstnameverylongfirstname", "Doe", true},
		{"John", "verylonglastnameverylonglastname", true},
	}

	for _, test := range tests {
		user, err := NewUser(NewUserID(), test.firstName, test.lastName)
		if test.wantErr {
			assert.Error(t, err)
			continue
		}
		assert.NoError(t, err)
		assert.Equal(t, test.firstName, user.FirstName)
		assert.Equal(t, test.lastName, user.LastName)
	}
}

func NewUserID() UserID {
	return UserID(uuid.New().String())
}

func TestGetCreatedAt(t *testing.T) {
	user := &User{
		ID:        UserID("some-id"),
		FirstName: "John",
		LastName:  "Doe",
		CreatedAt: shared.NewCreatedAt(),
	}

	got := user.GetCreatedAt()
	assert.NotZero(t, got.Value())
}
