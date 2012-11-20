package main

import (
    "log"
    "flag"
    "github.com/mikespook/goemphp/php"
)

var (
    info = flag.Bool("i", false, "Display phpinfo")
    file = flag.String("f", "", "Execute the file")
    run = flag.String("r", "", "Execute the script")
)

func init() {
    flag.Parse()
}

func main() {
    p := php.New()
    p.Startup()
    defer p.Close()
    if *info {
        if err := p.Eval("phpinfo();"); err != nil {
            log.Fatal(err)
        }
        return
    }
    if *file != "" {
        if err := p.Exec(*file); err != nil {
            log.Fatal(err)
        }
        return
    }
    if *run != "" {
        if err := p.Eval(*run); err != nil {
            log.Fatal(err)
        }
        return
    }
    flag.Usage()
}
