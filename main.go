package main

import (
    "log"
    "os"
    "path/filepath"
    "time"

    "github.com/alexraileanu/piper/pkg/aws"
    "github.com/alexraileanu/piper/pkg/cam"

    "github.com/go-co-op/gocron"
    _ "github.com/joho/godotenv/autoload"
)

func main() {
    scheduler := gocron.NewScheduler(time.Local)
    c := cam.Initialize()
    a := aws.Initialize()

    scheduler.Every(1).Minute().Do(func() {
        println("attempting to snap pic")
        err := c.Snap()
        if err != nil {
            log.Fatalf("error snapping: %v\n", err)
            return
        }
        a.Post(c.F, os.Getenv("BUCKET_NAME"), filepath.Base(c.F))
        err = c.Clean()
        if err != nil {
            log.Fatalf("error cleaning file: %v\n", err)
            return
        }
        println("finished snapping and sending pic :)")
    })

    <-scheduler.StartAsync()
}
