package main

import (
	"os"
	"strings"
)

var (
	Token          string
	queue          []ConnectionDetail
	guacURL        = os.Getenv("GUAC_URL")
	guacUser       = os.Getenv("GUAC_USER")
	guacPass       = os.Getenv("GUAC_PASS")
	guacDatasource = os.Getenv("GUAC_DATASOURCE")
	chatURL        = os.Getenv("CHAT_HOOK")
	authBody       = strings.NewReader("username=" + guacUser + "&password=" + guacPass)
)
