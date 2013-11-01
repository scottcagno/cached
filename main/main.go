package main 

import (
    "cached"
    "runtime"
)

func init() {
    runtime.GOMAXPROCS(1)
}

func main() {
    cached.InitStore().ListenAndServe(":31337")
}
