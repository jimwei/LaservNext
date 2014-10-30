// laser framework types
package types

//laser config struct
type LaserConfig struct {
	Connection ConnectionInfo
}

//database connection struct
type ConnectionInfo struct {
	Server   string
	Database string
	User     string
	Password string
}
