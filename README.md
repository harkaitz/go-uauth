# GO-UAUTH

Minimalist Go library for Google Oauth authentication using Gin.

## Testing

For using the library compose the settings of your program with the
"uauth.Settings" structure so that your configuration file incorporates
the following fields:

    {
        "UseHTTPS": false,          // Required false for the CLI.
        "Domain": "127.0.0.1:8080", // Required "127.0.0.1:8080" for the CLI.
        "GoogleClientID": "...",
        "GoogleClientSecret": "...",
        "GoogleCallbackPage": "/oauthcb/google",
        "RandomString": "<random-string"
    }

When creating test programs only call "uauth.CLIGetAuthentication". When
writting the actual program use the "RedirectToLogin" and "RedirectToLogout".

## Go documentation

    package uauth // import "github.com/harkaitz/go-uauth"
    
    func CLIAuthorizationFile() (file string)
    func CLIPleaseAuthenticateError() (err error)
    func GetLanguage(ctx *gin.Context, accepted ...string) (language string)
    func LoadJSON(s *Settings, file string) (err error)
    func RedirectToLogin(ctx *gin.Context)
    func RedirectToLogout(ctx *gin.Context, page string)
    type Authority struct{ ... }
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

## Help

uauth-fake

    Usage: uauth-fake { -V | -l | -s | EMAIL }
    
    Helper script to save and set authentications for testing.
    
      -l  List saved authentications.    -u  Log out.
      -s  Save current authentication.
    
    Environment variables: UAUTH_FAKE_DIR

uauth

    Usage: uauth [-f][-p]
    
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
    

## Collaborating

For making bug reports, feature requests and donations visit
one of the following links:

1. [gemini://harkadev.com/oss/](gemini://harkadev.com/oss/)
2. [https://harkadev.com/oss/](https://harkadev.com/oss/)
