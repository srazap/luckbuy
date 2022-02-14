package internal

import (
	"errors"

	"github.com/srazap/luckbuy/db"
	"github.com/srazap/luckbuy/models"
	"gorm.io/gorm"
)

func Register(user *models.User) error {

	tx := db.GetInstance().Db.Begin()
	if err := tx.Create(&user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// update points of user who referred me
	if user.ReferralCode != "" {
		if err := addRefferralPoints(user.ReferralCode, tx); err != nil {
			tx.Rollback()
			return err
		}
	}

	tx.Commit()

	return nil
}

func Login(email, password string) (string, error) {

	user := new(models.User)

	if err := db.GetInstance().Db.Where("email=?", email).Find(&user).Error; err != nil {
		return "", err
	}

	if user.Password == password {
		return newSession(email)
	}

	return "", errors.New("wrong email or password")
}

func Logout(sessionId string) error {
	return removeSession(sessionId)
}

func MyPoints(email string) ([]models.User, int64, error) {
	var (
		user *models.User
		list []models.User
		err  error
	)

	mainDb := db.GetInstance().Db

	// query user for points
	if err = mainDb.Where("email=?", email).Select("points").Find(&user).Error; err != nil {
		return nil, 0, err
	}

	// query user summary
	query := `SELECT * FROM users WHERE referral_code = ? UNION
			SELECT * FROM users WHERE referral_code IN (SELECT email FROM users WHERE referral_code = ?)`

	if err = mainDb.Raw(query, email, email).Scan(&list).Error; err != nil {
		return nil, 0, err
	}

	return list, user.Points, nil
}

func addRefferralPoints(email string, tx *gorm.DB) error {
	query := `UPDATE users SET points = points + 1 WHERE email = ? OR email IN (SELECT referral_code FROM users WHERE email = ?)`
	if err := tx.Exec(query, email, email).Error; err != nil {
		return err
	}
	return nil
}
