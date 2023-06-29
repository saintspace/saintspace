package datastore

import (
	"saintspace/app/models"
)

type Datastore interface {
	CreateAccountRequest(accountRequest *models.AccountRequest) error
	GetAccountRequestByInvitationCode(invitationCode string) (*models.AccountRequest, error)
	GetAccountRequestByEmail(email string) (*models.AccountRequest, error)
	GetAccountRequestsByStatus(status string) ([]*models.AccountRequest, error)
	DeleteAccountRequestByEmail(email string) error
	DeleteAccountRequestsByStatus(status string) error
}
