package environment

import (
	"fmt"
	"os"

	"github.com/spf13/pflag"
)

// flagToEnv maps flag names to environment variables
var flatToEnv = map[string]string{
	"bridge-name":    "BRIDGE_NAME",
	"collector-ip":   "COLLECTOR_IP",
	"collector-port": "COLLECTOR_PORT",
	"agent-ip":       "AGENT_IP",
	"header-bytes":   "HEADER_BYTES",
	"sampling-n":     "SAMPLING_N",
	"polling-secs":   "POLLING_SECS",
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
		fmt.Println(name, environmentVariable, v)
		fs.Set(name, v)
	}
}
