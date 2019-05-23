# ouiparser

IEEE OUI data parser

You can download `oui.txt` here: standards.ieee.org/develop/regauth/oui/oui.txt

## Install

`go get github.com/dfkdream/ouiparser`

## Example Code
```Go
package main

import (
    "log"
    "fmt"
    "net"
    "github.com/dfkdream/ouiparser"
)

func main(){
    oui,err:=ouiparser.ParseOUI("oui.txt")
    if err!=nil{
        log.Fatal(err)
    }

    mac,err:=net.ParseMAC("d4:38:9c:00:00:00")
    if err!=nil{
        log.Fatal(err)
    }

    result:=ouiparser.SearchOUI(oui,mac)
    fmt.Println(result.Organization)
}
```
### Output
```
Sony Mobile Communications Inc
```