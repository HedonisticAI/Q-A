package app

import (
	"golangqatestdesu/config"
	a_usecase "golangqatestdesu/internal/answers/usecase"
	q_usecase "golangqatestdesu/internal/questions/usecase"
	httpserver "golangqatestdesu/pkg/http_server"
	"golangqatestdesu/pkg/logger"
	"golangqatestdesu/pkg/postgresql"
)

type App struct {
	Server httpserver.HttpServer
	Ans    a_usecase.AnswersRepo
	Q      q_usecase.QuestionsRepo
	Logger logger.Logger
}

func NewApp(Cfg config.Config, Logger logger.Logger) App {
	HttpServ := httpserver.NewServer(Cfg)

	DB, err := postgresql.NewDB(Cfg)
	if err != nil {
		Logger.Error(err)
	}
	Ans := a_usecase.NewAnsRepo(DB, Logger)
	Q := q_usecase.NewQuesRepo(DB, Logger)

	return App{Server: *HttpServ, Logger: Logger, Ans: *Ans, Q: *Q}
}

func (A *App) Run() {}
