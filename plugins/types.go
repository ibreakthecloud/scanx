package plugins

const (
	YaraHunter    = "YaraHunter"
	SecretScanner = "SecretScanner"
)

var (
	PluginPathMap = map[string]string{
		YaraHunter:    "./plugins/yarahunter/yarahunter.so",
		SecretScanner: "./plugins/secretscanner/secretscanner.so",
	}
)
