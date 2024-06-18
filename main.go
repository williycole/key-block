package main

// TODO cleanup: consider this when thinking about external dependencies as apart of final cleanup https://chatgpt.com/share/95409b0a-d5fa-4fa9-89c4-61fbd00d4617
import (
	"fmt"
	"log"
	"os"
	"os/signal"

	// "syscall"
	// "unsafe"

	"github.com/moutend/go-hook/pkg/keyboard"
	"github.com/moutend/go-hook/pkg/types"

	// "golang.org/x/sys/windows"
	"key-block/stringutils"
)

func main() {
	fmt.Println("Hello World")

	log.SetFlags(0)
	log.SetPrefix("error: ")

	if err := runKeyBlock(); err != nil {
		log.Fatal(err)
	}
}

// TODO look up more of this code to understand it / how go work
func runKeyBlock() error {
	// Buffer size depends on your need. The 100 is placeholder value.
	keyboardChan := make(chan types.KeyboardEvent, 100)

	if err := keyboard.Install(nil, keyboardChan); err != nil {
		return err
	}

	defer keyboard.Uninstall()

	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, os.Interrupt)

	var trackedKeyEvents = stringutils.TrackPressedKeys(osSignalChan, keyboardChan)

	return trackedKeyEvents
}
