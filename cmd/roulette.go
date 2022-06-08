package cmd

import (
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

func newCommandRoulette() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "roulette",
		Short: "Choose one of the input lines to standard output. \nProbability increases as the index grows.",
		Run: func(cmd *cobra.Command, args []string) {
			escalationSelect(cmd, args)
		},
	}
	return cmd
}

func escalationSelect(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		cmd.Usage()
		return
	}

	cheat := make([]string, 0)
	cheatlen := len(args) * (len(args) + 1) / 2
	for k, v := range args {
		for i := 0; i <= k; i++ {
			cheat = append(cheat, v)
		}
	}
	rand.Seed(time.Now().UnixNano())
	cmd.Println(cheat[rand.Intn(cheatlen)])
}
