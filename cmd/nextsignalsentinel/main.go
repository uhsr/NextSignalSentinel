// cmd/nextsignalsentinel/main.go
package main

import (
"flag"
"log"
"os"

"nextsignalsentinel/internal/nextsignalsentinel"
)

func main() {
verbose := flag.Bool("verbose", false, "Enable verbose logging")
flag.Parse()

app := nextsignalsentinel.NewApp(*verbose)
if err := app.Run(); err != nil {
log.Fatal(err)
}
}
