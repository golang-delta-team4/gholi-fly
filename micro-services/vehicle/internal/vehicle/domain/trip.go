package domain

import "time"

type TripRequest struct {
	CompanyID      string    `json:"companyId"`
	TripType       string    `json:"tripType"`
	UserReleaseDate time.Time `json:"userReleaseDate"`
	TourReleaseDate time.Time `json:"tourReleaseDate"`
	UserPrice      float64   `json:"userPrice"`
	AgencyPrice    float64   `json:"agencyPrice"`
	PathID         string    `json:"pathId"`
	MinPassengers  int       `json:"minPassengers"`
	SoldTickets    int       `json:"soldTickets"`
	MaxTickets     int       `json:"maxTickets"`
	StartDate      time.Time `json:"startDate"`
	EndDate        time.Time `json:"endDate"`
}