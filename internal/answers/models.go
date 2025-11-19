package answers

import "time"

type Answer struct {
	ID          uint      `gorm:"primaryKey"`
	Question_id uint      `gorm:"column:q_id"`
	User_id     string    `gorm:"column:user_id"`
	Text        string    `gorm:"column:a_text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}
