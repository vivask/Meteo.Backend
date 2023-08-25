package wait

import (
	"net"
	"strings"
	"sync"
	"time"
)

type Services []string

var List Services

func (s *Services) Set(value string) error {
	*s = append(*s, value)
	return nil
}

func (s *Services) String() string {
	return strings.Join(*s, ", ")
}

// Wait waits for all services
func (s *Services) Wait(tSeconds int) bool {
	t := time.Duration(tSeconds) * time.Second
	now := time.Now()

	var wg sync.WaitGroup
	wg.Add(len(*s))

	success := make(chan bool, 1)

	go func() {
		for _, service := range List {
			go waitOne(service, &wg, now)
		}
		wg.Wait()
		success <- true
	}()

	select {
	case <-success:
		return true
	case <-time.After(t):
		return false
	}
}

func waitOne(service string, wg *sync.WaitGroup, start time.Time) {
	defer wg.Done()
	for {
		_, err := net.Dial("tcp", service)
		if err == nil {
			//Log(fmt.Sprintf("%s is available after %s", service, time.Since(start)))
			break
		}
		time.Sleep(time.Second)
	}
}
