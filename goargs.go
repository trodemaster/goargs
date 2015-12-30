
package main

import (
    "fmt"
    //"log"
    "os"
    "flag"
)

func main() {

    // Basic flag declarations are available for string,
    // integer, and boolean options. Here we declare a
    // string flag `word` with a default value `"foo"`
    // and a short description. This `flag.String` function
    // returns a string pointer (not a string value);
    // we'll see how to use this pointer below.
    wordPtr := flag.String("word", "foo", "a string")

    // This declares `numb` and `fork` flags, using a
    // similar approach to the `word` flag.
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

    // It's also possible to declare an option that uses an
    // existing var declared elsewhere in the program.
    // Note that we need to pass in a pointer to the flag
    // declaration function.
    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

//log.Println("Overriding configuration from env vars.")
    // set sain default for this option
    serviceUrl := "http://netjibbing.com/api"
    // override default from environment variable if it exists
    fmt.Println(os.Getenv("SERVICE_URL"))
    if os.Getenv("SERVICE_URL") != "" {
      serviceUrl = os.Getenv("SERVICE_URL")
    }
    // override default and env variable with option flag from command line
    // flag.StringVar( pointer for the variable, variable name, default value, description)
    flag.StringVar(&serviceUrl, "serviceUrl", serviceUrl, "service url")

    // Once all flags are declared, call `flag.Parse()`
    // to execute the command-line parsing.
    flag.Parse()

    // Here we'll just dump out the parsed options and
    // any trailing positional arguments. Note that we
    // need to dereference the pointers with e.g. `*wordPtr`
    // to get the actual option values.
    

    fmt.Println(serviceUrl)

    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("fork:", *boolPtr)
    fmt.Println("svar:", svar)
    fmt.Println("tail:", flag.Args())
}

