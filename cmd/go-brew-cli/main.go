package main

import (
	"github.com/robotjoosen/go-brew/cmd/go-brew-cli/cmd"
)

func init() {
	cmd.RootCmd.Flags().StringVarP(&cmd.Coffee, "coffee", "g", "16", "coffee weight in grams")
	cmd.RootCmd.Flags().StringVarP(&cmd.Flavor, "flavor", "f", "standard", "sweet, standard, bright")
	cmd.RootCmd.Flags().StringVarP(&cmd.Concentration, "concentration", "c", "strong", "light, medium, strong")
}

func main() {
	if err := cmd.RootCmd.Execute(); err != nil {
		return
	}
}
