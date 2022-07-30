package configs

import "os"

// NewAdmin returns an Admin object
func NewAdmin() Admin {
	return Admin{
		UserName:  os.Getenv("USERNAME"),
		Password:  os.Getenv("PASSWORD"),
		SecretKey: os.Getenv("SECRET_KEY"),
	}
}

// GetSecretKey returns the jwt secret key
func (a Admin) GetSecretKey() string {
	return a.SecretKey
}

// GetAdminName returns the Username
func (a Admin) GetAdminName() string {
	return a.UserName
}

// GetAdminPwd returns the Password
func (a Admin) GetAdminPwd() string {
	return a.Password
}
