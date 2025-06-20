package contextsamples

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

func sendRequest(ctx context.Context, url string) (*http.Response, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	return http.DefaultClient.Do(req)
}

func ContextWithCancelAndCause() {
	ctx, cancelFuncParent := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancelFuncParent()
	ctx, cancelFunc := context.WithCancelCause(ctx)
	defer cancelFunc(nil)
	ch := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			resp, err := sendRequest(ctx, "http://httpbin.org/status/200,200,200,500")
			if err != nil {
				cancelFunc(fmt.Errorf("in status goroutine: %w", err))
				return
			}
			if resp.StatusCode == http.StatusInternalServerError {
				cancelFunc(errors.New("bad status"))
				return
			}
			ch <- "success from status"
			time.Sleep(1 * time.Second)
		}
	}()
	go func() {
		defer wg.Done()
		for {
			resp, err := sendRequest(ctx, "http://httpbin.org/delay/1")
			if err != nil {
				fmt.Println("in delay goroutine:", err)
				cancelFunc(fmt.Errorf("in delay goroutine: %w", err))
				return
			}
			ch <- "success from delay: " + resp.Header.Get("date")
		}
	}()
loop:
	for {
		select {
		case s := <-ch:
			fmt.Println("in main:", s)
		case <-ctx.Done():
			fmt.Println("in main: cancelled with cause:", context.Cause(ctx), "err:", ctx.Err())
			break loop
		}
	}
	wg.Wait()
	fmt.Println("context cause:", context.Cause(ctx))
}
