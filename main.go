package main

import (
    "os"
    "path/filepath"
    "time"

    "github.com/alexraileanu/piper/pkg/aws"
    "github.com/alexraileanu/piper/pkg/cam"

    "github.com/go-co-op/gocron"
    _ "github.com/joho/godotenv/autoload"
)

func main() {
    println("hello hello")
    scheduler := gocron.NewScheduler(time.UTC)
    c := cam.Initialize()
    a := aws.Initialize()

    scheduler.Every(1).Minute().Do(func() {
        err := c.Snap()
        if err != nil {
            return
        }
        a.Post(c.F, os.Getenv("BUCKET_NAME"), filepath.Base(c.F))
        c.Clean()
    })

    <-scheduler.StartAsync()
}
