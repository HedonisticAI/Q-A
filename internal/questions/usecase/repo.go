package q_usecase

import (
	"context"
	"golangqatestdesu/internal/questions"
	"golangqatestdesu/pkg/logger"
	"golangqatestdesu/pkg/postgresql"
	"time"

	"gorm.io/gorm"
)

type answer struct {
	ID          uint      `gorm:"primaryKey"`
	Question_id uint      `gorm:"column:q_id"`
	User_id     string    `gorm:"column:user_id"`
	Text        string    `gorm:"column:a_text"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

type QuestionsRepo struct {
	DB     *postgresql.DB
	Logger logger.Logger
}

func NewQuesRepo(DB *postgresql.DB, LOgger logger.Logger) *QuestionsRepo {
	return &QuestionsRepo{DB: DB, Logger: LOgger}
}

func (Q *QuestionsRepo) Create(Question *questions.Question) (uint, error) {
	result := Q.DB.DB.Create(Question)
	if result.Error != nil {
		Q.Logger.Error(result.Error)
		return 0, result.Error
	}
	Q.Logger.Info("Created new Question")
	return Question.ID, nil
}

func (Q *QuestionsRepo) Delete(ID uint) error {
	ctx := context.Background()
	rows, err := gorm.G[questions.Question](Q.DB.DB).Where("id = ?", ID).Delete(ctx)
	if err != nil || rows == 0 {
		Q.Logger.Error(err)
		return err
	}
	Q.Logger.Debug("Deleted entry from questions table")
	_, err = gorm.G[answer](Q.DB.DB).Where("q_id = ?", ID).Delete(ctx)
	if err != nil && err != gorm.ErrRecordNotFound {
		Q.Logger.Error(err)
		return err
	}
	Q.Logger.Debug("Deleted all matchind questions")
	Q.Logger.Info("Question Deleted")
	return nil
}

func (Q *QuestionsRepo) GetByID(ID uint) questions.Question {
	Q.Logger.Info("Getting question by ID")
	var Question questions.Question
	Q.DB.DB.First(&Question, ID)
	return Question
}

func (Q *QuestionsRepo) GetAll() []questions.Question {
	Q.Logger.Info("Retrieving all questions")
	var Questions []questions.Question
	Q.DB.DB.Find(&Questions)
	return Questions
}
