package cmd

import (
	"fmt"
  
	"github.com/spf13/cobra"
  )
  
var cmdVersion = &cobra.Command{
    Use:   "version",
    Short: "Print app version",
    Long: `Print app version.`,
    Run: func(cmd *cobra.Command, args []string) {
      fmt.Println("version: v0.0.0")
    },
  }
