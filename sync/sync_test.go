package sync

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter(t *testing.T) {
	t.Run("increment", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assert.Equal(t, counter.Value(), 3)
	})
	t.Run("runs safely concurrently", func(t *testing.T) {
		count := 1000
		counter := Counter{}

		var wg sync.WaitGroup
		wg.Add(count)

		for i := 0; i < count; i++ {
			go func() {
				counter.Inc()
				wg.Done()
			}()
		}

		wg.Wait()

		assert.Equal(t, counter.Value(), 1000)
	})
}
