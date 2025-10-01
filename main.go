package main

import (
    "fmt"
    "os"
    "os/user"
    "github.com/HemanthBangera/PicoLang/repl"
)

func main() {
    user, err := user.Current()
    if err != nil {
        panic(err)
    }
    fmt.Printf("Hello %s! This is Charles Oliveira! \n Happy Birthday to you ! Lets go !\n",user.Usernam:)
    fmt.Printf("Feel free to type in commands\n")
    repl.Start(os.Stdin,os.Stdout)
}
