package a_usecase

import (
	"golangqatestdesu/pkg/logger"
	"golangqatestdesu/pkg/postgresql"
)

type AnswerUsecase struct {
	AnsRepo AnswersRepo
	Logger  logger.Logger
}

func NewAnsUseCase(DB *postgresql.DB, Logger logger.Logger) *AnswerUsecase {
	AnsRepo := NewAnsRepo(DB, Logger)
	return &AnswerUsecase{AnsRepo: *AnsRepo, Logger: Logger}
}
