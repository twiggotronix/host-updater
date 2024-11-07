package prompt

import (
	"errors"
	"fmt"
	"slices"

	"github.com/AlecAivazis/survey/v2"
	"github.com/twiggotronix/host-updater/network"
)

type SelectNetworkOptions struct {
	PreferWifi bool
}

func SelectNetwork(options *SelectNetworkOptions) (*string, error) {
	preferWifi := options != nil && options.PreferWifi

	localAddressesFactoy := network.LocalAddressesFactoy{}
	interfaces, err := localAddressesFactoy.GetLocalAddresses().GetNetworkInterfaces(&preferWifi)
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
		if preferWifi {
			fmt.Print("Wifi interface found")
		} else {
			fmt.Printf("Only one interface found")
		}
		fmt.Printf(" : selecting %s\n", interfaces[0].Addr)

		selectedAddress = interfaces[0].Addr
	}
	return &selectedAddress, nil
}
