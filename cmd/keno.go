package cmd

import (
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

type option struct {
	Count int `validate:"min=0"`
}

func newCommandKeno() *cobra.Command {
	var o = &option{}
	cmd := &cobra.Command{
		Use:   "keno",
		Short: "Choose one of given arguments",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return validateParams(*o)
		},
		Run: func(cmd *cobra.Command, args []string) {
			favoritismSelect(cmd, args, o)
		},
		SilenceErrors: true,
		SilenceUsage:  true,
	}

	cmd.Flags().IntVarP(&o.Count, "count", "n", 1, "Number to select")
	return cmd
}

func favoritismSelect(cmd *cobra.Command, args []string, options *option) {
	if len(args) == 0 {
		cmd.Usage()
		return
	}

	if options.Count >= len(args) {
		cmd.Printf("%d is out of range.\n", options.Count)
		os.Exit(2)
	}

	if len(args) == options.Count {
		cmd.Println(strings.Join(args, "\n"))
		return
	}

	cheat := append(args[:options.Count-1], args[options.Count:]...)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < options.Count; i++ {
		target := rand.Intn(len(cheat))
		cmd.Println(cheat[target])
		cheat = append(cheat[:target], cheat[target+1:]...)
	}
}
