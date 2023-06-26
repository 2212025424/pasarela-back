package subscription

import (
	"time"

	"github.com/google/uuid"

	"pasarelapago.go/model"
)

const (
	month = 1
	year  = 12
)

type UseCase struct {
	storage Storage
}

func New(s Storage) UseCase {
	return UseCase{storage: s}
}

func (uc UseCase) Create(s *model.Subscription) error {
	months := month
	if s.TypeSubs == model.Annual {
		months = year
	}

	s.ID = uuid.New()
	s.BeginsAt = time.Now()
	s.EndsAt = s.BeginsAt.AddDate(0, months, 0)
	s.Status = model.Active

	return uc.storage.Create(s)
}

func (uc UseCase) ByEmail(email string) (model.Subscriptions, error) {
	return uc.storage.ByEmail(email)
}
