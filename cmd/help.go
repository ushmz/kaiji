package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func newCommandHelp() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "help",
		Short: "Help about kaiji",
		Run: func(cmd *cobra.Command, args []string) {
			showHelp(cmd, args)
		},
	}

	cmd.Flags().BoolP("v", "v", false, "Set verbosity")
	cmd.Flags().Bool("vv", false, "Set verbosity")
	cmd.Flags().Bool("vvv", false, "Set verbosity")
	cmd.Flags().Bool("vvvv", false, "Set verbosity")
	cmd.Flags().Bool("vvvvv", false, "Set verbosity")
	return cmd
}

func showHelp(cmd *cobra.Command, args []string) {
	if v, _ := cmd.Flags().GetBool("v"); v {
		cmd.Parent().Usage()
		fmt.Println("There are no Easter Eggs in this program.")
		return
	}

	if v, _ := cmd.Flags().GetBool("vv"); v {
		fmt.Println("There really are no Easter Eggs in this program.")
		return
	}

	if v, _ := cmd.Flags().GetBool("vvv"); v {
		fmt.Println("Didn't I already tell you that there are no Easter Eggs in this program?")
		return
	}

	if v, _ := cmd.Flags().GetBool("vvvv"); v {
		fmt.Println("Okay, okay, if I give you an Easter Egg, will you go away?")
		return
	}

	if v, _ := cmd.Flags().GetBool("vvvvv"); v {
		fmt.Println("All right, you win.")
		time.Sleep(1 * time.Second)
		printAA(cmd)
		return
	}

	cmd.Parent().Usage()
}
