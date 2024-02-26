package system

import "sync"

func StopAll(ss ...*System) {
	if len(ss) == 0 {
		return
	}

	wg := new(sync.WaitGroup)
	n := len(ss)

	wg.Add(n)

	for i := 0; i < n; i++ {
		sys := ss[i]

		go func() {
			sys.Stop()

			wg.Done()
		}()
	}

	wg.Wait()
}
