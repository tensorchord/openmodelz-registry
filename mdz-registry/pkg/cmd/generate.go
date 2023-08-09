package cmd

import (
	"github.com/spf13/cobra"
)

var (
	output string
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:     "generate",
	Short:   "Generate JSON",
	Long:    `Generate JSON`,
	Example: `mdz-registery generate ./static`,
	GroupID: "basic",
	PreRunE: commandInit,
	Args:    cobra.ExactArgs(1),
	RunE:    commandGenerate,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	generateCmd.Flags().StringVarP(&output, "output", "o", "", "Output file")
}

func commandGenerate(cmd *cobra.Command, args []string) error {

	// Iterate the templates from the source directory
	// and generate the JSON files

	return nil
}
