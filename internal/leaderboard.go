package internal

import (
	"github.com/srazap/luckbuy/db"
	"github.com/srazap/luckbuy/models"
)

func GetLeaderboard() ([]models.Leaderboard, error) {
	var board []models.Leaderboard

	query := `SELECT email, points, ROW_NUMBER () OVER (ORDER BY points desc) AS rank FROM users ORDER BY points desc`
	if err := db.GetInstance().Db.Raw(query).Scan(&board).Error; err != nil {
		return nil, err
	}

	return board, nil
}
