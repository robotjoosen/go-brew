package brew

import "time"

type Pour struct {
	Duration time.Duration `json:"duration"`
	Grams    int           `json:"grams"`
}

type Recipe struct {
	Beans  string `json:"beans"`
	Ratio  string `json:"ratio"`
	Coffee int    `json:"coffee"`
	Water  int    `json:"water"`
	Schema []Pour `json:"schema"`
}
