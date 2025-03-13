package concurrency

import (
	"errors"
	"fmt"
	"net/http"
	"time"
)

type PressureGauge struct {
	ch chan struct{}
}

func NewPressureGauge(limit int) *PressureGauge {
	return &PressureGauge{make(chan struct{}, limit)}
}
func (pg *PressureGauge) Process(f func()) error {
	select {
	case pg.ch <- struct{}{}:
		f()
		<-pg.ch
		return nil
	default:
		return errors.New("no more capacity")
	}
}

func executeExpensiveOperation() string {
	time.Sleep(2 * time.Second)
	return "done"
}

func RunServer() {
	pg := NewPressureGauge(3)
	http.HandleFunc("/request", func(writer http.ResponseWriter, request *http.Request) {
		err := pg.Process(func() {
			status, err := writer.Write([]byte(executeExpensiveOperation()))
			logResponse("Executed:", status, err)
		})

		if err != nil {
			writer.WriteHeader(http.StatusTooManyRequests)
			status, tooManyRequestsWritingError := writer.Write([]byte("Too many requests"))
			logResponse("Validated:", status, tooManyRequestsWritingError)
		}
	})
	_ = http.ListenAndServe(":8080", nil)
}

func logResponse(prefix string, status int, err error) {
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Println(prefix, status)
	}
}
