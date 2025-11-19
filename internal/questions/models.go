package questions

import "time"

type Question struct {
	ID        uint      `gorm:"primaryKey"`
	Text      string    `gorm:"column:q_text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
