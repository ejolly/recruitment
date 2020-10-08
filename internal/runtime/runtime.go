package runtime

import (
	"context"
	"time"

	"github.com/empiricaly/recruitment/internal/ent"
	"github.com/empiricaly/recruitment/internal/ent/hook"
	runModel "github.com/empiricaly/recruitment/internal/ent/run"
	"github.com/empiricaly/recruitment/internal/storage"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

// Runtime manages the empirica recruitment run loop that will trigger timed
// events as needed (start run, go to next step, etc.)
type Runtime struct {
	// db conn
	conn *storage.Conn

	// runs tracked by the runtime. The map key is the Run ID.
	runs map[string]*runState

	// update wants Run IDs to update the runtime
	// When a new Run ID comes in, it is fetched from the DB, then
	// if it is already tracked, we check what change that implies
	// for the runtime, or if it doesn't exist, we add it to the runtime.
	updates chan string

	// If done == true, stop running hooks
	done bool
}

// Start the empirica recruitment runtime
func Start(conn *storage.Conn) (*Runtime, error) {
	r := &Runtime{
		conn:    conn,
		runs:    make(map[string]*runState),
		updates: make(chan string),
	}
	go r.processRuns()
	err := r.registerExistingSteps()
	if err != nil {
		return nil, err
	}
	r.registerHooks()

	return r, nil
}

// Stop the empirica recruitment runtime
func (r *Runtime) Stop() {
	r.done = true
	// stop run loop here?
}

type runState struct {
	run   *ent.Run
	timer *time.Timer
}

func (r *Runtime) addRun(run *ent.Run) {
	_, ok := r.runs[run.ID]
	if ok {
		log.Warn().Msgf("run %s already tracked", run.ID)
		return
	}

	if run.Status == runModel.StatusCREATED {

	}
}

func (r *Runtime) removeRun(run *ent.Run) {
	state, ok := r.runs[run.ID]
	if !ok {
		return
	}

	if !state.timer.Stop() {
		<-state.timer.C
	}
}

func (r *Runtime) processRun(runID string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	runRecord, err := r.conn.Run.Get(ctx, runID)
	if err != nil {
		if ent.IsNotFound(err) {
			r.removeRun(runRecord)
		} else if ent.IsNotSingular(err) {
			log.Error().Err(err).Msg("multiple runs with same ID!")
		} else {
			log.Error().Err(err).Msg("updating Run state in runtime")
		}
		return
	}

	switch runRecord.Status {
	case runModel.StatusDONE, runModel.StatusFAILED, runModel.StatusTERMINATED, runModel.StatusPAUSED:
		r.removeRun(runRecord)
	case runModel.StatusCREATED, runModel.StatusRUNNING:
		r.addRun(runRecord)
	default:
		log.Error().Msgf("unknown run status: %s", runRecord.Status.String())
	}
}

func (r *Runtime) processRuns() {
	for {
		runID := <-r.updates
		r.processRun(runID)
	}
}

func (r *Runtime) registerHooks() {
	r.conn.Run.Use(hook.On(
		func(next ent.Mutator) ent.Mutator {
			return hook.RunFunc(func(ctx context.Context, m *ent.RunMutation) (ent.Value, error) {
				if r.done {
					return next.Mutate(ctx, m)
				}

				// After the update happened, go ahead and add Run to Runtime
				defer func() {
					ID, exists := m.ID()
					if exists {
						r.updates <- ID
					}
				}()

				return next.Mutate(ctx, m)
			})
		},
		ent.OpUpdate|ent.OpUpdateOne|ent.OpDelete|ent.OpDeleteOne))
}

func (r *Runtime) registerExistingSteps() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	runIDs, err := r.conn.Run.Query().
		Where(runModel.StatusIn(runModel.StatusCREATED, runModel.StatusRUNNING)).
		Select(runModel.FieldID).
		Strings(ctx)
	if err != nil {
		return errors.Wrap(err, "initialize exisitng runs")
	}

	log.Debug().Interface("runIDs", runIDs).Msg("got runs")

	for _, runID := range runIDs {
		r.updates <- runID
	}

	return nil
}
