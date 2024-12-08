package configs

const (
	ModelConsole    = "Console"
	ModelController = "Controller"
)

const (
	ControllerLoginPath         = "api/login"
	ControllerLogoutPath        = "api/logout"
	ControllerClientActivePath  = "v2/api/site/default/clients/active"
	ControllerClientHistoryPath = "v2/api/site/default/clients/history"
	ControllerCmdRemovalPath    = "api/s/default/cmd/stamgr"
	ControllerWlanConfigPath    = "/v2/api/site/default/wlan/enriched-configuration"
)

const (
	ConsoleApiPrefix         = "proxy/network/"
	ConsoleLoginPath         = "api/auth/login"
	ConsoleLogoutPath        = "api/auth/logout"
	ConsoleClientActivePath  = ConsoleApiPrefix + ControllerClientActivePath
	ConsoleClientHistoryPath = ConsoleApiPrefix + ControllerClientHistoryPath
	ConsoleCmdRemovalPath    = ConsoleApiPrefix + ControllerCmdRemovalPath
	ConsoleWlanConfigPath    = ConsoleApiPrefix + ControllerWlanConfigPath
)
