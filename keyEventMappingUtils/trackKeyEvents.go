package keyEventMappingUtils

import (
	"fmt"
	"os"
	"reflect"
	"time"

	"github.com/moutend/go-hook/pkg/types"
)

func TrackPressedKeys(osSignalChan <-chan os.Signal, keyboardChan <-chan types.KeyboardEvent) error {
	fmt.Println("start capturing keyboard input")
	for {
		select {
		case <-time.After(5 * time.Minute):
			fmt.Println("Received timeout signal")
			return nil
		case <-osSignalChan:
			fmt.Println("Received shutdown signal")
			return nil
		case keyEvent := <-keyboardChan:
			var pressedKeyAsString = convertKeyEventToString(keyEvent)

			// todo better comment
			// handle uknown key events, we only care about a few numbers, all letters, and space
			if pressedKeyAsString != "" {
				// log it and build the word
				fmt.Printf("Received Key Down For: %v %v %v\n", keyEvent.Message, pressedKeyAsString, reflect.TypeOf(pressedKeyAsString))

				var wordSlice = processPressedKeyStrings(pressedKeyAsString)
				if len(wordSlice) > 0 {
					fmt.Printf("New word found: %v\n", wordSlice)
				}
			}

			continue
		}
	}
}

// todo consider renaming pressedKey and simplifying
func convertKeyEventToString(keyEvent types.KeyboardEvent) string {
	pressedKey := VirtualKeyToStringMap[uint16(keyEvent.VKCode)]

	var keyEventType = getKeyEventType(keyEvent)
	if keyEventType != "KeyDown" {
		return ""
	}

	return pressedKey
}

func getKeyEventType(event types.KeyboardEvent) string {
	const (
		WM_KEYDOWN = 0x0100
		WM_KEYUP   = 0x0101
	)

	switch event.Message {
	case WM_KEYDOWN:
		return "KeyDown"
	case WM_KEYUP:
		return "KeyUp"
	default:
		return "Unknown"
	}
}

// todo consider swtich statmetn use
// todo fix this keep it simple, get it to work, write tests, then make it better
func processPressedKeyStrings(pressedKey string) []string {
	fmt.Printf("Entering: processingPressedKeyStrings...\n")
	// todo consider using length as capacity default, or handling a warning log for ui elsewhere
	var wordSlice []string = make([]string, 10, 20)
	// add to slice when not space
	if pressedKey != "Space" {
		wordSlice = append(wordSlice, pressedKey)
	}
	// if space return word
	if pressedKey == "Space" {

		fmt.Printf("Exiting bc Space:  processingPressedKeyStrings...\n")
		return wordSlice
	}
	// else return the empty word slice and check it elsewhere
	emptySlice := make([]string, 0)
	fmt.Printf("Exiting empty:  processingPressedKeyStrings...\n")
	return emptySlice
}
