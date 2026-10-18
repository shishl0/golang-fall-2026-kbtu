package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
)

type Notifier interface {
	Notify(msg string) error
}

type Email struct {
	Address string
}

type SMS struct {
	Phone string
}

type Telegram struct {
	Chat string
	Sent int
}

func (e Email) Notify(msg string) error {
	fmt.Printf("[Email -> %s] %s\n", e.Address, msg)
	return nil
}

func (s SMS) Notify(msg string) error {
	if len([]rune(msg)) > 20 {
		return errors.New("sms to " + s.Phone + ": message is longer than 20 characters")
	}
	fmt.Printf("[SMS -> %s] %s\n", s.Phone, msg)
	return nil
}

func (t *Telegram) Notify(msg string) error {
	t.Sent++
	fmt.Printf("[Telegram -> %s] %s (sent: %d)\n", t.Chat, msg, t.Sent)
	return nil
}

var _ Notifier = Email{}
var _ Notifier = SMS{}
var _ Notifier = (*Telegram)(nil)

func Broadcast(ns []Notifier, msg string) []error {
	var errs []error
	for _, n := range ns {
		err := n.Notify(msg)
		if err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func printGrades(w io.Writer, grades map[string]int) {
	fmt.Fprintf(w, "%-10s | %s\n", "Name", "Grade")
	fmt.Fprintln(w, "-----------+------")
	for name, grade := range grades {
		fmt.Fprintf(w, "%-10s | %d\n", name, grade)
	}
}

func task1() {
	tg := &Telegram{Chat: "@group"}
	notifiers := []Notifier{
		Email{Address: "student@kbtu.kz"},
		SMS{Phone: "+77001234567"},
		tg,
	}

	errs := Broadcast(notifiers, "Hi!")
	fmt.Println("errors:", errs)

	errs = Broadcast(notifiers, "This message is definitely too long for SMS")
	fmt.Println("errors:", errs)
	fmt.Println("telegram sent total:", tg.Sent)

	fmt.Println("\n--- grades to os.Stdout ---")
	grades := map[string]int{"Aruzhan": 95, "Dias": 80, "Zhan": 88}
	printGrades(os.Stdout, grades)

	fmt.Println("\n--- grades to strings.Builder ---")
	var sb strings.Builder
	printGrades(&sb, grades)
	fmt.Print(sb.String())
}
