// cmd/ptai/main.go

package main

import (
	"fmt"
	"os"

	"prototypus-ai-go/internal/core" // コアロジックをインポート

	"github.com/spf13/cobra"
)

// グローバルなフラグ変数
var (
	inputPath    string
	outputFormat string
	prompt       string
)

// rootCmd はベースとなる 'ptai' コマンドを定義します。
var rootCmd = &cobra.Command{
	Use:   "ptai",
	Short: "Prototypus AI: The data transformation tool.",
	Long:  "Prototypus AI: The data transformation tool for the 21st century.",
	Run: func(cmd *cobra.Command, args []string) {
		// 引数なしで ptai が呼ばれた場合の処理
		cmd.Help()
	},
}

// convertCmd は 'ptai convert' サブコマンドを定義します。
var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts a file from one format to another using Gemini AI.",
	Long: `Converts a file from one format to another. 
It requires input-path and output-format flags.`,
	Run: func(cmd *cobra.Command, args []string) {
		if inputPath == "" || outputFormat == "" {
			fmt.Println("Error: --input-path and --output-format are required.")
			os.Exit(1)
		}

		fmt.Printf("[CLI] Starting conversion of '%s' to '%s'...\n", inputPath, outputFormat)

		// コアロジックの呼び出し
		result, err := core.PerformConversion(inputPath, outputFormat, prompt)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("\n--- Final Converted Content ---")
		fmt.Println(result)
		fmt.Println("-------------------------------")
	},
}

// init 関数は main 関数より前に実行され、コマンドの設定を行います。
func init() {
	// convert コマンドをルートコマンドに追加
	rootCmd.AddCommand(convertCmd)

	// convert コマンドにフラグ（引数）を追加
	convertCmd.Flags().StringVarP(&inputPath, "input-path", "i", "", "The path to the input file to be converted.")
	convertCmd.Flags().StringVarP(&outputFormat, "output-format", "o", "", "The desired output format (e.g., 'html', 'json', 'yaml').")
	convertCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "Optional. A custom prompt to guide the transformation.")

	// 必須フラグの設定
	convertCmd.MarkFlagRequired("input-path")
	convertCmd.MarkFlagRequired("output-format")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
