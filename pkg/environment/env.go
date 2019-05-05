package flags

import (
	"os"

	"github.com/spf13/pflag"
)

// flagToEnv maps flag names to environment variables
var flatToEnv = map[string]string{
	"collector_ip":   "COLLECTOR_IP",
	"collector_port": "COLLECTOR_PORT",
	"agent_ip":       "AGENT_IP",
	"header_bytes":   "HEADER_BYTES",
	"sampling_n":     "SAMPLING_N",
	"polling_secs":   "POLLING_SECS",
}

func FlagToEnv(fs *pflag.FlagSet) {
	for name, ennironmentVariable := range flatToEnv {
		setFlagFromEnv(name, ennironmentVariable, fs)
	}
}

// setFlagFromEnv sets flags to the value of environment variables
func setFlagFromEnv(name, environmentVariable string, fs *pflag.FlagSet) {
	if fs.Changed(name) {
		return
	}
	if v, ok := os.LookupEnv(environmentVariable); ok {
		fs.Set(name, v)
	}
}
