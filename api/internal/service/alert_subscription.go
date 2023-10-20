package service

import (
	"context"
	"fmt"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

const (
	unknown = ""
	email   = "email"
	profile = "profile"
)

type AlertSubscriptionService interface {
	SubscribeProfileToAlerts(ctx context.Context, alertConfigID, profileID uuid.UUID) (model.AlertSubscription, error)
	UnsubscribeProfileToAlerts(ctx context.Context, alertConfigID, profileID uuid.UUID) error
	GetAlertSubscription(ctx context.Context, alertConfigID, profileID uuid.UUID) (model.AlertSubscription, error)
	GetAlertSubscriptionByID(ctx context.Context, subscriptionID uuid.UUID) (model.AlertSubscription, error)
	ListMyAlertSubscriptions(ctx context.Context, profileID uuid.UUID) ([]model.AlertSubscription, error)
	UpdateMyAlertSubscription(ctx context.Context, s model.AlertSubscription) (model.AlertSubscription, error)
	SubscribeEmailsToAlertConfig(ctx context.Context, alertConfigID uuid.UUID, emails []model.EmailAutocompleteResult) (model.AlertConfig, error)
	UnsubscribeEmailsFromAlertConfig(ctx context.Context, alertConfigID uuid.UUID, emails []model.EmailAutocompleteResult) (model.AlertConfig, error)
	UnsubscribeAllFromAlertConfig(ctx context.Context, alertConfigID uuid.UUID) error
	UnregisterEmail(ctx context.Context, emailID uuid.UUID) error
}

type alertSubscriptionService struct {
	db *model.Database
	*model.Queries
}

func NewAlertSubscriptionService(db *model.Database, q *model.Queries) *alertSubscriptionService {
	return &alertSubscriptionService{db, q}
}

// SubscribeProfileToAlerts subscribes a profile to an instrument alert
func (s alertSubscriptionService) SubscribeProfileToAlerts(ctx context.Context, alertConfigID uuid.UUID, profileID uuid.UUID) (model.AlertSubscription, error) {
	var a model.AlertSubscription
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.SubscribeProfileToAlerts(ctx, alertConfigID, profileID); err != nil {
		return a, err
	}

	updated, err := qtx.GetAlertSubscription(ctx, alertConfigID, profileID)
	if err != nil {
		return a, err
	}

	if err := tx.Commit(); err != nil {
		return a, err
	}

	return updated, nil
}

// UpdateMyAlertSubscription updates properties on a AlertSubscription
func (s alertSubscriptionService) UpdateMyAlertSubscription(ctx context.Context, sub model.AlertSubscription) (model.AlertSubscription, error) {
	var a model.AlertSubscription
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdateMyAlertSubscription(ctx, sub); err != nil {
		return a, err
	}

	updated, err := qtx.GetAlertSubscription(ctx, sub.AlertConfigID, sub.ProfileID)
	if err != nil {
		return a, err
	}

	if err := tx.Commit(); err != nil {
		return a, err
	}

	return updated, nil
}

func (s alertSubscriptionService) SubscribeEmailsToAlertConfig(ctx context.Context, alertConfigID uuid.UUID, emails []model.EmailAutocompleteResult) (model.AlertConfig, error) {
	var a model.AlertConfig
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := registerAndSubscribe(ctx, qtx, alertConfigID, emails); err != nil {
		return a, err
	}

	// Register any emails that are not yet in system
	for idx, em := range emails {
		if em.UserType == unknown || em.UserType == email {
			newID, err := qtx.RegisterEmail(ctx, em.Email)
			if err != nil {
				return a, err
			}
			emails[idx].ID = newID
			emails[idx].UserType = email
		}
	}
	// Subscribe emails
	for _, em := range emails {
		if em.UserType == email {
			if err := qtx.SubscribeEmailToAlertConfig(ctx, alertConfigID, em.ID); err != nil {
				return a, err
			}
		} else if em.UserType == profile {
			if err := qtx.SubscribeProfileToAlertConfig(ctx, alertConfigID, em.ID); err != nil {
				return a, err
			}
		} else {
			return a, fmt.Errorf("unable to unsubscribe email %s: user type %s does not exist, aborting transaction", em.Email, em.UserType)
		}
	}

	acUpdated, err := qtx.GetOneAlertConfig(ctx, alertConfigID)
	if err != nil {
		return a, err
	}

	if err := tx.Commit(); err != nil {
		return a, err
	}

	return acUpdated, nil
}

func (s alertSubscriptionService) UnsubscribeEmailsFromAlertConfig(ctx context.Context, alertConfigID uuid.UUID, emails []model.EmailAutocompleteResult) (model.AlertConfig, error) {
	var a model.AlertConfig
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	for _, em := range emails {
		if em.UserType == unknown {
			return a, fmt.Errorf("required field user_type is null, aborting transaction")
		} else if em.UserType == email {
			if err := qtx.UnsubscribeEmailFromAlertConfig(ctx, alertConfigID, em.ID); err != nil {
				return a, err
			}
		} else if em.UserType == profile {
			if err := qtx.UnsubscribeProfileFromAlertConfig(ctx, alertConfigID, em.ID); err != nil {
				return a, err
			}
		} else {
			return a, fmt.Errorf("unable to unsubscribe email %s: user type %s does not exist, aborting transaction", em.Email, em.UserType)
		}
	}

	acUpdated, err := qtx.GetOneAlertConfig(ctx, alertConfigID)
	if err != nil {
		return a, err
	}

	if err := tx.Commit(); err != nil {
		return a, err
	}

	return acUpdated, nil
}

func (s alertSubscriptionService) UnsubscribeAllFromAlertConfig(ctx context.Context, alertConfigID uuid.UUID) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UnsubscribeAllEmailsFromAlertConfig(ctx, alertConfigID); err != nil {
		return err
	}

	if err := qtx.UnsubscribeAllProfilesFromAlertConfig(ctx, alertConfigID); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func registerAndSubscribe(ctx context.Context, q *model.Queries, alertConfigID uuid.UUID, emails []model.EmailAutocompleteResult) error {
	for idx, em := range emails {
		if em.UserType == unknown || em.UserType == email {
			newID, err := q.RegisterEmail(ctx, em.Email)
			if err != nil {
				return err
			}
			emails[idx].ID = newID
			emails[idx].UserType = email
		}
	}
	for _, em := range emails {
		if em.UserType == email {
			if err := q.SubscribeEmailToAlertConfig(ctx, alertConfigID, em.ID); err != nil {
				return err
			}
		} else if em.UserType == profile {
			if err := q.SubscribeProfileToAlertConfig(ctx, alertConfigID, em.ID); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unable to unsubscribe email %s: user type %s does not exist, aborting transaction", em.Email, em.UserType)
		}
	}
	return nil
}
