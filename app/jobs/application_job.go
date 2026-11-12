// Package jobs mirrors Rails' app/jobs directory and provides the base
// contract for background jobs.
package jobs

import "context"

// Job is anything that can be performed in the background.
type Job interface {
	Perform(ctx context.Context) error
}

// ApplicationJob is embedded by concrete jobs, analogous to
// ApplicationJob in Rails.
type ApplicationJob struct{}

// QueueName lets jobs opt into a named queue; defaults to "default".
func (ApplicationJob) QueueName() string {
	return "default"
}
