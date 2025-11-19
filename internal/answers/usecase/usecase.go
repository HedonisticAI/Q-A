package a_usecase

import (
	"encoding/json"
	"golangqatestdesu/internal/answers"
	"golangqatestdesu/pkg/logger"
	"golangqatestdesu/pkg/postgresql"
	"net/http"
	"strconv"
)

type AnswerUsecase struct {
	AnsRepo *AnswersRepo
	Logger  logger.Logger
}

type AnswerCreateRequest struct {
	UserID string `json:"UserID"`
	Text   string `json:"Text"`
}

func NewAnsUseCase(DB *postgresql.DB, Logger logger.Logger) answers.Answers {
	AnsRepo := NewAnsRepo(DB, Logger)
	return &AnswerUsecase{AnsRepo: AnsRepo, Logger: Logger}
}

func (AnswerUsecase *AnswerUsecase) AddAnswer(w http.ResponseWriter, r *http.Request) {
	var Ans AnswerCreateRequest
	QuestionID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		AnswerUsecase.Logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	decoder := json.NewDecoder(r.Body)
	AnswerUsecase.Logger.Info("Got create answer request")
	err = decoder.Decode(&Ans)
	if err != nil {
		AnswerUsecase.Logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	NewAnswer := &answers.Answer{Text: Ans.Text, User_id: Ans.UserID}
	Id, err := AnswerUsecase.AnsRepo.Create(NewAnswer, uint(QuestionID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	AnswerUsecase.Logger.Debug("Created DB enrty with id ", Id)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created Question with ID: " + strconv.Itoa(int(Id)))
}
func (AnswerUsecase *AnswerUsecase) GetAnswer(w http.ResponseWriter, r *http.Request) {
	AnswerUsecase.Logger.Info("Got Get Answer req")
	GetID, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		AnswerUsecase.Logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Answer, err := AnswerUsecase.AnsRepo.GetByID(uint(GetID))
	if err != nil {
		AnswerUsecase.Logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Answer)
}
func (AnswerUsecase *AnswerUsecase) Delete(w http.ResponseWriter, r *http.Request) {
	AnswerUsecase.Logger.Info("Got Answer Delete req")
	DeleteID, err := strconv.Atoi(r.PathValue("id"))
	AnswerUsecase.Logger.Debug(DeleteID)
	if err != nil {
		AnswerUsecase.Logger.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err = AnswerUsecase.AnsRepo.DeleteByID(uint(DeleteID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
}
