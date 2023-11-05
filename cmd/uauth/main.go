package main

import (
	"fmt"
	"os"
	"github.com/harkaitz/go-uauth"
	"github.com/pborman/getopt/v2"
	"github.com/gin-gonic/gin"
)

const help string =
`Usage: uauth -fp

Authenticate using Google Oauth, more in man "uauth(1)".

Copyright (c) 2023 - Harkaitz Agirre - All rights reserved.`

func main() {
	
	var err    error
	var user   uauth.User
	var found  bool
	
	defer func() {
		if err != nil {
			fmt.Fprintf(os.Stderr, "uauth: error: %v\n", err.Error())
			os.Exit(1)
		}
	}()
	
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
		
		uAuthSettings, err = uauth.LoadSettings(uauth.CLIConfigurationFile())
		if err != nil { return }
		
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
		
		uAuth, err = uauth.NewAuthority(uAuthSettings, r, true)
		if err != nil { return }
		
		go r.Run(uAuthSettings.Domain)
		user, err = uAuth.CLIAuthenticate()
		if err != nil { return }
	}
	
	if *pFlag == true {
		fmt.Println(user)
	}
}

