package list

import (
	"log"

	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
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

func init() {
	// 为 list 命令的 type 标志添加自定义补全
	ListCmd.Flags().StringP("type", "t", "", "Type of item to list")
	ListCmd.RegisterFlagCompletionFunc("type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		// 返回可能的补全选项
		return []string{"type1", "type2", "type3"}, cobra.ShellCompDirectiveNoFileComp
	})
}
