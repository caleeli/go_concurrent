package main

import (
	"github.com/caleeli/go_concurrent/cli"
	"github.com/caleeli/go_concurrent/storage/pokemon"
	"github.com/spf13/cobra"
)

func main() {
	repo := pokemon.NewRepository()

	rootCmd := &cobra.Command{Use: "pokemon-cli"}
	rootCmd.AddCommand(cli.InitPokemonCmd(repo))
	rootCmd.Execute()
}
