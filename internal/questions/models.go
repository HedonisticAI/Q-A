package questions

import "time"

type Question struct {
	ID        uint      `gorm:"primaryKey" json:"id,omitempty"`
	Text      string    `gorm:"column:q_text" json:"text"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created"`
}
