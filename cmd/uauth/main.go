package main

import (
	"fmt"
	"os"
	"github.com/harkaitz/go-uauth"
	"github.com/pborman/getopt/v2"
	"github.com/gin-gonic/gin"
)

const help string =
`Usage: uauth [-f][-p]

Helper program for programs that test libraries that authenticate/identify
clients using Google Oauth.

Test programs check the file "~/.uauth.json" exists to know whether the
user is authentified, and read user information from it with the following
functions:

    uauth.CLIGetAuthentication();

This authentication can be faked out using uauth-fake(1) or can be performed
with this program (uauth).

If the "~/.uauth.json" file doesn't exist (or "-f was specified) then
"~/.config.testing.json" file is read, and with the information gathered
an oauth authentication is performed.

The utility will listen in "127.0.0.1:8080" (remember to configure the
service to allow that URL). The authentication is performed with "xdg-open"
and the result is saved in "~/.uauth.json".

With "-p" the result is also printed out to the terminal.

Copyright (c) 2023 - Harkaitz Agirre - All rights reserved.`

func main() {
	
	var err    error
	var user   uauth.User
	var found  bool
	var home   string
	var config string
	
	defer func() {
		if err != nil {
			fmt.Fprintf(os.Stderr, "uauth: error: %v\n", err.Error())
			os.Exit(1)
		}
	}()
	
	// Read the environment.
	home, err = os.UserHomeDir()
	if err != nil { return }
	config = home + "/.config.testing.json"
	
	// Parse command line options.
	hFlag := getopt.BoolLong("help" , 'h')
	fFlag := getopt.BoolLong("force", 'f')
	pFlag := getopt.BoolLong("print", 'p')
	getopt.SetUsage(func() { fmt.Println(help) })
	getopt.Parse()
	if *hFlag { getopt.Usage(); return }
	
	if *fFlag == false {
		user, found, err = uauth.CLIGetAuthentication()
		if err != nil { return }
	}
	
	if found == false {
		var uAuth         uauth.Authority
		var uAuthSettings uauth.Settings
		var r            *gin.Engine
		
		err = uauth.LoadJSON(&uAuthSettings, config)
		if err != nil { return }
		err = uAuthSettings.VerifyUauth()
		if err != nil { return }
		
		
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		
		uAuth, err = uAuthSettings.NewAuthority(r, true)
		if err != nil { return }
		
		go r.Run(uAuthSettings.Domain)
		user, err = uAuth.CLIAuthenticate()
		if err != nil { return }
	}
	
	if *pFlag == true {
		fmt.Println(user)
	}
}

