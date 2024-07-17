package get

import (
	"log"

	"github.com/spf13/cobra"
)

var GetCmd = &cobra.Command{
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
