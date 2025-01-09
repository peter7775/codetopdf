package cmd

import (
	"fmt"
	"codetopdf/pkg/pdfprint"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "codetopdf [path]",
	Short: "Print code repository folder in one PDF file.",
	Long: `Print code repository folder in one PDF file. `,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
        err := pdfprint.PdfPrint(args[0])
        if err != nil {
            fmt.Println("Error:", err)
            os.Exit(1)
        }
    },
}


func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


