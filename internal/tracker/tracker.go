package tracker

import "time"

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

type TrackerService interface {
	StartActivity(name string) error
	StopActivity(name string) error
	SummaryizeActivity(name string)
}

func NewTracker() TrackerService {
	return &trackerService{}
}

type trackerService struct{}

// StartActivity implements TrackerService.
func (*trackerService) StartActivity(name string) error {
	panic("unimplemented")
}

// StopActivity implements TrackerService.
func (*trackerService) StopActivity(name string) error {
	panic("unimplemented")
}

// SummaryizeActivity implements TrackerService.
func (*trackerService) SummaryizeActivity(name string) {
	panic("unimplemented")
}
