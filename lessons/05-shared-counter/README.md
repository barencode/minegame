1. Detected data-race by two goroutines.
2. Mutex applied to counter protected it from concurrent read and write access. All the actions (2 * 100_000 conts make up to 200_000)