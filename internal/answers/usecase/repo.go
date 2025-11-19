package a_usecase

import (
	"context"
	"golangqatestdesu/internal/answers"
	"golangqatestdesu/pkg/logger"
	"golangqatestdesu/pkg/postgresql"
	"time"

	"gorm.io/gorm"
)

type AnswersRepo struct {
	Logger logger.Logger
	DB     *postgresql.DB
}

type question struct {
	ID        uint      `gorm:"primaryKey"`
	Text      string    `gorm:"column:q_text"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}

func NewAnsRepo(DB *postgresql.DB, Log logger.Logger) *AnswersRepo {
	return &AnswersRepo{DB: DB, Logger: Log}
}

func (A *AnswersRepo) DeleteByID(ID uint) error {
	ctx := context.Background()
	rows, err := gorm.G[answers.Answer](A.DB.DB).Where("id = ?", ID).Delete(ctx)
	if err != nil || rows == 0 {
		if rows == 0 {
			A.Logger.Error(gorm.ErrRecordNotFound)
			return gorm.ErrRecordNotFound
		}
		A.Logger.Error(err)
		return err
	}
	return nil
}

func (A *AnswersRepo) GetByID(ID uint) (*answers.Answer, error) {
	A.Logger.Info("Getting answer by ID")
	var Answer answers.Answer
	res := A.DB.DB.Find(&Answer, ID)
	if res.Error != nil {
		A.Logger.Error(res.Error)
		return nil, res.Error
	}
	return &Answer, nil
}

func (A *AnswersRepo) Create(Ans *answers.Answer, ID uint) (uint, error) {
	A.Logger.Info("Creating new answer")
	var Q question
	res := A.DB.DB.First(&Q, ID)
	if res.Error != nil || res.RowsAffected == 0 {
		A.Logger.Error(res.Error)
		if res.RowsAffected == 0 {
			A.Logger.Error(gorm.ErrRecordNotFound)
			return 0, gorm.ErrRecordNotFound
		}
		return 0, res.Error
	}
	Ans.Question_id = ID
	result := A.DB.DB.Create(Ans)
	if result.Error != nil {
		A.Logger.Error(result.Error)
		return 0, result.Error
	}
	return Ans.ID, nil
}
