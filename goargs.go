
package main

import (
    "fmt"
    "log"
    "os"
    "flag"
)

func main() {

    log.Println("Setting Defaults...")
    // set sain default for this option
    serviceUrl := "http://netjibbing.com/api"
    
    log.Println("Overriding defaults with environment variables...")
    // override default from environment variable if it exists
    if os.Getenv("SERVICE_URL") != "" {
      serviceUrl = os.Getenv("SERVICE_URL")
    }
 
    log.Println("Overriding current options with flags...")

    // override default and env variable with option flag from command line
    // flag.StringVar( pointer for the variable, variable name, default value, description)
    flag.StringVar(&serviceUrl, "serviceUrl", serviceUrl, "service url")

    // Once all flags are declared, call `flag.Parse()`
    // to execute the command-line parsing.
    flag.Parse()

    // Print the options out for inspection
    fmt.Println("The Serivce URL is: ",serviceUrl)
}

