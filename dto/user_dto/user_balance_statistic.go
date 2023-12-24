package user_dto

import "time"

type GetStatResp struct {
	TotalUsage            float64
	TotalUsageCurrency    string
	LastWeekUsage         float64
	LastWeekUsageCurrency string
}

type GetTotalUsageResp struct {
	TotalUsage         float64
	TotalUsageCurrency string
}

type GetLastWeekUsageResp struct {
	LastWeekUsage         float64
	LastWeekUsageCurrency string
}

type GetTrendResp struct {
	Data map[string]float64
}

type UserBalanceStatistic struct {
	ID         string    `json:"id"`
	UserID     string    `json:"user_id"`
	Balance    float64   `json:"balance"`
	Currency   string    `json:"currency"`
	Date       string    `json:"date"`
	CreateTime time.Time `json:"create_time"`
	ModifyTime time.Time `json:"modify_time"`
}
