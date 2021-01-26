package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/caarlos0/env/v6"
)

type config struct {
	GitBranch string `env:"OKTETO_GIT_BRANCH,required"`
	GitCommit string `env:"OKTETO_GIT_COMMIT,required"`
	//GitCloneDuration time.Duration `env:"OKTETO_CLONE_DURATION,required"`
}

func main() {
	cfg := config{}
	if err := env.Parse(&cfg); err != nil {
		log.Fatalf(err.Error())
	}

	http.HandleFunc("/results", cfg.results)

	log.Println(http.ListenAndServe(":8080", nil))
}

func (c config) results(w http.ResponseWriter, r *http.Request) {
	bytes, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		http.Error(w, "marshaling config to json", http.StatusInternalServerError)
	}
	fmt.Fprintf(w, string(bytes))
}
