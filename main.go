package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/getlantern/systray"
)

func main() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(getIcon("sennuki.ico"))
	systray.SetTitle("Opener")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()

	os.Stderr, _ = os.Create("stderr.txt")

    o, err := NewOpener(os.Stderr)
	if err != nil {
		log.Fatal(err)
		systray.Quit()
	}

	if err = o.Run(); err != nil {
		log.Fatal(err)
		systray.Quit()
	}
}

func onExit() {
	// clean up here
}

func getIcon(s string) []byte {
    b, err := ioutil.ReadFile(s)
    if err != nil {
        log.Fatal(err)
    }
    return b
}
