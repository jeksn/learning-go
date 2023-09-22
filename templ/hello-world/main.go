package main

import (
    "context"
    "os"
)

func main() {
    component := hello("Johan")
    component.Render(context.Background(), os.Stdout)
}
