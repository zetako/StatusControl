package scontrol

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestBasicLogic(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	sc := New()

	go func() {
		for sc.Check() != StatusStop {
			t.Log("Routine 1: Working...")
			time.Sleep(time.Microsecond * time.Duration(rand.Intn(200)))
		}
		wg.Done()
	}()

	go func() {
		for sc.Check() != StatusStop {
			t.Log("Routine 2: Working...")
			time.Sleep(time.Microsecond * time.Duration(rand.Intn(100)))
		}
		wg.Done()
	}()

	time.Sleep(time.Second)
	sc.Set(StatusPause)
	t.Log("pause!")
	time.Sleep(time.Second)
	t.Log("running!")
	sc.Set(StatusRunning)
	time.Sleep(time.Second)
	sc.Set(StatusStop)

	wg.Wait()
}
