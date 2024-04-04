# GO-UAUTH

Minimalist Go library for Google Oauth authentication using Gin. You
can read the manual page for more info [here](./uauth.1.md).

## Go documentation

    package uauth // import "github.com/harkaitz/go-uauth"
    
    func CLIAuthorizationFile() (file string)
    func CLIConfigurationFile() (file string)
    func CLIPleaseAuthenticateError() (err error)
    func GetLanguage(ctx *gin.Context, accepted ...string) (language string)
    func LoadJSON(s *Settings, file string) (err error)
    func RedirectToLogin(ctx *gin.Context)
    func RedirectToLogout(ctx *gin.Context, page string)
    type Authority struct{ ... }
    type Message struct{ ... }
    type Settings struct{ ... }
    type UID string
    type User goth.User
        func CLIGetAuthentication() (user User, found bool, err error)
        func VerifyUser(ctx *gin.Context) (user User, found bool)

## Go struct Authority

    package uauth // import "."
    
    type Authority struct {
        Settings   *Settings
        CLIChannel chan User
    }
        Authority is the main object to be used to authenticate.
    
    func (u Authority) CLIAuthenticate() (user User, err error)
    func (u Authority) CLIAuthenticateURL() (url string)

## Go programs

    Usage: uauth -fp
    
    Authenticate using Google Oauth, more in man "uauth(1)".
    
    Copyright (c) 2023 - Harkaitz Agirre - All rights reserved.

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
