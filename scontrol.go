package scontrol

import "sync"

// Status is possible status of controller
type Status int

const (
	StatusRunning Status = 0 // StatusRunning is normal status
	StatusPause   Status = 1 // StatusPause will block any Check() until it change
	StatusStop    Status = 2 // StatusStop is stop status, routine should exit after received this
)

// Controller provide cross routine control.
// DO NOT copy it!
type Controller struct {
	s      Status          // status now
	locker *sync.RWMutex   // protect s
	wg     *sync.WaitGroup // act as Semaphore
}

// New return a new Controller
func New() *Controller {
	return &Controller{
		s:      StatusRunning,
		wg:     &sync.WaitGroup{},
		locker: &sync.RWMutex{},
	}
}

// Set controller to a new status
func (c *Controller) Set(s Status) {
	c.locker.Lock()
	defer c.locker.Unlock()

	// Set Value
	old := c.s
	c.s = s

	// Set WaitGroup/Semaphore
	if old != StatusPause && s == StatusPause {
		c.wg.Add(1)
	} else if old == StatusPause && s != StatusPause {
		c.wg.Done()
	}
}

// Check returns status now; if status is StatusPause, it will block until it change.
func (c *Controller) Check() Status {
	c.wg.Wait() // if Status is not Pause, it will return immediately
	c.locker.RLock()
	defer c.locker.RUnlock()
	tmp := c.s
	return tmp
}

// Get return status immediately; no blocking.
func (c *Controller) Get() Status {
	c.locker.RLock()
	defer c.locker.RUnlock()
	tmp := c.s
	return tmp
}
