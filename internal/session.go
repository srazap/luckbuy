package internal

import (
	"github.com/srazap/luckbuy/db"
	"github.com/srazap/luckbuy/models"
)

func newSession(email string) (string, error) {
	session := models.NewSession(email)

	tx := db.GetInstance().Db.Begin()
	if err := tx.FirstOrCreate(&session).Error; err != nil {
		tx.Rollback()
		return "", err
	}

	tx.Commit()

	return session.SessionId, nil
}

func removeSession(sessionId string) error {
	tx := db.GetInstance().Db.Begin()
	if err := tx.Unscoped().Where("session_id=?", sessionId).Delete(&models.Session{}).Error; err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func GetEmailBySessionId(sId string) (string, error) {
	s := new(models.Session)
	if err := db.GetInstance().Db.Where("session_id=?", sId).First(&s).Error; err != nil {
		return "", err
	}
	return s.Email, nil
}

func ValidateSession(sId string) (bool, error) {
	email, err := GetEmailBySessionId(sId)
	if err != nil {
		return false, err
	}
	return email != "", nil
}
