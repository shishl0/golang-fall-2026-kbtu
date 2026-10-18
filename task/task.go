package task

import "fmt"

type Strength string

const (
	Weak   Strength = "weak"
	Medium Strength = "medium"
	Strong Strength = "strong"
)

type Result struct {
	Length   int
	Score    int
	Strength Strength
	Failed   []string
}

type Task struct {
	ID     string
	Input  string
	Status Status
	Result *Result
	Error  string
}

func (t *Task) changeStatus(next Status) error {
	if !t.Status.CanChangeTo(next) {
		return fmt.Errorf("%w: %s -> %s", ErrInvalidTransition, t.Status, next)
	}
	t.Status = next
	return nil
}

func (t Task) clone() Task {
	if t.Result != nil {
		r := *t.Result
		r.Failed = append([]string(nil), t.Result.Failed...)
		t.Result = &r
	}
	return t
}
