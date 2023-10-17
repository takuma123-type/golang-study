package user

func GenWhenCreate(first, last string) (*User, error) {
	return NewUser(NewUserID(), first, last)
}

func GenForTest(id UserID, first, last string) (*User, error) {
	return NewUser(id, first, last)
}
