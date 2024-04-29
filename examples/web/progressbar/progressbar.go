package progressbar

import (
	"fmt"
	"net/http"
	"strconv"
	"sync/atomic"
	"time"

	htmx "github.com/angelofallars/htmx-go"

	"github.com/will-wow/typed-htmx-go/examples/web/progressbar/exgom"
	"github.com/will-wow/typed-htmx-go/examples/web/progressbar/extempl"
	"github.com/will-wow/typed-htmx-go/examples/web/progressbar/shared"
)

type example struct {
	gom  bool
	jobs *jobs
}

func NewHandler(gom bool) http.Handler {
	mux := http.NewServeMux()

	ex := example{
		gom:  gom,
		jobs: newJobs(),
	}

	mux.HandleFunc("GET /{$}", ex.demo)
	mux.HandleFunc("POST /job/{$}", ex.start)
	mux.HandleFunc("GET /job/{id}/progress/{$}", ex.progress)
	mux.HandleFunc("GET /job/{id}/{$}", ex.job)

	return mux
}

func (ex *example) demo(w http.ResponseWriter, r *http.Request) {
	if ex.gom {
		_ = exgom.Page().Render(w)
	} else {
		_ = extempl.Page().Render(r.Context(), w)
	}
}

func (ex *example) start(w http.ResponseWriter, r *http.Request) {
	id := ex.jobs.add()

	if ex.gom {
		_ = exgom.JobRunning(id, 0).Render(w)
	} else {
		_ = extempl.JobRunning(id, 0).Render(r.Context(), w)
	}
}

func (ex *example) progress(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	if idParam == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	progress, err := ex.jobs.getProgress(id)
	if err != nil {
		http.Error(w, "job not found", http.StatusNotFound)
		return
	}

	res := htmx.NewResponse()

	if progress >= 100 {
		res = res.AddTrigger(htmx.Trigger(shared.TriggerDone))
	}

	if ex.gom {
		res.MustWrite(w)
		_ = exgom.ProgressBar(progress).Render(w)
	} else {
		res.MustRenderTempl(r.Context(), w, extempl.ProgressBar(progress))
	}
}

func (ex *example) job(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	if idParam == "" {
		http.Error(w, "missing id", http.StatusBadRequest)
		return
	}

	id, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	progress, err := ex.jobs.getProgress(id)
	if err != nil {
		http.Error(w, "job not found", http.StatusNotFound)
		return
	}

	if ex.gom {
		_ = exgom.Job(id, progress).Render(w)
	} else {
		_ = extempl.Job(id, progress).Render(r.Context(), w)
	}
}

type job struct {
	duration  time.Duration
	startTime time.Time
}

type jobs struct {
	nextID atomic.Int64
	active map[int64]job
}

func newJobs() *jobs {
	return &jobs{
		nextID: atomic.Int64{},
		active: map[int64]job{},
	}
}

func (j *jobs) getNextID() int64 {
	return j.nextID.Add(1)
}

func (j *jobs) add() int64 {
	id := j.getNextID()

	startTime := time.Now()
	duration := time.Second * 6
	job := job{
		startTime: startTime,
		duration:  duration,
	}

	j.active[id] = job
	return id
}

func (j *jobs) getProgress(id int64) (int, error) {
	job, ok := j.active[id]
	if !ok {
		return 0, fmt.Errorf("job not found")
	}

	timeElapsed := time.Since(job.startTime)

	if timeElapsed >= job.duration {
		return 100, nil
	}

	return int((float64(timeElapsed) / float64(job.duration)) * 100), nil
}
