package store

import (
	"context"
	"fmt"
	"log"

	"github.com/USACE/instrumentation-api/api/internal/model"
	"github.com/google/uuid"
)

const (
	unknown = ""
	email   = "email"
	profile = "profile"
)

type AlertSubscriptionStore interface {
	SubscribeProfileToAlerts(ctx context.Context, alertConfigID *uuid.UUID, profileID *uuid.UUID) (*model.AlertSubscription, error)
	UnsubscribeProfileToAlerts(ctx context.Context, alertConfigID *uuid.UUID, profileID *uuid.UUID) error
	GetAlertSubscription(ctx context.Context, alertConfigID *uuid.UUID, profileID *uuid.UUID) (*model.AlertSubscription, error)
	GetAlertSubscriptionByID(ctx context.Context, id *uuid.UUID) (*model.AlertSubscription, error)
	ListMyAlertSubscriptions(ctx context.Context, profileID *uuid.UUID) ([]model.AlertSubscription, error)
	UpdateMyAlertSubscription(ctx context.Context, s *model.AlertSubscription) (*model.AlertSubscription, error)
	SubscribeEmailsToAlertConfig(ctx context.Context, alertConfigID *uuid.UUID, emails *model.EmailAutocompleteResultCollection) (*model.AlertConfig, error)
	UnsubscribeEmailsToAlertConfig(ctx context.Context, alertConfigID *uuid.UUID, emails *model.EmailAutocompleteResultCollection) (*model.AlertConfig, error)
	DeleteEmail(ctx context.Context, emailID *uuid.UUID) error
}

type alertSubscriptionStore struct {
	db *model.Database
	q  *model.Queries
}

func NewAlertSubscriptionStore(db *model.Database, q *model.Queries) *alertSubscriptionStore {
	return &alertSubscriptionStore{db, q}
}

// SubscribeProfileToAlerts subscribes a profile to an instrument alert
func (s alertSubscriptionStore) SubscribeProfileToAlerts(ctx context.Context, alertConfigID uuid.UUID, profileID uuid.UUID) (model.AlertSubscription, error) {
	var a model.AlertSubscription
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

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

// UnsubscribeProfileToAlerts subscribes a profile to an instrument alert
func (s alertSubscriptionStore) UnsubscribeProfileToAlerts(ctx context.Context, alertConfigID, profileID uuid.UUID) error {
	return s.q.UnsubscribeProfileToAlerts(ctx, alertConfigID, profileID)
}

// GetAlertSubscription returns a AlertSubscription
func (s alertSubscriptionStore) GetAlertSubscription(ctx context.Context, alertConfigID, profileID uuid.UUID) (model.AlertSubscription, error) {
	return s.q.GetAlertSubscription(ctx, alertConfigID, profileID)
}

// GetAlertSubscriptionByID returns an alert subscription
func (s alertSubscriptionStore) GetAlertSubscriptionByID(ctx context.Context, subID uuid.UUID) (model.AlertSubscription, error) {
	return s.q.GetAlertSubscriptionByID(ctx, subID)
}

// ListMyAlertSubscriptions returns all profile_alerts for a given profile ID
func (s alertSubscriptionStore) ListMyAlertSubscriptions(ctx context.Context, profileID uuid.UUID) ([]model.AlertSubscription, error) {
	return s.q.ListMyAlertSubscriptions(ctx, profileID)
}

// UpdateMyAlertSubscription updates properties on a AlertSubscription
func (s alertSubscriptionStore) UpdateMyAlertSubscription(ctx context.Context, sub model.AlertSubscription) (model.AlertSubscription, error) {
	var a model.AlertSubscription
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

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

func (s alertSubscriptionStore) SubscribeEmailsToAlertConfig(ctx context.Context, alertConfigID uuid.UUID, emails model.EmailAutocompleteResultCollection) (model.AlertConfig, error) {
	var a model.AlertConfig
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

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

func (s alertSubscriptionStore) UnsubscribeEmailsFromAlertConfig(ctx context.Context, alertConfigID uuid.UUID, emails model.EmailAutocompleteResultCollection) (model.AlertConfig, error) {
	var a model.AlertConfig
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return a, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

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

func (s alertSubscriptionStore) UnsubscribeAllFromAlertConfigTxn(ctx context.Context, alertConfigID uuid.UUID) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	qtx := s.q.WithTx(tx)

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

func (s alertSubscriptionStore) UnregisterEmail(ctx context.Context, emailID uuid.UUID) error {
	return s.q.UnregisterEmail(ctx, emailID)
}

func registerAndSubscribe(ctx context.Context, q *model.Queries, alertConfigID uuid.UUID, emails model.EmailAutocompleteResultCollection) error {
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
