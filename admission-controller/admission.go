package main

import (
	"flag"

	"github.com/qqbuby/kube-admission/webhook"
	"github.com/spf13/cobra"

	"k8s.io/klog/v2"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "kube",
		Version: "1.32",
	}

	rootCmd.AddCommand(webhook.CmdWebhook)

	loggingFlags := &flag.FlagSet{}
	klog.InitFlags(loggingFlags)
	rootCmd.PersistentFlags().AddGoFlagSet(loggingFlags)
	rootCmd.Execute()
}
