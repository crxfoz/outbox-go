package worker

import "time"

const (
	defaultPollingDelay = time.Second * 10
)

type Config struct {
	// PollingDelay is a delay that Worker makes each time it pools awaiting records to be sent
	PollingDelay time.Duration
	// BatchSize is an maximum amount of records that Worker can receive and process.
	// Increasing this parameter may improve throughput as it takes less communication
	// between application and DB, on the other hand it may lead to more duplicates (in
	// the worst case the whole batch will be duplicated) because it calls Commit to
	// DB only when the whole Batch is processed.
	BatchSize  uint
	SkipLocked bool
}

func NewConfigOrdered() Config {
	return Config{
		PollingDelay: defaultPollingDelay,
		BatchSize:    1,
		SkipLocked:   false,
	}
}

func NewConfigThroughput() Config {
	return Config{
		PollingDelay: defaultPollingDelay,
		BatchSize:    25,
		SkipLocked:   true,
	}
}
