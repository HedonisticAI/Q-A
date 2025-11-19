package answers

import "time"

type Answer struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Question_id uint      `gorm:"column:q_id" json:"question_id"`
	User_id     string    `gorm:"column:user_id" json:"user_id"`
	Text        string    `gorm:"column:a_text" json:"text"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"omitempty"`
}
