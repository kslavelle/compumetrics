package identity

// RegisterUser represents a user who is signing up
type RegisterUser struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

// User stores email password and a hash of password.
type User struct {
	Email        string
	Password     string
	PasswordHash string
}
