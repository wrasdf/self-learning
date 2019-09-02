package main

import (
  "fmt"
  "context"
  "log"
  "net/http"
  "os"
  "os/signal"
  "time"
  "encoding/json"

  "github.com/gorilla/mux"
)

func main() {

  fmt.Println("Go API Server")
  var gracefulTimeout time.Duration = time.Second * 15

  r := mux.NewRouter()

  apiv1 := r.PathPrefix("/api/v1").Subrouter()

  apiv1.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{"health": "ok"})
	})

  // r.HandleFunc("/employees", ALLEmployees).Methods("GET")
  // r.HandleFunc("/employees/{id}", GetEmployee).Methods("GET")
  // r.HandleFunc("/employees/{id}", UpdateEmployee).Methods("PUT")
  // r.HandleFunc("/employees/{id}", UpdateEmployee).Methods("DELETE")

  http.Handle("/", r)

  srv := &http.Server{
        Addr:         "0.0.0.0:8081",
        // Good practice to set timeouts to avoid Slowloris attacks.
        WriteTimeout: time.Second * 15,
        ReadTimeout:  time.Second * 15,
        IdleTimeout:  time.Second * 60,
        Handler: r, // Pass our instance of gorilla/mux in.
    }

  // Run our server in a goroutine so that it doesn't block.
  go func() {
      fmt.Println("http server running on 0.0.0.0:8081")
      if err := srv.ListenAndServe(); err != nil {
          log.Println(err)
      }
  }()

  c := make(chan os.Signal, 1)
  // We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
  // SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
  signal.Notify(c, os.Interrupt)

  // Block until we receive our signal.
  <-c

  // Create a deadline to wait for.
  ctx, cancel := context.WithTimeout(context.Background(), gracefulTimeout)
  defer cancel()

  // Doesn't block if no connections, but will otherwise wait until the gracefulTimeout deadline.
  srv.Shutdown(ctx)
  // Optionally, you could run srv.Shutdown in a goroutine and block on
  // <-ctx.Done() if your application should wait for other services
  // to finalize based on context cancellation.
  log.Println("shutting down")
  os.Exit(0)
}
