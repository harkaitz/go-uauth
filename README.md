# GO-UAUTH

Minimalist Go library for Google Oauth authentication using Gin. You
can read the manual page for more info [here](./uauth.1.md).

## Go documentation

    package uauth // import "github.com/harkaitz/go-uauth"
    
    func CLIAuthorizationFile() (file string)
    func CLIConfigurationFile() (file string)
    func CLIPleaseAuthenticateError() (err error)
    type Authority struct{ ... }
        func NewAuthority(s Settings, e *gin.Engine, useCLI bool) (u Authority, err error)
    type Message struct{ ... }
    type Settings struct{ ... }
        func LoadSettings(file string) (s Settings, err error)
    type UID string
    type User goth.User
        func CLIGetAuthentication() (user User, found bool, err error)

## Go struct Authority

    package uauth // import "."
    
    type Authority struct {
        Settings   Settings
        CLIChannel chan User
    }
        Authority is the main object to be used to authenticate.
    
    func NewAuthority(s Settings, e *gin.Engine, useCLI bool) (u Authority, err error)
    func (u Authority) CLIAuthenticate() (user User, err error)
    func (u Authority) CLIAuthenticateURL() (url string)
    func (u Authority) VerifyUser(ctx *gin.Context) (user User, found bool)

## Go programs

    Usage: uauth -fp
    
    Authenticate using Google Oauth, more in man "uauth(1)".
    
    Copyright (c) 2023 - Harkaitz Agirre - All rights reserved.

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
