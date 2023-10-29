package main

import (
	"fmt"
	"os"
	"encoding/json"
	"github.com/harkaitz/go-uauth"
	"github.com/pborman/getopt/v2"
	"github.com/gin-gonic/gin"
)

const help string =
`Usage: uauth -fp

Authenticate using Google Oauth, more in man "uauth(1)".

Copyright (c) 2023 - Harkaitz Agirre - All rights reserved.`

func main() {
	
	var uAuthSettings uauth.Settings
	var uAuth         uauth.Authority
	var r            *gin.Engine
	var err           error
	var user          uauth.User
	var userB       []byte
	var found         bool
	
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
	
	
	gin.SetMode(gin.ReleaseMode)
	r = gin.New()
	
	err = uauth.LoadSettings(&uAuthSettings, uauth.CLIConfigurationFile())
	if err != nil { return }
	
	uAuth, err = uauth.NewAuthority(uAuthSettings, r, true)
	if err != nil { return }
	
	if *fFlag == false {
		user, found, err = uAuth.CLIGetAuthentication()
		if err != nil { return }
	}
	
	if found == false {
		go r.Run(uAuthSettings.Domain)
		user, err = uAuth.CLIAuthenticate()
		if err != nil { return }
	}
	
	if *pFlag == true {
		userB, err = json.Marshal(user)
		if err != nil { return }
		fmt.Println(string(userB))
	}
}

