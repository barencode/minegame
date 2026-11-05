1. if worker() gets parameters, we should use wrapper for wg.Go: wg.Go( func() { worker(...) } )
2. if we call close() before getter and not within goroutine, we'll get deadlock