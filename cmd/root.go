package cmd

import (
	"github.com/mizuki1412/go-core-kit/init/initkit"
	"github.com/mizuki1412/go-core-kit/service/logkit"
	"github.com/mizuki1412/go-core-kit/service/restkit"
	"github.com/spf13/cobra"
	"wechat-work-pusher/constant"
	"wechat-work-pusher/controller"
)

func init() {
	DefFlags(rootCmd)
}

var rootCmd = &cobra.Command{
	Use: "Wechat-Work-Pusher",
	Run: func(cmd *cobra.Command, args []string) {
		initkit.BindFlags(cmd)
		restkit.AddActions(controller.Init)
		_ = restkit.Run()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		logkit.Fatal(err.Error())
	}
}

func DefFlags(cmd *cobra.Command) {
	cmd.Flags().String(constant.ConfigKeyWorkCorpId, "", "企业ID")
	cmd.Flags().String(constant.ConfigKeyWorkAgentId, "", "应用AgentID")
	cmd.Flags().String(constant.ConfigKeyWorkCorpSecret, "", "应用SecretID")
	cmd.Flags().String(constant.ConfigKeyDefaultReceiver, "", "默认被推送ID")
	cmd.Flags().String(constant.ConfigKeyToken, "", "接口调用Token")
}
