package model

type Pagination struct {
	ResultsPerPage string `json:"results_per_page"`
	Page           string `json:"page"`
}
