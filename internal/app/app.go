package app

import (
	"golangqatestdesu/config"
	"golangqatestdesu/internal/answers"
	a_usecase "golangqatestdesu/internal/answers/usecase"
	"golangqatestdesu/internal/questions"
	q_usecase "golangqatestdesu/internal/questions/usecase"
	httpserver "golangqatestdesu/pkg/http_server"
	"golangqatestdesu/pkg/logger"
	"golangqatestdesu/pkg/postgresql"
)

type App struct {
	Server    httpserver.HttpServer
	Qusetions questions.Questions
	Answers   answers.Answers
	Logger    logger.Logger
}

func NewApp(Cfg config.Config, Logger logger.Logger) App {
	HttpServ := httpserver.NewServer(Cfg)
	Logger.Debug("Http Ready")
	DB, err := postgresql.NewDB(Cfg)
	if err != nil {
		Logger.Error(err)
	}
	Logger.Debug("DB ready")
	Ans := a_usecase.NewAnsUseCase(DB, Logger)
	Q := q_usecase.NewQueUseCase(DB, Logger)
	HttpServ.Map("GET /questions/{id}", Q.GetByID)
	HttpServ.Map("GET /questions/", Q.GetAll)
	HttpServ.Map("GET /answers/{id}", Ans.GetAnswer)
	HttpServ.Map("POST /questions/", Q.Create)
	HttpServ.Map("POST /questions/{id}/answers/", Ans.AddAnswer)
	HttpServ.Map("DELETE /answers/{id}", Ans.Delete)
	HttpServ.Map("DELETE /questions/{id}", Q.Delete)
	return App{Server: *HttpServ, Logger: Logger, Answers: Ans, Qusetions: Q}
}

func (A *App) Run() {
	A.Logger.Info("Running")
	A.Server.Run()
}
