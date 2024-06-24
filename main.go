package main

// TODO cleanup: Add a thoughtful folder structure to code chaning method names as appropriate
// TODO cleanup: Add idomatic comments to all methods
// TODO cleanup: consider this when thinking about external dependencies as apart of final
// TODO cleanup: consider error handling approach like anthony gg
// TODO prev --> this may not matter though so look into binarys
// TODO prev --> https://chatgpt.com/share/95409b0a-d5fa-4fa9-89c4-61fbd00d4617

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
	"key-block/keyEventMappingUtils"
	// "key-block/wordProcessingUtils"
)

func main() {
	fmt.Println("Hello from main")

	log.SetFlags(0)
	log.SetPrefix("error: ")

	if err := startup(); err != nil {
		log.Fatal(err)
	}
}

func startup() error {
	fmt.Println("Starting up key-block application...")

	log.SetFlags(0)
	log.SetPrefix("error: ")

	keyboardChan := make(chan types.KeyboardEvent, 100)
	if err := keyboard.Install(nil, keyboardChan); err != nil {
		return err
	}
	defer keyboard.Uninstall()

	osSignalChan := make(chan os.Signal, 1)
	signal.Notify(osSignalChan, os.Interrupt)

	// Start tracking key events in a separate goroutine
	trackedKeyEvents := make(chan error)
	go func() {
		trackedKeyEvents <- keyEventMappingUtils.TrackPressedKeys(osSignalChan, keyboardChan)
	}()

	// concurrently doing other tasks
	var word string
	for {
		select {
		case <-osSignalChan:
			fmt.Println("Received shutdown signal...")
			return nil
		case err := <-trackedKeyEvents:
			if err != nil {
				return err
			}
			// Perform other tasks here based on state or collected data
			fmt.Printf("Current word: %s\n", word)
			// Reset word for next iteration if needed
			word = ""
		}
	}
}
