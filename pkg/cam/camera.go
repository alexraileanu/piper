package cam

import (
    "fmt"
    "os"
    "os/exec"
    "time"
)

type Cam struct {
    F string
}

func Initialize() *Cam {
    return &Cam{}
}

func (c *Cam) Snap() error {
    fileName := fmt.Sprintf("%s.jpg", time.Now().Format("20060102150405"))
    filePath := fmt.Sprintf("/tmp/%s", fileName)
    c.F = filePath

    cmd := exec.Command(
        "raspistill",
        "-o",
        filePath,
        "-q",
        "75",
        "-w",
        "1280",
        "-h",
        "720",
    )
    err := cmd.Start()
    if err != nil {
        return err
    }

    return cmd.Wait()
}

func (c *Cam) Clean() error {
    return os.Remove(c.F)
}
