package task

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

type brokenProcessor struct{}

func (brokenProcessor) Process(input string) (Result, error) {
	return Result{}, errors.New("something broke")
}

func TestCreate(t *testing.T) {
	s := NewService(PasswordValidator{})
	got, err := s.Create("Secret123!")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID == "" {
		t.Error("task must have an ID")
	}
	if got.Input != "Secret123!" {
		t.Errorf("input = %q", got.Input)
	}
	if got.Status != StatusPending {
		t.Errorf("status = %s, want pending", got.Status)
	}
}

func TestCreateInvalidInput(t *testing.T) {
	tests := []struct {
		name  string
		input string
	}{
		{"empty", ""},
		{"spaces only", "     "},
		{"too long", strings.Repeat("a", MaxPasswordLength+1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewService(PasswordValidator{})
			_, err := s.Create(tt.input)
			if !errors.Is(err, ErrInvalidInput) {
				t.Errorf("err = %v, want ErrInvalidInput", err)
			}
		})
	}
}

func TestGet(t *testing.T) {
	s := NewService(PasswordValidator{})
	created, _ := s.Create("Secret123!")
	got, err := s.Get(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if got.ID != created.ID || got.Input != created.Input {
		t.Errorf("got %+v, want %+v", got, created)
	}
}

func TestGetNotFound(t *testing.T) {
	s := NewService(PasswordValidator{})
	_, err := s.Get("task-99")
	if !errors.Is(err, ErrNotFound) {
		t.Errorf("err = %v, want ErrNotFound", err)
	}
}

func TestProcess(t *testing.T) {
	s := NewService(PasswordValidator{})
	created, _ := s.Create("Secret123!")
	done, err := s.Process(created.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if done.Status != StatusCompleted {
		t.Errorf("status = %s, want completed", done.Status)
	}
	if done.Result == nil {
		t.Fatal("result is nil")
	}
	stored, _ := s.Get(created.ID)
	if stored.Status != StatusCompleted || stored.Result == nil {
		t.Errorf("stored task was not updated: %+v", stored)
	}
}

func TestPasswordResults(t *testing.T) {
	tests := []struct {
		name         string
		password     string
		wantScore    int
		wantStrength Strength
		wantFailed   []string
	}{
		{"strong", "Secret123!", 7, Strong, nil},
		{"only lowercase", "abcdefgh", 4, Weak, []string{"an uppercase letter", "a digit", "a special character"}},
		{"too short", "Ab1!", 6, Medium, []string{"at least 8 characters"}},
		{"common password", "Password1", 5, Weak, []string{"a special character", "not a common password"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PasswordValidator{}.Process(tt.password)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if got.Score != tt.wantScore {
				t.Errorf("score = %d, want %d", got.Score, tt.wantScore)
			}
			if got.Strength != tt.wantStrength {
				t.Errorf("strength = %s, want %s", got.Strength, tt.wantStrength)
			}
			if !reflect.DeepEqual(got.Failed, tt.wantFailed) {
				t.Errorf("failed = %v, want %v", got.Failed, tt.wantFailed)
			}
		})
	}
}

func TestProcessTwice(t *testing.T) {
	s := NewService(PasswordValidator{})
	created, _ := s.Create("Secret123!")
	_, err := s.Process(created.ID)
	if err != nil {
		t.Fatalf("first process failed: %v", err)
	}
	_, err = s.Process(created.ID)
	if !errors.Is(err, ErrInvalidTransition) {
		t.Errorf("err = %v, want ErrInvalidTransition", err)
	}
}

func TestProcessorFails(t *testing.T) {
	s := NewService(brokenProcessor{})
	created, _ := s.Create("Secret123!")
	got, err := s.Process(created.ID)
	if err == nil {
		t.Fatal("expected an error")
	}
	if got.Status != StatusFailed || got.Result != nil {
		t.Errorf("task = %+v", got)
	}
}

func TestMultipleTasks(t *testing.T) {
	s := NewService(PasswordValidator{})
	a, _ := s.Create("Secret123!")
	b, _ := s.Create("abcdefgh")
	if a.ID == b.ID {
		t.Fatalf("IDs are not unique: %s", a.ID)
	}

	s.Process(a.ID)
	s.Process(b.ID)

	ga, _ := s.Get(a.ID)
	gb, _ := s.Get(b.ID)
	if ga.Result.Strength != Strong || ga.Input != "Secret123!" {
		t.Errorf("task a = %+v", ga)
	}
	if gb.Result.Strength != Weak || gb.Input != "abcdefgh" {
		t.Errorf("task b = %+v", gb)
	}
}
