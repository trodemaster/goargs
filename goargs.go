// example code for providing tool configuration from enviornemnet variables and option flags
package main

import (
    "fmt"
    "log"
    "os"
    "flag"
)

func main() {

    log.Println("Setting Defaults...")
    // set sane default for this option
    serviceUrl := "netjibbing.com/api"
    serviceUser := "testuser"
    servicePass := "testpass"
    serviceOption1 := "download"
    serviceOption2 := "2938472934"

    log.Println("Overriding defaults with environment variables...")
    // override default from environment variable if they exist
    if os.Getenv("SERVICE_URL") != "" {
      serviceUrl = os.Getenv("SERVICE_URL")
    }
    if os.Getenv("SERVICE_USER") != "" {
      serviceUser = os.Getenv("SERVICE_USER")
    }
    if os.Getenv("SERVICE_PASS") != "" {
      servicePass = os.Getenv("SERVICE_PASS")
    }
    if os.Getenv("SERVICE_OPTION_1") != "" {
      serviceOption1 = os.Getenv("SERVICE_OPTION_1")
    }
    if os.Getenv("SERVICE_OPTION_2") != "" {
      serviceOption2 = os.Getenv("SERVICE_OPTION_2")
    }
 
    log.Println("Overriding current options with flags...")

    // override default and env variable with option flag from command line
    // flag.StringVar( pointer for the variable, variable name, default value, description)
    flag.StringVar(&serviceUrl, "url", serviceUrl, "service url, env var SERVICE_URL")
    flag.StringVar(&serviceUser, "user", serviceUser, "service username, env var SERVICE_USER")
    flag.StringVar(&servicePass, "pass", servicePass, "service password, env var SERVICE_PASS")
    flag.StringVar(&serviceOption1, "option1", serviceOption1, "Service Option 1, env var SERVICE_OPTION_1")
    flag.StringVar(&serviceOption2, "option2", serviceOption2, "Service Option 2, env var SERVICE_OPTION_2")

    // Once all flags are declared, call `flag.Parse()`
    // to execute the command-line parsing.
    flag.Parse()

    // Print the options out for inspection
    log.Println("Outputting current options")

    fmt.Println("The Serivce URL is: ",serviceUrl)
    fmt.Println("The Serivce username is: ",serviceUser)
    fmt.Println("The Serivce password is: ",servicePass)
    fmt.Println("The Serivce option 1 is: ",serviceOption1)
    fmt.Println("The Serivce option 1 is: ",serviceOption2)

    // compose a url from the options
    log.Println("Constructing a full url...")
    fmt.Println("https://"+serviceUser+":"+servicePass+"@"+serviceUrl+"/"+serviceOption1+"/"+serviceOption2)

    // log end of run
    log.Println("completed...")

}

