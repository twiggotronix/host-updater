package prompt

import (
	"fmt"
	"slices"

	"github.com/AlecAivazis/survey/v2"
	"github.com/twiggotronix/host-updater/network"
)

func SelectNetwork() (string, error) {
	interfaces, err := network.GetNetworkInterfaces()
	if err != nil {
		fmt.Println(err)
	}
	interfaceAdresses := network.ToIpAdresses(interfaces)
	prompt := &survey.Select{
		Message: "Pick a network interface",
		Options: interfaceAdresses,
		Description: func(interfaceAdress string, index int) string {
			intfIndex := slices.IndexFunc(interfaces, func(c network.Intf) bool { return c.Addr == interfaceAdress })
			return fmt.Sprintf("%s (%s)", interfaces[intfIndex].Name, interfaceAdress)
		},
	}
	selectedAddress := ""
	surveyError := survey.AskOne(prompt, &selectedAddress)
	if surveyError != nil {
		fmt.Println(surveyError.Error())
		return "", surveyError
	}
	return selectedAddress, nil
}
