package server

import (
	"time"
	"github.com/spf13/cobra"
	"github.com/continuul/random-names/command"
	"net/http"
	"k8s.io/apimachinery/pkg/util/wait"
	"github.com/golang/glog"
	"io"
	"math/rand"
	"github.com/continuul/random-names/pkg/namesgenerator"
	"fmt"
)

type serverInfo struct {
	cli command.Cli
}

func New(cli command.Cli) *cobra.Command {
	stopCh := SetupSignalHandler()

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Starts the server",
		Run: func(cmd *cobra.Command, args []string) {
			serverConfig := &serverInfo{
				cli: cli,
			}
			Run(serverConfig, stopCh)
		},
	}
	return cmd
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func randomNames(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rand.Seed(time.Now().UTC().UnixNano())
	io.WriteString(w, fmt.Sprintf(`{"name": "%s"}`, namesgenerator.GetRandomName(0)))
}

func home(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "hello")
}

func Run(s *serverInfo, stopCh <-chan struct{}) error {
	done := make(chan struct{})

	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/api/v1/name", randomNames)
	mux.HandleFunc("/healthz", healthCheck)
	go wait.Until(func() {
		err := http.ListenAndServe(":8000", mux)
		if err != nil {
			glog.Errorf("Starting health server failed: %v", err)
		}
	}, 5*time.Second, wait.NeverStop)

	select {
	case <-done:
		break
	case <-stopCh:
		break
	}

	return nil
}
