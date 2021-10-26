package sender

import (
	"srv-verification-api/internal/model"
)

type EventSender interface {
	Send(subdomain *model.VerificationEvent) error
}