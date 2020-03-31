package handlers

import (
	"net/http"
	"golang-fifa-world-cup-web-service/data"
)

// RootHandler returns an empty body status code
func RootHandler(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusNoContent)
	return
}

// ListWinners returns winners from the list
func ListWinners(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	year := req.URL.Query().Get("year")
	
	winners, err := data.ListAllJSON()
	
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	
	if year == ""{
		res.Write(winners)
	} else{
		filteredWinners, err := data.ListAllByYear(year)
		if err != nil{
			res.WriteHeader(http.StatusBadRequest)
			return
		}
		res.Write(filteredWinners)
		return
	}

}

// AddNewWinner adds new winner to the list
func AddNewWinner(res http.ResponseWriter, req *http.Request) {
	accessToken := req.Header.Get("X-ACCESS-TOKEN")
	isTokenValid := data.IsAccessTokenValid(accessToken)

	if !isTokenValid{
		res.WriteHeader(http.StatusUnauthorized)
	}

	err := data.AddNewWinner(req.Body)

	if err != nil{
		res.WriteHeader(http.StatusUnprocessableEntity)
		return
	}

	res.WriteHeader(http.StatusCreated)
	return
}

// WinnersHandler is the dispatcher for all /winners URL
func WinnersHandler(res http.ResponseWriter, req *http.Request) {
	switch method := req.Method; method{
	case http.MethodGet:
		ListWinners(res, req)
	case http.MethodPost:
		AddNewWinner(res, req)
	default:
		res.WriteHeader(http.StatusMethodNotAllowed)
	}
}
