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
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.AlertSubscription{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.SubscribeProfileToAlerts(ctx, alertConfigID, profileID); err != nil {
		return model.AlertSubscription{}, err
	}

	updated, err := qtx.GetAlertSubscription(ctx, alertConfigID, profileID)
	if err != nil {
		return model.AlertSubscription{}, err
	}

	err = tx.Commit()

	return updated, err
}

// UpdateMyAlertSubscription updates properties on a AlertSubscription
func (s alertSubscriptionService) UpdateMyAlertSubscription(ctx context.Context, sub model.AlertSubscription) (model.AlertSubscription, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return sub, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := qtx.UpdateMyAlertSubscription(ctx, sub); err != nil {
		return sub, err
	}

	updated, err := qtx.GetAlertSubscription(ctx, sub.AlertConfigID, sub.ProfileID)
	if err != nil {
		return sub, err
	}

	err = tx.Commit()

	return updated, err
}

func (s alertSubscriptionService) SubscribeEmailsToAlertConfig(ctx context.Context, alertConfigID uuid.UUID, emails []model.EmailAutocompleteResult) (model.AlertConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.AlertConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	if err := registerAndSubscribe(ctx, qtx, alertConfigID, emails); err != nil {
		return model.AlertConfig{}, err
	}

	// Register any emails that are not yet in system
	for idx, em := range emails {
		if em.UserType == unknown || em.UserType == email {
			newID, err := qtx.RegisterEmail(ctx, em.Email)
			if err != nil {
				return model.AlertConfig{}, err
			}
			emails[idx].ID = newID
			emails[idx].UserType = email
		}
	}
	// Subscribe emails
	for _, em := range emails {
		if em.UserType == email {
			if err := qtx.SubscribeEmailToAlertConfig(ctx, alertConfigID, em.ID); err != nil {
				return model.AlertConfig{}, err
			}
		} else if em.UserType == profile {
			if err := qtx.SubscribeProfileToAlertConfig(ctx, alertConfigID, em.ID); err != nil {
				return model.AlertConfig{}, err
			}
		} else {
			return model.AlertConfig{}, fmt.Errorf("unable to unsubscribe email %s: user type %s does not exist, aborting transaction", em.Email, em.UserType)
		}
	}

	acUpdated, err := qtx.GetOneAlertConfig(ctx, alertConfigID)
	if err != nil {
		return model.AlertConfig{}, err
	}

	err = tx.Commit()

	return acUpdated, err
}

func (s alertSubscriptionService) UnsubscribeEmailsFromAlertConfig(ctx context.Context, alertConfigID uuid.UUID, emails []model.EmailAutocompleteResult) (model.AlertConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return model.AlertConfig{}, err
	}
	defer model.TxDo(tx.Rollback)

	qtx := s.WithTx(tx)

	for _, em := range emails {
		if em.UserType == unknown {
			return model.AlertConfig{}, fmt.Errorf("required field user_type is null, aborting transaction")
		} else if em.UserType == email {
			if err := qtx.UnsubscribeEmailFromAlertConfig(ctx, alertConfigID, em.ID); err != nil {
				return model.AlertConfig{}, err
			}
		} else if em.UserType == profile {
			if err := qtx.UnsubscribeProfileFromAlertConfig(ctx, alertConfigID, em.ID); err != nil {
				return model.AlertConfig{}, err
			}
		} else {
			return model.AlertConfig{}, fmt.Errorf("unable to unsubscribe email %s: user type %s does not exist, aborting transaction", em.Email, em.UserType)
		}
	}

	acUpdated, err := qtx.GetOneAlertConfig(ctx, alertConfigID)
	if err != nil {
		return model.AlertConfig{}, err
	}

	err = tx.Commit()

	return acUpdated, err
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

	return tx.Commit()
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
