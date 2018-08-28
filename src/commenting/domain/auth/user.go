package auth

type User struct {
	UserID string
}

func NewUser(userID string) *User {
	return &User{
		UserID: userID,
	}
}
