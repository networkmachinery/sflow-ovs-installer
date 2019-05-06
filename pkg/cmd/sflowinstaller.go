package cmd

import (
	"github.com/digitalocean/go-openvswitch/ovs"
	"github.com/networkmachinery/sflow-ovs-installer/log"
	"github.com/networkmachinery/sflow-ovs-installer/pkg/environment"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type sFlowInstallerCmd struct {
	bridgeName    string
	collectorIP   string
	collectorPort string
	agentIP       string
	headerBytes   string
	samplingN     string
	pollingSecs   string
}

// NewCmdSFlowInstaller creates a new sFlowInstallerCmd
func NewCmdSFlowInstaller(args []string) *cobra.Command {
	sFlowCmd := sFlowInstallerCmd{}
	cmd := &cobra.Command{
		Use:   "sflow-ovs-installer",
		Short: "A tool to install to setup and install sFlow for OVS with Kubernetes",
		RunE: func(cmd *cobra.Command, args []string) error {
			return sFlowCmd.run()
		},
	}
	flags := cmd.PersistentFlags()
	sFlowCmd.AddFlags(flags)
	environment.FlagToEnv(flags)

	cmd.AddCommand(versionCmd)
	flags.Parse(args)

	return cmd
}

// AddFlags binds flags to the given flagset.
func (s *sFlowInstallerCmd) AddFlags(fs *pflag.FlagSet) {
	fs.StringVar(&s.bridgeName, "bridge-name", "", "the name of the OVS bridge to configure")
	fs.StringVar(&s.collectorIP, "collector-ip", "", "is the sFlow collector IP")
	fs.StringVar(&s.collectorPort, "collector-port", "6343", "is the default port number for sFlowTrend")
	fs.StringVar(&s.agentIP, "agent-ip", "", "indicates the interface / ip that the sFlow agent should send traffic from")
	fs.StringVar(&s.headerBytes, "header-bytes", "128", "the header bytes")
	fs.StringVar(&s.samplingN, "sampling-n", "64", "is the type of sampling")
	fs.StringVar(&s.pollingSecs, "polling-secs", "10", "frequency of sampling i.e., samples/sec")
}

// TODO: Ensure OVS is installed and validated flags
func (s *sFlowInstallerCmd) validate() error {
	return nil
}

func (s *sFlowInstallerCmd) run() error {
	ovsService := ovs.New()
	createdSFLowID, err := ovsService.VSwitch.CreateSFlow(s.bridgeName, s.agentIP, s.collectorIP, s.collectorPort, s.headerBytes, s.samplingN, s.pollingSecs)
	if err != nil {
		return err
	}

	log.Infof("sFlow was succesfully configured with ID: %s", createdSFLowID)
	return nil
}
