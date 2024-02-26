package system

import "fmt"

func StartAll(ss ...*System) error {
	if len(ss) == 0 {
		return nil
	}

	started := make([]*System, 0, len(ss))

	for _, sys := range ss {
		if err := sys.Start(); err != nil {
			StopAll(started...)

			return fmt.Errorf("failed to start system: %w", err)
		}

		started = append(started, sys)
	}

	return nil
}
