package model

import "time"

type Verification struct {
	ID        uint64    `db:"id"`
	Name      string    `db:"name"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
	IsRemoved bool      `db:"is_removed"`
}

type EventType string

type EventStatus string

const (
	Created EventType = "CREATED"
	Updated EventType = "UPDATED"
	Removed EventType = "REMOVED"

	Deferred  EventStatus = "DEFERRED"
	Processed EventStatus = "PROCESSED"
)

type VerificationEvent struct {
	ID             uint64
	VerificationID uint64
	Type           EventType
	Status         EventStatus
	Entity         *Verification
}

func (s *VerificationEvent) Lock(batchSize uint64) ([]VerificationEvent, error) {
	events := make([]VerificationEvent, batchSize)
	return events, nil
}

func (s *VerificationEvent) Unlock(eventsID []uint64) error {
	return nil
}

func (s *VerificationEvent) Add(event []VerificationEvent) error {
	return nil
}

func (s *VerificationEvent) Remove(eventsID []uint64) error {
	return nil
}
