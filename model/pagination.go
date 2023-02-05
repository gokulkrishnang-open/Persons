package model

type Pagination struct {
	ResultsPerPage string `json:"results_per_page"`
	Offset         string
	Page           string `json:"page"`
}
