package model

type Verification struct {
	ID uint64
    Name string
}

type EventType uint8

type EventStatus uint8

const (
	Created EventType = iota
	Updated
	Removed

	Deferred EventStatus = iota
	Processed
)

type VerificationEvent struct {
	ID     uint64
	Type   EventType
	Status EventStatus
	Entity *Verification
}

func (s * VerificationEvent) Lock(batchSize uint64) ([]VerificationEvent, error){
	events := make([]VerificationEvent, batchSize)
	return events, nil
}

func (s * VerificationEvent) Unlock(eventsID []uint64) error{
	return nil
}

func (s * VerificationEvent) Add(event []VerificationEvent) error{
	return nil
}

func (s * VerificationEvent) Remove(eventsID []uint64) error{
	return nil
}