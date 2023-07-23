package outbox

import "fmt"

type Status = string

const (
	StatusPrepared Status = "prepared"
	StatusDone     Status = "done"
	StatusFailed   Status = "failed"
)

func NewRetryableError(orig error) RetryableError {
	return RetryableError{orig}
}

var (
	ErrNoMoreEntries = fmt.Errorf("no more entries to process")
)

type RetryableError struct {
	err error
}

func (re RetryableError) Error() string {
	return fmt.Sprintf("retryable error: %s", re.err.Error())
}

func (re RetryableError) Unwrap() error {
	return re.err
}
