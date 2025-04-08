package language

import (
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

func HttpClientExample() {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	req := createTodoListRequest()

	res, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println("Error while closing response body", err.Error())
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("unexpected status: got %v", res.Status))
	}
	fmt.Println(res.Header.Get("Content-Type"))

	data := decodeResponseBody(res)
	fmt.Printf("%+v\n", data)
}

func ServeMux() {
	terribleSecurity := terribleSecurityProvider("GOPHER")

	mux := http.NewServeMux()

	// to apply the middleware to just the single route
	mux.Handle("/hello", terribleSecurity(requestTimer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			_, err := w.Write([]byte("Hello!\n"))
			if err != nil {
				panic(err)
			}
		}))))

	// or to apply the middleware to every route in the mux:
	//
	//	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
	//		w.Write([]byte("Hello!\n"))
	//	})
	//	mux = terribleSecurity(RequestTimer(mux))

	s := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	err := s.ListenAndServe()
	if err != nil {
		if err != http.ErrServerClosed {
			panic(err)
		}
	}
}

func requestTimer(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		end := time.Now()
		slog.Info("request time", "path", r.URL.Path, "duration", end.Sub(start))
	})
}

var securityMsg = []byte("You didn't give the secret password\n")

func terribleSecurityProvider(password string) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Header.Get("X-Secret-Password") != password {
				w.WriteHeader(http.StatusUnauthorized)
				_, err := w.Write(securityMsg)
				if err != nil {
					panic(err)
				}
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
