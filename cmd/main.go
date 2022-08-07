package main

import (
	"deus-task/api"
	"deus-task/core"
	inmemory "deus-task/repository/InMemory"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	repo, _ := inmemory.NewInMemoryStore("", 1)
	service := core.NewVisitService(repo)
	h := api.NewVisitsHandler(service)
	http.HandleFunc("/get-visits", h.Get)
	http.HandleFunc("/add-visit", h.Post)

	port := httpPort()

	errs := make(chan error, 2)
	go func() {
		fmt.Println("Listening on port" + port)
		errs <- http.ListenAndServe("localhost"+port, nil)
	}()

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}

func httpPort() string {
	port := "8080"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	return fmt.Sprintf(":%s", port)
}

func choseRepo() core.VisitsRepository {
	switch os.Getenv("URL_DB") {
	case "inmemory":
		repo, err := inmemory.NewInMemoryStore("", 1)
		if err != nil {
			log.Fatal(err)
		}
		return repo
	}
	return nil
}
