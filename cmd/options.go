package main

type Option struct {
	LogLevel       string
	Port           int
	Driver         string
	Connection     string
	PrivateKeyPath string
	PublicKeyPath  string
	SyncDB         bool
	HTTPScheme     string //服务器有可能放在nginx做了SSL卸载，所以无法直接判断是https还是http
	RedisHost      string
	RedisPort      int
	RedisDB        int
}
