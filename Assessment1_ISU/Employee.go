package main

type Employee struct {
	EmpID   int64   `json:"id"`
	EmpName string  `json:"empname"`
	Balance float64 `json:"balance"`
}
