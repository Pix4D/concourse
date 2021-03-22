package strategies_test

import (
	"testing"

	"github.com/concourse/concourse/integration/internal/dctest"
	"github.com/concourse/concourse/integration/internal/flytest"
	"github.com/stretchr/testify/require"
)

func TestLimitActiveTasksStrategy(t *testing.T) {
	t.Parallel()

	limitActiveTasksDC := dctest.Init(t, "../docker-compose.yml", "overrides/limit_active_tasks.yml")

	t.Run("deploy limit active tasks", func(t *testing.T) {
		limitActiveTasksDC.Run(t, "up", "-d")
	})

	fly := flytest.Init(t, limitActiveTasksDC)
	setupTestPipeline(t, fly)
	verifyLimitedTaskActivity(t, fly)
}

func setupTestPipeline(t *testing.T, fly flytest.Cmd) {
	t.Run("set up test pipeline", func(t *testing.T) {
		fly.Run(t, "set-pipeline", "-p", "test", "-c", "pipelines/limited-active-tasks-pipeline.yml", "-n")
		fly.Run(t, "unpause-pipeline", "-p", "test")

		fly.ExpectExit(1).Run(t, "trigger-job", "-j", "test/serialized-by-active-tasks-limit", "-w")
	})
}

func verifyLimitedTaskActivity(t *testing.T, fly flytest.Cmd) {
	t.Run("parallel tasks have to wait", func(t *testing.T) {
		out := fly.Output(t, "watch", "-j", "test/serialized-by-active-tasks-limit")
		require.Contains(t, out, "All workers are busy at the moment, please stand-by.")
		require.Contains(t, out, "Found a free worker after waiting")
	})

	t.Run("can still run pipeline builds", func(t *testing.T) {
		fly.ExpectExit(1).Run(t, "trigger-job", "-j", "test/serialized-by-active-tasks-limit", "-w")
	})

}
