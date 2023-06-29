package models

import (
	"time"
)

type AccountRequest struct {
	FullName       string    `json:"fullName"`
	Email          string    `json:"email"`
	Phone          string    `json:"phone"`
	InvitationCode string    `json:"invitationCode"`
	RequestStatus  string    `json:"status"`
	Reviewer       string    `json:"reviewer"`
	StatusReason   string    `json:"statusReason"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type AccountRequestStatus string

const (
	AccountRequestStatusPending  AccountRequestStatus = "pending"
	AccountRequestStatusApproved AccountRequestStatus = "approved"
	AccountRequestStatusRejected AccountRequestStatus = "rejected"
)
