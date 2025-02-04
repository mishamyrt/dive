package types

// HostNameParts is the part of host name
type HostNameParts struct {
	User      string
	Host      string
	Namespace string
}

// RSyncFile is description of rsync target file
type RSyncFile struct {
	IsRemote bool
	Path     string
}

// RemotesList is list with namespace keys and URL values
type RemotesList struct {
	Remotes map[string]string
}

// NamespaceHostConfig is part of main configuration file
type NamespaceHostConfig struct {
	Port, User string `yaml:",omitempty"`
}

// GlobalConfig is main configuration file struct
type GlobalConfig struct {
	Namespaces map[string]NamespaceHostConfig
}

// LocalConfig is local configuration file with aliases
type LocalConfig struct {
	Aliases map[string]string
}

// Host is basic SSH connection config
type Host struct {
	Host            string
	Port, User, Key string `yaml:",omitempty"`
}

// RemoteFile is a combination of file path and host
type RemoteFile struct {
	Host     Host
	FilePath string
}

// NamespaceConfig is list of hosts
type NamespaceConfig struct {
	Namespace string
	Hosts     map[string]Host
}
