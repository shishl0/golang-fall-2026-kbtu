# Golang Application Development Kazakh-British Technical University Fall 2026 

## Lectures

1. [Lecture 01 — Go Foundations](./lectures/01-go-foundations/README.md)
2. [Lecture 02 — Arrays and Slices](./lectures/02-arrays-slices/README.md)
3. [Lecture 03 — Strings and Maps](./lectures/03-strings-maps/README.md)
4. [Lecture 04 — Structs and Methods](./lectures/04-structs-methods/README.md)
5. [Lecture 05 — Interfaces and Generics](./lectures/05-interfaces-generics/README.md)
6. [Lecture 06 — Functions, Errors, Defer, Panic and Recover](./lectures/06-functions-errors/README.md)
7. [Lecture 07 — Memory Management](./lectures/07-memory-management/README.md)

## Checkpoints (SIS)

1. [Checkpoint 01 — Go Foundations](sis/SIS_CP1_SPECIFICATION.md)

### SIS 1 — Password Validator

A task is a unit of work: check a password. Input: the password. Processing: check 7 rules (length of at least 8, uppercase letter, lowercase letter, digit, special character, no spaces, not a common password). Result: length, score (0 to 7), strength (weak, medium, strong) and the list of failed rules.

A weak password is still valid input. Input is rejected only when it is empty or whitespace-only, is not valid UTF-8, or is longer than 128 characters.

Task lifecycle:

```
pending -> processing -> completed
                      -> failed
```

- `Create` validates the input and stores a new task in `pending`.
- `Process` moves the task to `processing`, runs the validator and moves it to `completed` (result saved) or `failed` (error saved).
- `completed` and `failed` are final. Any other transition, for example processing a completed task again, returns `ErrInvalidTransition`.

```bash
go run ./cmd/sis
gofmt -l .
go vet ./...
go test ./...
```
