package prompt

import (
	"errors"
	"fmt"
	"slices"

	"github.com/AlecAivazis/survey/v2"
	"github.com/twiggotronix/host-updater/network"
)

func SelectNetwork() (*string, error) {
	localAddressesFactoy := network.LocalAddressesFactoy{}
	interfaces, err := localAddressesFactoy.GetLocalAddresses().GetNetworkInterfaces()
	if err != nil {
		fmt.Println(err)
	}
	if len(interfaces) == 0 {
		return nil, errors.New("No network interfaces to select")
	}
	selectedAddress := ""
	if len(interfaces) > 1 {
		interfaceAdresses := network.ToIpAdresses(interfaces)
		prompt := &survey.Select{
			Message: "Pick a network interface",
			Options: interfaceAdresses,
			Description: func(interfaceAdress string, index int) string {
				intfIndex := slices.IndexFunc(interfaces, func(c network.Intf) bool { return c.Addr == interfaceAdress })
				return fmt.Sprintf("%s (%s)", interfaces[intfIndex].Name, interfaceAdress)
			},
		}
		surveyError := survey.AskOne(prompt, &selectedAddress)
		if surveyError != nil {
			fmt.Println(surveyError.Error())
			return nil, surveyError
		}
	} else {
		fmt.Println("Only one interface found, selecting...")
		selectedAddress = interfaces[0].Addr
	}
	return &selectedAddress, nil
}
