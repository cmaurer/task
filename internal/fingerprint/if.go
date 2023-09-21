package fingerprint

import (
	"context"

	"github.com/go-task/task/v3/internal/env"
	"github.com/go-task/task/v3/internal/execext"
	"github.com/go-task/task/v3/internal/logger"
	"github.com/go-task/task/v3/taskfile"
)

type IfChecker struct {
	logger *logger.Logger
}

func NewIfChecker(logger *logger.Logger) IfCheckable {
	return &IfChecker{
		logger: logger,
	}
}

func (checker *IfChecker) ShouldSkip(ctx context.Context, t *taskfile.Task, c *taskfile.Cmd) (bool, error) {
	var command string
	if len(c.If.Static) > 0 {
		command = c.If.Static
	} else {
		command = c.If.Sh
	}
	err := execext.RunCommand(ctx, &execext.RunCommandOptions{
		Command: command,
		Dir:     t.Dir,
		Env:     env.Get(t),
	})
	if err != nil {
		return false, nil
	}
	return true, nil
}
