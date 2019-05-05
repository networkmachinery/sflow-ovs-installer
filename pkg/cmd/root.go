package cmd

import (
        "fmt"
        "io"
        "os"

        "github.com/spf13/cobra"
)

var cfgFile string


type createCmd struct {
        collectorIP    helmpath.Home
        collectorPort    string
             io.Writer
        starter string
}

// rootCmd represents the base command when called without any subcommands
var sflowInstallerCmd = &cobra.Command{
        Use:   "networkmachinery-sflow",
        Short: "A tool to install to setup and install sFlow for OVS with Kubernetes",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
        if err := rootCmd.Execute(); err != nil {
                fmt.Println(err)
                os.Exit(1)
        }
}

func init() {
        cobra.OnInitialize()
        rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}