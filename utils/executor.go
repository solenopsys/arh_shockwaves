package utils

import (
	"sync"
)

// todo need implement

type Executor struct {
	maxConcurrent int
	wg            sync.WaitGroup
	channels      []chan func()
}
