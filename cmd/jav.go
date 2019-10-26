package cmd

import (
	"encoding/base64"
	"fmt"

	"github.com/spf13/cobra"
)

var JAV = &cobra.Command{
	Use: "jav",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(base64En(args[0]))
	},
}

func designation(code string) {}

func base64En(value string) string {
	v := []byte(value)
	s := base64.StdEncoding.EncodeToString(v)
	fmt.Println(s)
	return s
}
