package server

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/musafir-V/log-grepper/internal/model"
	"github.com/musafir-V/log-grepper/internal/service"
	"log/slog"
	"net/http"
	"time"
)

type SearchHandler interface {
	GetMatchingLogs(w http.ResponseWriter, r *http.Request)
}

type searchHandler struct {
	grepper service.LogGrepperService
}

func NewSearchHandler(grepper service.LogGrepperService) SearchHandler {
	return &searchHandler{
		grepper: grepper,
	}
}

func (h searchHandler) GetMatchingLogs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	req := getRequestFromQueryParams(r)
	fmt.Println(req)
	if req.SearchKeyword == "" || req.From == "" || req.To == "" {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	slog.Info("Request received", slog.String("search_keyword", req.SearchKeyword), slog.String("from", req.From), slog.String("to", req.To))
	err := validateRequest(req)

	if err != nil {
		slog.Error("Invalid request", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	to, _ := time.Parse(time.DateOnly, req.To)     // Ignoring error as it is already validated
	from, _ := time.Parse(time.DateOnly, req.From) // Ignoring error as it is already validated
	greppedResp := h.grepper.GrepLogs(to, from, req.SearchKeyword)

	retResp := model.Response{
		Code:    200,
		Message: "success",
		Data:    greppedResp,
	}
	respBytes, err := json.Marshal(retResp)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(respBytes)
}

func validateRequest(req *model.Request) error {
	from := req.From
	fromTimestamp, err := time.Parse(time.DateOnly, from)
	if err != nil {
		return err
	}

	to := req.To
	toTimestamp, err := time.Parse(time.DateOnly, to)
	if err != nil {
		return err
	}

	if fromTimestamp.After(toTimestamp) {
		return errors.New("from timestamp cannot be after to timestamp")
	}

	// Setting a hard limit at 7 days, the time range can be of only 7 days.
	if toTimestamp.Sub(fromTimestamp) > 7*24*time.Hour {
		return errors.New("time range cannot be more than 7 days")
	}
	return nil
}

func getRequestFromQueryParams(r *http.Request) *model.Request {
	var req model.Request
	req.To = r.URL.Query().Get("to")
	req.From = r.URL.Query().Get("from")
	req.SearchKeyword = r.URL.Query().Get("search_keyword")
	return &req
}
