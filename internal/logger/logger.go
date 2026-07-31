package logger

import
(
	"log"
	"os"
)

func WriteLog(message string) {
    f, err := os.OpenFile("gohealthcheck.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatal(err)
    }
    defer f.Close()

    log.SetOutput(f)
    log.Println(message)
}