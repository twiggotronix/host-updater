package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/twiggotronix/host-updater/constants"
	"github.com/twiggotronix/host-updater/prompt"
	"github.com/twiggotronix/host-updater/utils"
)

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Update your host file interactively or automagically",
	Long:  `This CLI tool will help you update your host file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run command")
		hostFile, _ := cmd.Flags().GetString("dest")
		flag, _ := cmd.Flags().GetString("flag")
		dryRun, _ := cmd.Flags().GetBool("dryRun")
		preferWifi, _ := cmd.Flags().GetBool("preferWifi")
		fmt.Printf("hostFile : %s\n", hostFile)
		fmt.Printf("flag : %s\n", flag)
		fmt.Printf("dryRun : %t\n", dryRun)
		fmt.Printf("preferWifi : %t\n", preferWifi)
		networkSelectOptions := prompt.SelectNetworkOptions{
			PreferWifi: preferWifi,
		}
		selectedAddress, err := prompt.SelectNetwork(&networkSelectOptions)
		if err != nil {
			panic("Could not select network")
		}

		fileUtilsFactory := utils.FileUtilsFactoy{}
		fileUtils := fileUtilsFactory.GetFileUtils()
		fileUtils.SetNewIp(hostFile, flag, *selectedAddress, dryRun)
	},
}

func init() {
	fmt.Println("Init command")
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("dest", "e", constants.DefaultHostFilePath, "The file to edit")
	updateCmd.Flags().StringP("flag", "f", "[location-host]", "the flag used in the host file")
	updateCmd.Flags().BoolP("dryRun", "d", false, "if true, don't try to write to the file")
	updateCmd.Flags().BoolP("preferWifi", "w", false, "if true, prefere a wifi network")
}
