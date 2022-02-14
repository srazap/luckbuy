package models

import (
	"encoding/base64"
	"fmt"
	"time"
)

const SessionLimit = 3600

type Session struct {
	Email      string `gorm:"primaryKey"`
	SessionId  string
	ExpiryTime time.Time
}

func NewSession(email string) *Session {

	// get current time
	ct := time.Now()

	// base64 encode current email + timestamp
	sessionId := base64.URLEncoding.EncodeToString([]byte(fmt.Sprintf("%s%d", email, ct.Unix())))

	return &Session{
		Email:      email,
		SessionId:  sessionId,
		ExpiryTime: ct.Add(time.Second * SessionLimit),
	}
}
