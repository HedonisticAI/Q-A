package answers

import "net/http"

type Answers interface {
	AddAnswer(w http.ResponseWriter, r *http.Request)
	GetAnswer(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}
