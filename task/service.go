package task

import "fmt"

type Service struct {
	tasks     map[string]*Task
	nextID    int
	processor Processor
}

func NewService(processor Processor) *Service {
	return &Service{
		tasks:     make(map[string]*Task),
		nextID:    1,
		processor: processor,
	}
}

func (s *Service) Create(input string) (Task, error) {
	err := ValidateInput(input)
	if err != nil {
		return Task{}, err
	}

	t := &Task{
		ID:     fmt.Sprintf("task-%d", s.nextID),
		Input:  input,
		Status: StatusPending,
	}
	s.nextID++
	s.tasks[t.ID] = t
	return t.clone(), nil
}

func (s *Service) Get(id string) (Task, error) {
	t, ok := s.tasks[id]
	if !ok {
		return Task{}, fmt.Errorf("%w: %s", ErrNotFound, id)
	}
	return t.clone(), nil
}

func (s *Service) Process(id string) (Task, error) {
	t, ok := s.tasks[id]
	if !ok {
		return Task{}, fmt.Errorf("%w: %s", ErrNotFound, id)
	}

	err := t.changeStatus(StatusProcessing)
	if err != nil {
		return Task{}, err
	}

	result, err := s.processor.Process(t.Input)
	if err != nil {
		t.Error = err.Error()
		t.Status = StatusFailed
		return t.clone(), err
	}

	t.Result = &result
	t.Status = StatusCompleted
	return t.clone(), nil
}
