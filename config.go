package main

import "github.com/stanchan/ansible-agent/ansible"

type Config struct {
	SSL  SSLSection
	Ldap ansible.LdapOptions
}

type SSLSection struct {
	Enabled     bool
	Certificate string
	PrivateKey  string `toml:"private_key"`
	ClientCA    string `toml:"client_ca"`
}

func DefaultConfig() *Config {
	return &Config{}
}
