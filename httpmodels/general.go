package httpmodels

import (
    "time"
)

type Respose struct {
  Status    bool
  Data      string
  Msg       string
  Error     error
}

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Role     string `form:"role" json:"role" binding:"required"`
}

type Search struct {
	Year    int     `form:"year" json:"year"`
	Month   int     `form:"month" json:"month"`
	Day     int     `form:"day" json:"day"`
    Count   float64 `form:"count" json:"count"`
}

type SearchResponse struct {
	Year        int     `form:"year" json:"year"`
	Month       int     `form:"month" json:"month"`
	Day         int     `form:"day" json:"day"`
    Count       float64 `form:"count" json:"count"`
    Percentage  float64 `form:"percentage" json:"percentage"`
}

type SearchYearly struct {
	Year    int     `form:"year" json:"year"`
	Month   int     `form:"month" json:"month"`
    Count   float64 `form:"count" json:"count"`
}

type SearchYearlyResponse struct {
	Year        int     `form:"year" json:"year"`
	Month       int     `form:"month" json:"month"`
    Count       float64 `form:"count" json:"count"`
    Percentage  float64 `form:"percentage" json:"percentage"`
}

type AtentionResponse struct {
	Date        time.Time   `form:"date" json:"date"`
	Average     int     `form:"average" json:"average"`
}

type Minutes struct {
	Minutes     int
}

type DateTime struct {
	Date        time.Time
}

type AffiliationResponse struct {
	Nationality     string      `form:"nationality" json:"nationality"`
	Count           int         `form:"count" json:"count"`
}
