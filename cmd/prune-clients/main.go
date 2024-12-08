package main

import (
	"fmt"
	"unifi-toolkits/utils"
)

func main() {
	args, err := utils.NewArgs()
	if err != nil {
		return
	}
	unifi := utils.NewUnifi(args.Model, args.IP, args.Port, args.User, args.Password)
	defer unifi.Recover()
	err = unifi.PruneOfflineClients(args.DryRun)
	if err != nil {
		fmt.Println(err)
	}
}
