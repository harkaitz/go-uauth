# NAME

uauth - A Go library and small program for oauth debugging.

# SYNOPSIS

    uauth [-fp]
    
    import "github.com/harkaitz/go-uauth"

# DESCRIPTION

Authenticate users using Google Oauth. This library is intended for
web service creation.

The idea is to use the "User" type on the libraries you create, and
to use the "CLI*" family of functions and the "uauth" command line
utility for testing.

# CONFIGURATION FILE

To configure the library use "uauth.LoadSettings". It expects the
following arguments.

    {
        "UseHTTPS": false,          // Required false for the CLI.
        "Domain": "127.0.0.1:8080", // Required "127.0.0.1:8080" for the CLI.
        "GoogleClientID": "...",
        "GoogleClientSecret": "...",
        "GoogleCallbackPage": "/oauthcb/google",
        "RandomString": "<random-string"
    }

# COMMAND LINE UTILITY (DO NOT USE IN PRODUCTION)

It binds to "127.0.0.1:8080" and opens a web browser for authentication.

This of course requires the programmer to configure the oauth service
to allow "127.0.0.1:8080" domain.

The configuration is readen from "~/.config.testing.json". The info of
the logged user is saved in "~/.uauth.json".

With "-f" the authentication happens always. With "-p" the resulting
JSON is printed to the standard output.

# SEE ALSO

**JQ(1)**
