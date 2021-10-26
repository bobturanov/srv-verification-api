package repo

import "srv-verification-api/internal/model"

type EventRepo interface {
	Lock(n uint64) ([]model.VerificationEvent, error)
	Unlock(eventIDs []uint64) error

	Add(event []model.VerificationEvent) error
	Remove(eventIDs []uint64) error
}
