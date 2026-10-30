# Golang Application Development — KBTU, Fall 2026

```
labs/
  week01/            Go foundations
  week02/task1..2    arrays and slices
  week03/task1..5    strings and maps
  week04/task1..7    structs and methods (+ bonus)
  week05/task1..7    interfaces and generics (+ bonus)
  week06/task1..7    functions, errors, defer, panic, recover (+ bonus)
  week07/task1..7    memory management (+ bonus)
sis/                 SIS project (checkpoint 1)
  main.go            demo
  task/              tasks, validation, processing, storage, tests
```

Every task is its own program:

```bash
go run ./labs/week04/task3
go test -bench=. -benchmem ./labs/week07/task4
```

## SIS — Password Validator

A task is a unit of work: check one password.

- **Input:** the password.
- **Processing:** 7 rules — at least 8 characters, an uppercase letter, a lowercase letter, a digit, a special character, no spaces, not a common password.
- **Result:** length, score (0–7), strength (`weak`, `medium`, `strong`) and the list of failed rules.

A weak password is still valid input. Input is rejected when it is empty or whitespace-only, is not valid UTF-8, or is longer than 128 characters.

### Task lifecycle

```
pending -> processing -> completed
                      -> failed
```

- `Create` validates the input and stores a new task in `pending`.
- `Process` moves the task to `processing`, runs the validator and moves it to `completed` (result saved) or `failed` (error saved).
- `completed` and `failed` are final. Any other transition, for example processing a completed task again, returns `ErrInvalidTransition`.

### Run

```bash
go run ./sis
go run ./sis 'MyPass#2026' 'qwerty'
go test ./sis/...
```

### Checks

```bash
gofmt -l .
go vet ./...
go test ./...
```
