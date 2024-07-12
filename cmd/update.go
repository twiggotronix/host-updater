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
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		hostFile, _ := cmd.Flags().GetString("dest")
		flag, _ := cmd.Flags().GetString("flag")
		dryRun, _ := cmd.Flags().GetBool("dryRun")

		selectedAddress, err := prompt.SelectNetwork()
		if err != nil {
			panic("Could not select network")
		}
		fmt.Printf("Selected Address : %s\n", selectedAddress)
		fmt.Printf("hostFile : %s\n", hostFile)
		fmt.Printf("flag : %s\n", flag)
		fmt.Printf("dryRun : %b\n", dryRun)

		utils.SetNewIp(hostFile, flag, selectedAddress, dryRun)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)

	updateCmd.Flags().StringP("hostFile", "h", constants.DefaultHostFilePath, "The file to edit")
	updateCmd.Flags().StringP("flag", "f", "[location-host]", "the flag used in the host file")
	updateCmd.Flags().BoolP("dryRun", "d", false, "if true, don't try to write to the file")
}
