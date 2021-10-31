package sender

import (
	"github.com/ozonmp/srv-verification-api/internal/model"
)

type EventSender interface {
	Send(subdomain *model.VerificationEvent) error
}