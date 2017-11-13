package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"time"
)

var (
	g errgroup.Group
)

func main() {

	s1 := &http.Server{
		Addr:         ":8080",
		Handler:      PingPongHandler(),
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
	}

	s2 := &http.Server{
		Addr:         ":8081",
		Handler:      ProfileHandler(),
		ReadTimeout:  500 * time.Millisecond,
		WriteTimeout: 500 * time.Millisecond,
	}

	g.Go(func() error {
		return s1.ListenAndServe()
	})
	g.Go(func() error {
		return s2.ListenAndServe()
	})

	if err := g.Wait(); err != nil {
		fmt.Fprintf(os.Stderr, "err? %v\n", err)
	}

}
