package cmd

import (
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func newCommandKeno() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "keno",
		Short: "Choose one of given arguments",
		Run: func(cmd *cobra.Command, args []string) {
			favoritismSelect(cmd, args)
		},
	}

	cmd.Flags().IntP("count", "n", 1, "Number to select")
	return cmd
}

func favoritismSelect(cmd *cobra.Command, args []string) {
	count, err := cmd.Flags().GetInt("count")
	if err != nil {
		cmd.Println("Failed to parse flags")
	}

	if len(args) == 0 {
		cmd.Usage()
		return
	}

	if count >= len(args) {
		cmd.Printf("%d is out of range.\n", count)
		os.Exit(2)
	}

	if count <= 0 || len(args) == count {
		cmd.Println(strings.Join(args, "\n"))
		return
	}

	cheat := append(args[:count-1], args[count:]...)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < count; i++ {
		target := rand.Intn(len(cheat))
		cmd.Println(cheat[target])
		cheat = append(cheat[:target], cheat[target+1:]...)
	}
}
