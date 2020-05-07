package db

import (
	"database/sql"
	"fmt"
	"time"

	"code.cloudfoundry.org/lager"
	sq "github.com/Masterminds/squirrel"
)

//go:generate counterfeiter . TaskQueue

// A TaskQueue handles multiple queues for the tasks in the database.
// It is used along the placement strategy: limit-active-tasks
type TaskQueue interface {
	FindOrAppend(string, string, int, string, lager.Logger) (int, int, error)
	FindQueue(string) (string, int, string, error)
	Dequeue(string, lager.Logger)
	Elapsed(string) (int, error)
	Length(string) (int, error)
	Position(string) (int, error)
}

type taskQueue struct {
	conn Conn
}

// NewTaskQueue initializes a queue with a database connection
func NewTaskQueue(conn Conn) TaskQueue {
	return &taskQueue{
		conn: conn,
	}
}

func (queue *taskQueue) FindOrAppend(id string, platform string, teamID int, workerTag string, logger lager.Logger) (int, int, error) {
	var length int
	// Returns the position and the total queue length for a given id
	exPlatform, exTeamID, exWorkerTag, err := queue.FindQueue(id)
	if err != nil && err != sql.ErrNoRows {
		return 0, 0, err
	}
	// Check that the id is not already present in a different queue, remove it from that queue in that case
	if exPlatform != platform || exTeamID != teamID || exWorkerTag != workerTag {
		logger.Info(fmt.Sprintf("%s.already-present-in-different-queue", id))
		queue.Dequeue(id, logger)
	}
	position, err := queue.Position(id)
	if err != nil {
		return 0, 0, err
	}
	if position > 0 { // Already in the queue, return position and total queue length
		length, err = queue.Length(id)
		if err != nil {
			return 0, 0, err
		}
	} else { // Append to the queue, then check its position and total queue length
		_, err := psql.Insert("tasks_queue").
			Values(id, platform, teamID, workerTag, sq.Expr("now()")).
			RunWith(queue.conn).
			Exec()
		if err != nil {
			return 0, 0, err
		}

		position, err = queue.Position(id)
		if err != nil {
			return 0, 0, err
		}
		length, err = queue.Length(id)
		if err != nil {
			return 0, 0, err
		}
	}
	return position, length, nil
}

func (queue *taskQueue) Dequeue(id string, logger lager.Logger) {
	_, err := psql.Delete("tasks_queue").
		Where(sq.Eq{"id": id}).
		RunWith(queue.conn).
		Exec()
	if err != nil {
		logger.Error("failed-to-dequeue-task", err)
	}
}

func (queue *taskQueue) Position(id string) (int, error) {
	// Return 0 if the id is not present,
	// its position if found, where 1 is the front of the queue,
	// an error in any other case.
	var position int
	platform, teamID, workerTag, err := queue.FindQueue(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	tasksPositions := psql.Select("row_number() over (order by insert_time), id").
		From("tasks_queue").
		Where(sq.Eq{"platform": platform, "team_id": teamID, "worker_tag": workerTag})
	err = psql.Select("row_number").
		FromSelect(tasksPositions, "subq").
		Where(sq.Eq{"id": id}).
		RunWith(queue.conn).
		QueryRow().
		Scan(&position)
	if err != nil {
		return 0, err
	}
	return position, nil
}

func (queue *taskQueue) FindQueue(id string) (string, int, string, error) {
	var (
		platform  string
		teamID    int
		workerTag string
	)
	err := psql.Select("platform, team_id, worker_tag").
		From("tasks_queue").
		Where(sq.Eq{"id": id}).
		RunWith(queue.conn).
		QueryRow().
		Scan(&platform, &teamID, &workerTag)
	if err != nil {
		return "", 0, "", err
	}
	return platform, teamID, workerTag, nil
}

func (queue *taskQueue) Length(id string) (int, error) {
	var length int
	platform, teamID, workerTag, err := queue.FindQueue(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, err
	}
	err = psql.Select("count(*)").
		From("tasks_queue").
		Where(sq.Eq{"platform": platform, "team_id": teamID, "worker_tag": workerTag}).
		RunWith(queue.conn).
		QueryRow().
		Scan(&length)
	if err != nil {
		return 0, err
	}
	return length, nil
}

func (queue *taskQueue) Elapsed(id string) (int, error) {
	var insertTime time.Time
	err := psql.Select("insert_time").
		From("tasks_queue").
		Where(sq.Eq{"id": id}).
		RunWith(queue.conn).
		QueryRow().
		Scan(&insertTime)
	if err != nil {
		return 0, err
	}
	duration := int(time.Since(insertTime).Seconds()) + 1
	return duration, err
}
