package models

type User struct {
	Email        string `json:"email" gorm:"primaryKey"`
	Password     string `json:"password"`
	ReferralCode string `json:"refferral_code"`
	Points       int64  `json:"points"`
}

func NewUser(email, password, refCode string) *User {
	return &User{
		Email:        email,
		Password:     password,
		ReferralCode: refCode,
	}
}
