package server

type Config struct {
	// BindAddr is used to control the address we bind to.
	BindAddr string `mapstructure:"bind_addr"`

	// Port configurations
	Ports PortConfig

	// EnableUI enables the statically-compiled assets for the Consul web UI and
	// serves them at the default /ui/ endpoint automatically.
	EnableUI bool `mapstructure:"ui"`

	// UIDir is the directory containing the Web UI resources.
	// If provided, the UI endpoints will be enabled.
	UIDir string `mapstructure:"ui_dir"`

	// CAFile is a path to a certificate authority file. This is used with VerifyIncoming
	// or VerifyOutgoing to verify the TLS connection.
	CAFile string `mapstructure:"ca_file"`

	// CAPath is a path to a directory of certificate authority files. This is used with
	// VerifyIncoming or VerifyOutgoing to verify the TLS connection.
	CAPath string `mapstructure:"ca_path"`

	// CertFile is used to provide a TLS certificate that is used for serving TLS connections.
	// Must be provided to serve TLS connections.
	CertFile string `mapstructure:"cert_file"`

	// KeyFile is used to provide a TLS key that is used for serving TLS connections.
	// Must be provided to serve TLS connections.
	KeyFile string `mapstructure:"key_file"`
}

// Ports is used to simplify the configuration by
// providing default ports, and allowing the addresses
// to only be specified once
type PortConfig struct {
	HTTP  int // HTTP API
	HTTPS int // HTTPS API
}

func DefaultConfig() *Config {
	return &Config{
		BindAddr: "0.0.0.0",
		Ports: PortConfig{
			HTTP:  9000,
			HTTPS: -1,
		},
	}
}

// MergeConfig merges two configurations together to make a single new
// configuration.
func MergeConfig(a, b *Config) *Config {
	var result = *a

	if b.BindAddr != "" {
		result.BindAddr = b.BindAddr
	}
	if b.Ports.HTTP != 0 {
		result.Ports.HTTP = b.Ports.HTTP
	}
	if b.Ports.HTTPS != 0 {
		result.Ports.HTTPS = b.Ports.HTTPS
	}
	if b.EnableUI {
		result.EnableUI = true
	}
	if b.UIDir != "" {
		result.UIDir = b.UIDir
	}

	return &result
}
