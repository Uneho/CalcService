package calculator

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type CalcRequest struct {
	Expression string `json:"expression"`
}

type CalcResponse struct {
	Result string `json:"result,omitempty"`
	Error  string `json:"error,omitempty"`
}

func CalculateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, `{"error":"Method not allowed"}`, http.StatusMethodNotAllowed)
		return
	}

	var req CalcRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, `{"error":"Invalid request body"}`, http.StatusUnprocessableEntity)
		return
	}

	result, err := Calc(req.Expression)
	if err != nil {
		message := GetErrorMessage(err)
		if err == ErrInvalidExpression {
			http.Error(w, `{"error":"`+message+`"}`, http.StatusUnprocessableEntity)
		} else {
			http.Error(w, `{"error":"`+message+`"}`, http.StatusInternalServerError)
		}
		return
	}

	resp := CalcResponse{Result: strconv.FormatFloat(result, 'f', 2, 64)}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
