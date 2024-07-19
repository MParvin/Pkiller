package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	tools "github.com/mparvin/pkiller/tools"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "pkiller",
	Short: "A simple golang program that gets a pid and kill it",
	Long:  `pkiller is a simple utility to kill a process by its PID`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Fatalln("Please provide a PID to kill")
		}
		pid, err := strconv.Atoi(args[0])
		if err != nil {
			log.Fatalln(err)
		}
		if !tools.IsProcessRunning(pid) {
			log.Fatalln("Process is not running")
		}
		pCommand, pOwner, err := tools.GetProcInfo(pid)
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println("Here is information about the process you are about to kill:")
		fmt.Printf("PID: %d\nCommand: %s\nOwner: %s\n", pid, pCommand, pOwner)
		tools.KillProc(pid)
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
