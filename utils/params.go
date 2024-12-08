package utils

import (
	"errors"
	"flag"
	"fmt"
	"golang.org/x/term"
	"syscall"
	"unifi-toolkits/configs"
)

type Args struct {
	Model    string
	IP       string
	Port     string
	User     string
	Password string
	DryRun   bool
}

var InvalidParams = errors.New("invalid params")
var ValidModels = map[string]bool{
	configs.ModelConsole:    true,
	configs.ModelController: true,
}

func NewArgs() (*Args, error) {
	var model string
	var ip string
	var user string
	var password string
	var port string
	var dryRun bool

	flag.StringVar(&model, "m", "", "")
	flag.StringVar(&ip, "g", "", "")
	flag.StringVar(&user, "u", "", "")
	flag.StringVar(&port, "p", "", "")
	flag.BoolVar(&dryRun, "d", false, "")
	flag.Usage = func() {
		usage := `Usages:
-m <Console/Controller>
    Unifi model
-g
    Unifi console/controller ip address
-p
    Unifi console/controller port (Optional)
-u
    Unifi console/controller user
-d <true>
    Dry run`
		fmt.Println(usage)
	}
	flag.Parse()

	if model == "" || ip == "" || user == "" {
		flag.Usage()
		return nil, InvalidParams
	}
	if _, ok := ValidModels[model]; !ok {
		flag.Usage()
		return nil, InvalidParams
	}

	fmt.Println("Please enter password:")
	bytePassword, err := term.ReadPassword(syscall.Stdin)
	if err != nil {
		fmt.Println(err)
		return nil, InvalidParams
	}
	password = string(bytePassword)

	args := &Args{
		Model:    model,
		IP:       ip,
		Port:     port,
		User:     user,
		Password: password,
		DryRun:   dryRun,
	}
	return args, nil
}

type Params struct {
	LoginPath         string
	LogoutPath        string
	ClientActivePath  string
	ClientHistoryPath string
	CmdRemovalPath    string
	WlanConfigPath    string
}

func NewConsoleParams() *Params {
	params := &Params{
		LoginPath:         configs.ConsoleLoginPath,
		LogoutPath:        configs.ConsoleLogoutPath,
		ClientActivePath:  configs.ConsoleClientActivePath,
		ClientHistoryPath: configs.ConsoleClientHistoryPath,
		CmdRemovalPath:    configs.ConsoleCmdRemovalPath,
		WlanConfigPath:    configs.ConsoleWlanConfigPath,
	}
	return params
}

func NewControllerParams() *Params {
	params := &Params{
		LoginPath:         configs.ControllerLoginPath,
		LogoutPath:        configs.ControllerLogoutPath,
		ClientActivePath:  configs.ControllerClientActivePath,
		ClientHistoryPath: configs.ControllerClientHistoryPath,
		CmdRemovalPath:    configs.ControllerCmdRemovalPath,
		WlanConfigPath:    configs.ControllerWlanConfigPath,
	}
	return params
}
