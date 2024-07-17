package main

import (
	"log"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cobra-example",
	Short: "练手项目，演示如何使用cobra的自动补全",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all items",
	Long: `
演示 flag 自动补全定义：
listCmd.RegisterFlagCompletionFunc("type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// 返回可能的补全选项
		return []string{"type1", "type2", "type3"}, cobra.ShellCompDirectiveNoFileComp
	})
`,
	Run: func(cmd *cobra.Command, args []string) {
		// Command logic here
		typ, err := cmd.Flags().GetString("type")
		if err != nil {
			log.Println("获取type 失败", err)
			return
		}

		log.Printf("获取type 值：%s \n\n", typ)
	},
}

var getCmd = &cobra.Command{
	Use:   "get [id]",
	Short: "Get an item by ID",
	Long: `
演示 无 flag 自动补全定义：

	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// 返回可能的补全选项
		return []string{"item1", "item2", "item3"}, cobra.ShellCompDirectiveNoFileComp
	},
`,

	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		// Command logic here

		id := args[0]

		log.Printf("获取id 值：%s \n\n", id)
	},
	// 自定义补全 get 命令的 id 参数
	ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// 返回可能的补全选项
		return []string{"item1", "item2", "item3"}, cobra.ShellCompDirectiveNoFileComp
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(getCmd)

	// 为 list 命令的 type 标志添加自定义补全
	listCmd.Flags().StringP("type", "t", "", "Type of item to list")
	listCmd.RegisterFlagCompletionFunc("type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// 返回可能的补全选项
		return []string{"type1", "type2", "type3"}, cobra.ShellCompDirectiveNoFileComp
	})
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}
