package main

import (
    "fmt"
    "os"
    "path/filepath"
    "time"

    "github.com/alexraileanu/piper/pkg/aws"
    "github.com/alexraileanu/piper/pkg/cam"

    "github.com/go-co-op/gocron"
    _ "github.com/joho/godotenv/autoload"
)

func main() {
    scheduler := gocron.NewScheduler(time.UTC)
    c := cam.Initialize()
    a := aws.Initialize()

    scheduler.Every(1).Minute().Do(func() {
        fmt.Printf("attempting to snap pic\n")
        err := c.Snap()
        if err != nil {
            fmt.Printf("error snapping: %v\n", err)
            return
        }
        a.Post(c.F, os.Getenv("BUCKET_NAME"), filepath.Base(c.F))
        err = c.Clean()
        if err != nil {
            fmt.Printf("error cleaning file: %v\n", err)
            return
        }
    })

    <-scheduler.StartAsync()
}
