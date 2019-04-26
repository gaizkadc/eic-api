/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package main

import (
	"github.com/nalej/eic-api/cmd/eic-api/commands"
	"github.com/nalej/eic-api/version"
)

// MainVersion with the application version.
var MainVersion string
// MainCommit with the commit id.
var MainCommit string

func main() {
	version.AppVersion = MainVersion
	version.Commit = MainCommit
	commands.Execute()
}
