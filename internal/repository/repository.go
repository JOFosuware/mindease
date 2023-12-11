package repository

import "github.com/jofosuware/mindease/internal/models"

type DatabaseRepo interface {
	AllUsers() bool
	InsertClient(models.Client) (int, error)
	FetchClientByEmail(string) (models.Client, error)
	InsertProvider(models.Provider) error
	FetchProvider(string) (models.Provider, error)
	FetchProviders() ([]models.Provider, error)
	InsertNotification(n models.Notification) error
	FetchNotifications() ([]models.Notification, error)
	InsertPrescription(p models.Prescription) (int, error)
	FetchPrescription(formId string) (models.Prescription, error)
}
