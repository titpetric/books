package main

import (
	"errors"
	"os"
	"time"

	"github.com/apex/log"
	"github.com/apex/log/handlers/multi"
	"github.com/apex/log/handlers/papertrail"
	"github.com/apex/log/handlers/text"
)

func main() {
	handlerText := text.New(os.Stdout)
	handlerPapertrail := papertrail.New(&papertrail.Config{
		Host:     "logs5",
		Port:     41062,
		Hostname: "app",
		Tag:      "production",
	})
	handler := multi.New(handlerPapertrail, handlerText)

	log.SetHandler(handler)

	ctx := log.WithFields(log.Fields{
		"file": "something.png",
		"type": "image/png",
		"user": "tobi",
	})

	for range time.Tick(time.Second * 2) {
		ctx.Info("upload")
		ctx.Info("upload complete")
		ctx.Warn("upload retry")
		ctx.WithError(errors.New("unauthorized")).Error("upload failed")
		ctx.Errorf("failed to upload %s", "img.png")
	}
}
