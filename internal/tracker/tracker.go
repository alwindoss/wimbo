package tracker

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"go.etcd.io/bbolt"
)

const (
	TrackerBucket  = "TRACKER_BUCKET"
	ActivityBucket = "ACTIVITY_BUCKET"
	AuditLogBucket = "AUDIT_LOG_BUCKET"
)

type Activity struct {
	StartedAt time.Time
	StoppedAt time.Time
	Name      string
	Active    bool
}

type Summary struct {
	CreatedAt  time.Time
	Activities []Activity
}

type Service interface {
	StartActivity(name string) error
	StopActivity(name string) error
	SummaryizeActivity(name string)
}

func NewTrackerService(db *bbolt.DB) (Service, error) {
	err := db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(TrackerBucket))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte(ActivityBucket))
		if err != nil {
			return err
		}
		_, err = tx.CreateBucketIfNotExists([]byte(AuditLogBucket))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &trackerService{
		DB: db,
	}, nil
}

type trackerService struct {
	DB *bbolt.DB
}

// StartActivity implements TrackerService.
func (s *trackerService) StartActivity(name string) error {
	// TODO: Must implement logic to check if there is an existing activity already ongoing,
	// if no then create a new activity. If an activity exists then stop that existing activity and then
	// start the new activity

	// update the data store with the activity details
	s.DB.Update(func(tx *bbolt.Tx) error {
		id := []byte(uuid.New().String())
		act := Activity{
			StartedAt: time.Now(),
			Name:      name,
			Active:    true,
		}
		dat, err := json.Marshal(&act)
		if err != nil {
			return err
		}
		err = tx.Bucket([]byte(ActivityBucket)).Put(id, dat)
		if err != nil {
			return err
		}
		return nil
	})
	return nil
}

// StopActivity implements TrackerService.
func (*trackerService) StopActivity(name string) error {
	panic("unimplemented")
}

// SummaryizeActivity implements TrackerService.
func (*trackerService) SummaryizeActivity(name string) {
	panic("unimplemented")
}
