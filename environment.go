package main

import (
	"os"
)

var (
	listenPort     = os.Getenv("LISTEN_PORT")
	guacURL        = os.Getenv("GUAC_URL")
	guacUser       = os.Getenv("GUAC_USER")
	guacPass       = os.Getenv("GUAC_PASS")
	guacDatasource = os.Getenv("GUAC_DATASOURCE")
	chatURL        = os.Getenv("CHAT_HOOK")
	queue          []string
)
