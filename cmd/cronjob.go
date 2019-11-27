package cmd

import "github.com/spf13/cobra"

var CronJobCmd = &cobra.Command{
	Use:   "cronJob",
	Short: "execute cronJob",
	PreRun: func(cmd *cobra.Command, args []string) {

	},
	Run: func(cmd *cobra.Command, args []string) {

	},
	PostRun: func(cmd *cobra.Command, args []string) {

	},
}

func heartBeat() {
	// * */10 * * *

}
