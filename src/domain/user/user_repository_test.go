package user

import (
	"context"
	"errors"
	"testing"
)

func TestStore(t *testing.T) {
	mockRepo := &MockUserRepository{}
	ctx := context.Background()

	// エラーを返すモックを設定
	mockRepo.StoreErr = errors.New("some error")

	err := mockRepo.Store(ctx, &User{})
	if err == nil || err.Error() != "some error" {
		t.Errorf("expected error, got %v", err)
	}

	// エラーを返さないモックを設定
	mockRepo.StoreErr = nil

	err = mockRepo.Store(ctx, &User{})
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestFindAll(t *testing.T) {
	mockRepo := &MockUserRepository{}
	ctx := context.Background()

	// ユーザーを返すモックを設定
	mockRepo.FindAllUsers = []*User{{ID: UserID("some-id")}}

	users, err := mockRepo.FindAll(ctx)
	if err != nil || len(users) != 1 || users[0].ID != UserID("some-id") {
		t.Errorf("unexpected result: users = %v, error = %v", users, err)
	}

	// エラーを返すモックを設定
	mockRepo.FindAllErr = errors.New("some error")
	mockRepo.FindAllUsers = nil

	users, err = mockRepo.FindAll(ctx)
	if err == nil || err.Error() != "some error" {
		t.Errorf("expected error, got users = %v, error = %v", users, err)
	}
}

func TestFindByID(t *testing.T) {
	mockRepo := &MockUserRepository{}
	ctx := context.Background()

	// ユーザーを返すモックを設定
	mockRepo.FindByIDUser = &User{ID: UserID("some-id")}

	user, err := mockRepo.FindByID(ctx, UserID("some-id"))
	if err != nil || user.ID != UserID("some-id") {
		t.Errorf("unexpected result: user = %v, error = %v", user, err)
	}

	// エラーを返すモックを設定
	mockRepo.FindByIDErr = errors.New("some error")
	mockRepo.FindByIDUser = nil

	user, err = mockRepo.FindByID(ctx, UserID("some-id"))
	if err == nil || err.Error() != "some error" {
		t.Errorf("expected error, got user = %v, error = %v", user, err)
	}
}
