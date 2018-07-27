package server

import (
	"encoding/json"
	"fmt"
	"github.com/continuul/random-names/command"
	"github.com/continuul/random-names/pkg/namesgenerator"
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"io"
	"k8s.io/apimachinery/pkg/util/wait"
	"math/rand"
	"net/http"
	"time"
)

type serverInfo struct {
	Cli    command.Cli
	Config *Config
}

// New creates a Cobra CLI command.
func New(cli command.Cli) *cobra.Command {
	stopCh := SetupSignalHandler()

	server := &serverInfo{
		Cli:    cli,
		Config: nil,
	}

	cmd := &cobra.Command{
		Use:   "server",
		Short: "Starts the server",
		Run: func(cmd *cobra.Command, args []string) {

			PrettyPrint(server.Config)

			//cfg := DefaultConfig()
			//cfg = MergeConfig(cfg, &cmdCfg)
			//server.Config = cfg

			server.Run(stopCh)
		},
	}

	var cmdCfg Config
	cmd.Flags().StringVarP(&cmdCfg.BindAddr, "bind", "b", cmdCfg.BindAddr, "Sets the bind address for the server.")
	cmd.Flags().IntVarP(&cmdCfg.Ports.HTTP, "http-port", "p", 0, "Sets the HTTP API port to listen on.")
	cmd.Flags().BoolVar(&cmdCfg.EnableUI, "ui", false, "Enables the built-in static web UI server.")
	cmd.Flags().StringVar(&cmdCfg.UIDir, "ui-dir", "", "Path to directory containing the web UI resources.")
	server.Config = &cmdCfg
	return cmd
}

func PrettyPrint(v interface{}) (err error) {
	b, err := json.MarshalIndent(v, "", "  ")
	if err == nil {
		fmt.Println(string(b))
	}
	return
}

// Run the server.
func (s *serverInfo) Run(stopCh <-chan struct{}) error {
	done := make(chan struct{})

	// handle tls, sockets, addresses here (early)...

	mux := http.NewServeMux()
	mux.HandleFunc("/", s.index)
	//mux.HandleFunc("/ui/", s.ui)
	mux.HandleFunc("/api/v1/name", s.generate)
	mux.HandleFunc("/healthz", s.healthz)
	if s.Config.UIDir != "" {
		mux.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(http.Dir(s.Config.UIDir))))
	} else if s.Config.EnableUI {
		mux.Handle("/ui/", http.StripPrefix("/ui/", http.FileServer(assetFS())))
	}

	go wait.Until(func() {
		err := http.ListenAndServe(":9000", mux)
		if err != nil {
			glog.Errorf("Starting health server failed: %v", err)
		}
	}, 5*time.Second, wait.NeverStop)

	fmt.Fprintln(s.Cli.Out(), "Random names server running!")
	fmt.Fprintf(s.Cli.Out(), "   Version: '%s'\n", command.Version)
	fmt.Fprintf(s.Cli.Out(), "   Bind Addr: %v (HTTP: %d, HTTPS: %d)\n", s.Config.BindAddr,
		s.Config.Ports.HTTP, s.Config.Ports.HTTPS)

	select {
	case <-done:
		break
	case <-stopCh:
		break
	}

	return nil
}

func (s *serverInfo) index(resp http.ResponseWriter, req *http.Request) {
	// Check if this is a non-index path
	if req.URL.Path != "/" {
		resp.WriteHeader(http.StatusNotFound)
		return
	}

	// Redirect to the UI endpoint
	http.Redirect(resp, req, "/ui/", http.StatusMovedPermanently) // 301
}

func (s *serverInfo) healthz(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func (s *serverInfo) generate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	rand.Seed(time.Now().UTC().UnixNano())
	io.WriteString(w, fmt.Sprintf(`{"name": "%s"}`, namesgenerator.GetRandomName(0)))
}

func (s *serverInfo) ui(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Random Names Server")
}
