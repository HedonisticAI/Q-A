package q_usecase

import (
	"context"
	"golangqatestdesu/internal/questions"
	"golangqatestdesu/pkg/logger"
	"golangqatestdesu/pkg/postgresql"
	"time"

	"gorm.io/gorm"
)

type Answer struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	Question_id uint      `gorm:"column:q_id" json:"question_id"`
	User_id     string    `gorm:"column:user_id" json:"user_id"`
	Text        string    `gorm:"column:a_text" json:"text"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"omitempty"`
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
		if rows == 0 {
			Q.Logger.Error(gorm.ErrRecordNotFound)
			return gorm.ErrRecordNotFound
		}
		Q.Logger.Error(err)
		return err
	}
	Q.Logger.Debug("Deleted entry from questions table")
	_, err = gorm.G[Answer](Q.DB.DB).Where("q_id = ?", ID).Delete(ctx)
	if err != nil && err != gorm.ErrRecordNotFound {
		Q.Logger.Error(err)
		return err
	}
	Q.Logger.Debug("Deleted all matchind questions")
	Q.Logger.Info("Question Deleted")
	return nil
}

func (Q *QuestionsRepo) GetByID(ID int) (*questions.Question, []Answer, error) {
	Q.Logger.Info("Getting question by ID")
	var Answers []Answer
	var Question questions.Question
	res := Q.DB.DB.First(&Question, ID)
	if res.Error != nil {
		Q.Logger.Error(res.Error)
		return nil, nil, res.Error
	}
	Q.Logger.Debug("Getting answer for Question")
	rows := Q.DB.DB.Where("q_id = ?", ID).Find(&Answers)
	if rows.Error != nil || rows.RowsAffected == 0 {
		Q.Logger.Error(rows.Error)
		if rows.Error != gorm.ErrRecordNotFound {
			Q.Logger.Debug("No related answers found")
			return &Question, nil, nil
		}
		return nil, nil, rows.Error
	}
	Q.Logger.Debug("Got answers")
	Q.Logger.Debug(Answers)
	return &Question, Answers, nil
}

func (Q *QuestionsRepo) GetAll() ([]questions.Question, error) {
	Q.Logger.Info("Retrieving all questions")
	var Questions []questions.Question
	res := Q.DB.DB.Find(&Questions)
	if res.Error != nil {
		Q.Logger.Debug(res.Error)
		return nil, res.Error
	}
	return Questions, nil
}
