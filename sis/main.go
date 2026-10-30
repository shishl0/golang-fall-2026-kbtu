package main

import (
	"errors"
	"fmt"

	"github.com/shishl0/golang-fall-2026-kbtu/sis/task"
)

func main() {
	service := task.NewService(task.PasswordValidator{})

	passwords := []string{"Secret123!", "password1", "abc", "   "}
	for _, p := range passwords {
		t, err := service.Create(p)
		if err != nil {
			fmt.Printf("create %q: %v\n", p, err)
			continue
		}
		done, err := service.Process(t.ID)
		if err != nil {
			fmt.Println("process:", err)
			continue
		}
		fmt.Printf("%s %s: %s, score %d/7, failed: %v\n", done.ID, done.Status, done.Result.Strength, done.Result.Score, done.Result.Failed)
	}

	_, err := service.Process("task-1")
	if errors.Is(err, task.ErrInvalidTransition) {
		fmt.Println("second process rejected:", err)
	}
	_, err = service.Get("task-100")
	if errors.Is(err, task.ErrNotFound) {
		fmt.Println("get rejected:", err)
	}
}
