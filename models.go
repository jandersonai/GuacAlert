package main

type Server struct {
	Name              string                 `json:"name"`
	Identifier        string                 `json:"identifier"`
	ParentIdentifier  string                 `json:"parentIdentifier"`
	Protocol          string                 `json:"protocol"`
	Attributes        map[string]interface{} `json:"attributes"`
	ActiveConnections int                    `json:"activeConnections"`
	LastActive        int64                  `json:"lastActive"`
}

type Connection struct {
	Identifier           string `json:"identifier"`
	ConnectionIdentifier string `json:"connectionIdentifier"`
	StartDate            int64  `json:"startDate"`
	RemoteHost           string `json:"remoteHost"`
	Username             string `json:"username"`
	Connectable          bool   `json:"connectable"`
}

type ConnectionDetail struct {
	ServerName   string `json:"serverName"`
	Username     string `json:"username"`
	ConnectionID string `json:"connectionID"`
}
