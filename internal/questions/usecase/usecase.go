package q_usecase

import (
	"encoding/json"
	"golangqatestdesu/internal/questions"
	"golangqatestdesu/pkg/logger"
	"golangqatestdesu/pkg/postgresql"
	"net/http"
	"strconv"
)

type QuestionsUsecase struct {
	QuesRepo *QuestionsRepo
	Logger   logger.Logger
}

type QuestionCreateRequest struct {
	Text string `json:"Text"`
}

type GetAllResp struct {
	Questions []questions.Question `json:"Questions"`
}

type GetByIDResponse struct {
	Question *questions.Question `json:"Question"`
	Answers  []Answer            `json:"Answers,omitempty"`
}

func NewQueUseCase(DB *postgresql.DB, Logger logger.Logger) questions.Questions {
	QuesRepo := NewQuesRepo(DB, Logger)
	return &QuestionsUsecase{QuesRepo: QuesRepo, Logger: Logger}
}

func (QuestionsUsecase *QuestionsUsecase) GetByID(w http.ResponseWriter, r *http.Request) {
	QuestionsUsecase.Logger.Info("Got Get Question Request")
	var Res GetByIDResponse
	GetID, err := strconv.Atoi(r.PathValue("id"))
	QuestionsUsecase.Logger.Debug(GetID)
	if err != nil {
		QuestionsUsecase.Logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Res.Question, Res.Answers, err = QuestionsUsecase.QuesRepo.GetByID(GetID)
	if err != nil {
		QuestionsUsecase.Logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Res)
}
func (QuestionsUsecase *QuestionsUsecase) GetAll(w http.ResponseWriter, r *http.Request) {
	Questions, err := QuestionsUsecase.QuesRepo.GetAll()
	if err != nil {
		QuestionsUsecase.Logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Res := &GetAllResp{Questions: Questions}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Res)
}
func (QuestionsUsecase *QuestionsUsecase) Create(w http.ResponseWriter, r *http.Request) {
	var QCreate QuestionCreateRequest

	decoder := json.NewDecoder(r.Body)
	QuestionsUsecase.Logger.Info("Got create Question request")
	err := decoder.Decode(&QCreate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Ques := &questions.Question{Text: QCreate.Text}
	Id, err := QuestionsUsecase.QuesRepo.Create(Ques)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	QuestionsUsecase.Logger.Debug("Created DB enrty with id ", Id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created Question with ID: " + strconv.Itoa(int(Id)))
}
func (QuestionsUsecase *QuestionsUsecase) Delete(w http.ResponseWriter, r *http.Request) {
	QuestionsUsecase.Logger.Info("Got Question Delete req")
	DeleteID, err := strconv.Atoi(r.PathValue("id"))
	QuestionsUsecase.Logger.Debug(DeleteID)
	if err != nil {
		QuestionsUsecase.Logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = QuestionsUsecase.QuesRepo.Delete(uint(DeleteID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
