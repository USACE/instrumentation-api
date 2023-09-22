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
}

func NewAlertSubscriptionStore(db *model.Database) *alertSubscriptionStore {
	return &alertSubscriptionStore{db}
}

// SubscribeProfileToAlerts subscribes a profile to an instrument alert
func (s alertSubscriptionStore) SubscribeProfileToAlerts(ctx context.Context, alertConfigID *uuid.UUID, profileID *uuid.UUID) (*model.AlertSubscription, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	q := model.NewQueries(s.db).WithTx(tx)

	if err := q.SubscribeProfileToAlerts(ctx, alertConfigID, profileID); err != nil {
		return nil, err
	}

	updated, err := q.GetAlertSubscription(ctx, alertConfigID, profileID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return updated, nil
}

// UnsubscribeProfileToAlerts subscribes a profile to an instrument alert
func (s alertSubscriptionStore) UnsubscribeProfileToAlerts(ctx context.Context, alertConfigID *uuid.UUID, profileID *uuid.UUID) error {
	q := model.NewQueries(s.db)

	if err := q.UnsubscribeProfileToAlerts(ctx, alertConfigID, profileID); err != nil {
		return err
	}

	return nil
}

// GetAlertSubscription returns a AlertSubscription
func (s alertSubscriptionStore) GetAlertSubscription(ctx context.Context, alertConfigID *uuid.UUID, profileID *uuid.UUID) (*model.AlertSubscription, error) {
	q := model.NewQueries(s.db)

	pa, err := q.GetAlertSubscription(ctx, alertConfigID, profileID)
	if err != nil {
		return nil, err
	}
	return pa, nil
}

// GetAlertSubscriptionByID returns an alert subscription
func (s alertSubscriptionStore) GetAlertSubscriptionByID(ctx context.Context, subID *uuid.UUID) (*model.AlertSubscription, error) {
	q := model.NewQueries(s.db)

	aa, err := q.GetAlertSubscriptionByID(ctx, subID)
	if err != nil {
		return nil, err
	}
	return aa, nil
}

// ListMyAlertSubscriptions returns all profile_alerts for a given profile ID
func (s alertSubscriptionStore) ListMyAlertSubscriptions(ctx context.Context, profileID *uuid.UUID) ([]model.AlertSubscription, error) {
	q := model.NewQueries(s.db)

	aa, err := q.ListMyAlertSubscriptions(ctx, profileID)
	if err != nil {
		return nil, err
	}
	return aa, nil
}

// UpdateMyAlertSubscription updates properties on a AlertSubscription
func (s alertSubscriptionStore) UpdateMyAlertSubscription(ctx context.Context, sub *model.AlertSubscription) (*model.AlertSubscription, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	q := model.NewQueries(s.db).WithTx(tx)

	if err := q.UpdateMyAlertSubscription(ctx, sub); err != nil {
		return nil, err
	}

	updated, err := q.GetAlertSubscription(ctx, &sub.AlertConfigID, &sub.ProfileID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return updated, nil
}

func (s alertSubscriptionStore) SubscribeEmailsToAlertConfig(ctx context.Context, alertConfigID *uuid.UUID, emails model.EmailAutocompleteResultCollection) (*model.AlertConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	q := model.NewQueries(s.db).WithTx(tx)

	if err := registerAndSubscribe(ctx, q, alertConfigID, emails); err != nil {
		return nil, err
	}

	// Register any emails that are not yet in system
	for idx, em := range emails {
		if em.UserType == unknown || em.UserType == email {
			newID, err := q.RegisterEmail(ctx, em.Email)
			if err != nil {
				return nil, err
			}
			if newID == nil {
				continue
			}
			emails[idx].ID = *newID
			emails[idx].UserType = email
		}
	}
	// Subscribe emails
	for _, em := range emails {
		if em.UserType == email {
			if err := q.SubscribeEmailToAlertConfig(ctx, alertConfigID, &em.ID); err != nil {
				return nil, err
			}
		} else if em.UserType == profile {
			if err := q.SubscribeProfileToAlertConfig(ctx, alertConfigID, &em.ID); err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("unable to unsubscribe email %s: user type %s does not exist, aborting transaction", em.Email, em.UserType)
		}
	}

	acUpdated, err := q.GetOneAlertConfig(ctx, alertConfigID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return acUpdated, nil
}

func (s alertSubscriptionStore) UnsubscribeEmailsFromAlertConfig(ctx context.Context, alertConfigID *uuid.UUID, emails model.EmailAutocompleteResultCollection) (*model.AlertConfig, error) {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	q := model.NewQueries(s.db).WithTx(tx)

	for _, em := range emails {
		if em.UserType == unknown {
			return nil, fmt.Errorf("required field user_type is null, aborting transaction")
		} else if em.UserType == email {
			if err := q.UnsubscribeEmailFromAlertConfig(ctx, alertConfigID, &em.ID); err != nil {
				return nil, err
			}
		} else if em.UserType == profile {
			if err := q.UnsubscribeProfileFromAlertConfig(ctx, alertConfigID, &em.ID); err != nil {
				return nil, err
			}
		} else {
			return nil, fmt.Errorf("unable to unsubscribe email %s: user type %s does not exist, aborting transaction", em.Email, em.UserType)
		}
	}

	acUpdated, err := q.GetOneAlertConfig(ctx, alertConfigID)
	if err != nil {
		return nil, err
	}

	if err := tx.Commit(); err != nil {
		return nil, err
	}

	return acUpdated, nil
}

func (s alertSubscriptionStore) UnsubscribeAllFromAlertConfigTxn(ctx context.Context, alertConfigID *uuid.UUID) error {
	tx, err := s.db.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		if err := tx.Rollback(); err != nil {
			log.Print(err.Error())
		}
	}()

	q := model.NewQueries(s.db).WithTx(tx)

	if err := q.UnsubscribeAllEmailsFromAlertConfig(ctx, alertConfigID); err != nil {
		return err
	}

	if err := q.UnsubscribeAllProfilesFromAlertConfig(ctx, alertConfigID); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}

func (s alertSubscriptionStore) UnregisterEmail(ctx context.Context, emailID *uuid.UUID) error {
	q := model.NewQueries(s.db)
	if err := q.UnregisterEmail(ctx, emailID); err != nil {
		return err
	}

	return nil
}

type registerFn func(ctx context.Context, email string) (*uuid.UUID, error)
type subFn func(ctx context.Context, alertConfigID, emailID *uuid.UUID) error

func registerAndSubscribe(ctx context.Context, q *model.Queries, alertConfigID *uuid.UUID, emails model.EmailAutocompleteResultCollection) error {
	// Register any emails that are not yet in system
	for idx, em := range emails {
		if em.UserType == unknown || em.UserType == email {
			newID, err := q.RegisterEmail(ctx, em.Email)
			if err != nil {
				return err
			}
			if newID == nil {
				continue
			}
			emails[idx].ID = *newID
			emails[idx].UserType = email
		}
	}
	// Subscribe emails
	for _, em := range emails {
		if em.UserType == email {
			if err := q.SubscribeEmailToAlertConfig(ctx, alertConfigID, &em.ID); err != nil {
				return err
			}
		} else if em.UserType == profile {
			if err := q.SubscribeProfileToAlertConfig(ctx, alertConfigID, &em.ID); err != nil {
				return err
			}
		} else {
			return fmt.Errorf("unable to unsubscribe email %s: user type %s does not exist, aborting transaction", em.Email, em.UserType)
		}
	}
	return nil
}
