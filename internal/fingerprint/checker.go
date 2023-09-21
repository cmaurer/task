package fingerprint

import (
	"context"

	"github.com/go-task/task/v3/taskfile"
)

// StatusCheckable defines any type that can check if the status of a task is up-to-date.
type StatusCheckable interface {
	IsUpToDate(ctx context.Context, t *taskfile.Task) (bool, error)
}

// IfCheckable defines any type that can check if a command or task should not execute
type IfCheckable interface {
	ShouldSkip(ctx context.Context, t *taskfile.Task, c *taskfile.Cmd) (bool, error)
}

// SourcesCheckable defines any type that can check if the sources of a task are up-to-date.
type SourcesCheckable interface {
	IsUpToDate(t *taskfile.Task) (bool, error)
	Value(t *taskfile.Task) (any, error)
	OnError(t *taskfile.Task) error
	Kind() string
}
