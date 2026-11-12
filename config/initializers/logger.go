package initializers

import "log"

// Default initializers live in this package. Each one mirrors a file you
// would drop into Rails' config/initializers directory.
func init() {
	Register("logger", func() error {
		// Consistent, UTC-based timestamps for all log output.
		log.SetFlags(log.LstdFlags | log.LUTC)
		return nil
	})
}
