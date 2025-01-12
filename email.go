package models

type Email struct {
    From    string `json:"from"`
    To      string `json:"to"`
    Subject string `json:"subject"`
    Body    string `json:"body"`
}

type EmailStats struct {
    TotalEmailsSent int `json:"total_emails_sent"`
}