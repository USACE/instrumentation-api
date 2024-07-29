package model

import (
	"github.com/google/uuid"
)

var (
	GreenSubmittalStatusID  uuid.UUID = uuid.MustParse("0c0d6487-3f71-4121-8575-19514c7b9f03")
	YellowSubmittalStatusID uuid.UUID = uuid.MustParse("ef9a3235-f6e2-4e6c-92f6-760684308f7f")
	RedSubmittalStatusID    uuid.UUID = uuid.MustParse("84a0f437-a20a-4ac2-8a5b-f8dc35e8489b")

	MeasurementSubmittalAlertTypeID uuid.UUID = uuid.MustParse("97e7a25c-d5c7-4ded-b272-1bb6e5914fe3")
	EvaluationSubmittalAlertTypeID  uuid.UUID = uuid.MustParse("da6ee89e-58cc-4d85-8384-43c3c33a68bd")
)

const (
	warning  = "Warning"
	alert    = "Alert"
	reminder = "Reminder"
)

type AlertCheck struct {
	AlertConfigID uuid.UUID `db:"alert_config_id"`
	SubmittalID   uuid.UUID `db:"submittal_id"`
	ShouldWarn    bool      `db:"should_warn"`
	ShouldAlert   bool      `db:"should_alert"`
	ShouldRemind  bool      `db:"should_remind"`
	Submittal     Submittal `db:"-"`
}

func (ck AlertCheck) GetShouldWarn() bool {
	return ck.ShouldWarn
}

func (ck AlertCheck) GetShouldAlert() bool {
	return ck.ShouldAlert
}

func (ck AlertCheck) GetShouldRemind() bool {
	return ck.ShouldRemind
}

func (ck AlertCheck) GetSubmittal() Submittal {
	return ck.Submittal
}

func (ck *AlertCheck) SetSubmittal(sub Submittal) {
	ck.Submittal = sub
}

type AlertConfigMap map[uuid.UUID]AlertConfigScheduler

type SubmittalMap map[uuid.UUID]Submittal
