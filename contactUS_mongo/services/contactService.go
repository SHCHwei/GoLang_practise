package services

import (
    "contentAPI/models"
    "time"
)

type ContactService interface {
    CreateContact(*models.ContactOne) error
    SearchContact(time.Time, time.Time) ([]models.Contacts, error)
}
