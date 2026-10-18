package task

type Status int

const (
	StatusPending Status = iota
	StatusProcessing
	StatusCompleted
	StatusFailed
)

func (s Status) String() string {
	switch s {
	case StatusPending:
		return "pending"
	case StatusProcessing:
		return "processing"
	case StatusCompleted:
		return "completed"
	case StatusFailed:
		return "failed"
	}
	return "unknown"
}

func (s Status) CanChangeTo(next Status) bool {
	switch s {
	case StatusPending:
		return next == StatusProcessing
	case StatusProcessing:
		return next == StatusCompleted || next == StatusFailed
	}
	return false
}
