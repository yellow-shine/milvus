// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package concurrency

import (
	"sync"

	"github.com/panjf2000/ants/v2"
)

// A goroutine pool
type Pool struct {
	inner *ants.Pool
	opt   *poolOption
}

// Return error if provides invalid parameters
// cap: the number of workers
func NewPool(cap int, opts ...PoolOption) (*Pool, error) {
	opt := defaultPoolOption()
	for _, o := range opts {
		o(opt)
	}

	pool, err := ants.NewPool(cap, opt.antsOptions()...)
	if err != nil {
		return nil, err
	}

	return &Pool{
		inner: pool,
		opt:   opt,
	}, nil
}

// Submit a task into the pool,
// executes it asynchronously.
// This will block if the pool has finite workers and no idle worker.
func (pool *Pool) Submit(method func() (interface{}, error)) *Future {
	future := newFuture()
	err := pool.inner.Submit(func() {
		defer close(future.ch)

		if pool.opt != nil && pool.opt.preHandler != nil {
			pool.opt.preHandler()
		}

		res, err := method()
		if err != nil {
			future.err = err
		} else {
			future.value = res
		}
	})
	if err != nil {
		future.err = err
		close(future.ch)
	}

	return future
}

// The number of workers
func (pool *Pool) Cap() int {
	return pool.inner.Cap()
}

// The number of running workers
func (pool *Pool) Running() int {
	return pool.inner.Running()
}

func (pool *Pool) Release() {
	pool.inner.Release()
}

// WarmupPool do warm up logic for each goroutine in pool
func WarmupPool(pool *Pool, warmup func()) {
	cap := pool.Cap()
	ch := make(chan struct{})
	wg := sync.WaitGroup{}
	wg.Add(cap)
	for i := 0; i < cap; i++ {
		pool.Submit(func() (any, error) {
			warmup()
			wg.Done()
			<-ch
			return struct{}{}, nil
		})
	}
	wg.Wait()
	close(ch)
}
