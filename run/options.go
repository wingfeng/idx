package main

type Option struct {
	Port           int
	Driver         string
	Connection     string
	PrivateKeyPath string
	PublicKeyPath  string
	SyncDB         bool
}
