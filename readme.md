
# ======= 补全的几种方式及示例 ====
### 命令补全:
 1. Cobra 默认支持命令补全。只要将命令和子命令添加到 rootCmd 中，Cobra 就会自动处理它们的补全。例如：

 2. 自定义补全:

 - 参数补全:
 在 getCmd 命令中，通过 ValidArgsFunction 为 id 参数提供自定义补全选项
  
```sh  
ValidArgsFunction: func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
  return []string{"item1", "item2", "item3"}, cobra.ShellCompDirectiveNoFileComp
}
```

- 标志补全:

在 listCmd 命令中，通过 RegisterFlagCompletionFunc 为 --type 标志提供自定义补全选项：

```bash 
 listCmd.Flags().StringP("type", "t", "", "Type of item to list")
 listCmd.RegisterFlagCompletionFunc("type", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
     return []string{"type1", "type2", "type3"}, cobra.ShellCompDirectiveNoFileComp
 })
 ```

 **！！！ ⚠️ 完成补全的先觉条件是 安装了 bash-completion 包**

# 是否安装 bash-complection  

检查 在centos 中检查是否已安装了 bash-complection 包的方法如下：

1. `rpm -q bash-completion`
2. `type compgen`
3. `test -f /etc/bash_completion.d/authselect-completion.sh `, 检查文件是否存在


#  bash-completion 常用函数

bash-completion 脚本主要定义了以下几个常用的函数，用于实现命令补全：

1. _init_completion: 初始化补全函数的环境。
2. _filedir: 用于文件和目录的补全。
3. _get_comp_words_by_ref: 获取命令行中某些特定位置的补全词。
4. _expand: 处理路径名展开（如 ~ 和 *）。
5. _complete: 自动补全的核心函数。
6. _command_offset: 确定当前补全的命令行的位置偏移。
7. _have: 检查命令是否在当前系统中可用。

# 使用

全局生成自动补全代码

`./cobra-example completion bash > /etc/bash_completion.d/cobra-example`

### 参考示例

- 项目： https://github.com/kekek/cobra-example

- 分支： master