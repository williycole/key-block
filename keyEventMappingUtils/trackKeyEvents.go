package keyEventMappingUtils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/moutend/go-hook/pkg/types"
)

// todo consider this being the main function that gets kicked off for each new word if needed
// func TrackPressedKeys(osSignalChan <-chan os.Signal, keyboardChan <-chan types.KeyboardEvent) error {
// 	fmt.Println("start capturing keyboard input")
// 	var wordSlice []string = make([]string, 10)
// 	for {
// 		select {
// 		case <-time.After(5 * time.Minute):
// 			fmt.Println("Received timeout signal")
// 			return nil
// 		case <-osSignalChan:
// 			fmt.Println("Received shutdown signal")
// 			return nil
// 		case keyEvent := <-keyboardChan:
// 			var pressedKeyAsString = convertKeyEventToString(keyEvent)

// 			// todo better comment
// 			// handle uknown key events, we only care about a few numbers, all letters, and space
// 			if pressedKeyAsString != "" {
// 				// log it and build the word
// 				fmt.Printf("Received Key Down For: %v %v %v\n", keyEvent.Message, pressedKeyAsString, reflect.TypeOf(pressedKeyAsString))
// 				wordSlice = append(wordSlice, pressedKeyAsString)
// 			}

// 			continue
// 		}
// 	}

// }

func TrackPressedKeys(osSignalChan <-chan os.Signal, keyboardChan <-chan types.KeyboardEvent) error {
	fmt.Println("Start capturing keyboard input")
	var wordSlice []string

	for {
		select {
		case <-time.After(5 * time.Minute):
			fmt.Println("Received timeout signal")
			return nil
		case <-osSignalChan:
			fmt.Println("Received shutdown signal")
			return nil
		case keyEvent := <-keyboardChan:
			pressedKey := convertKeyEventToString(keyEvent)

			// Handle known key events
			if pressedKey != "" {
				fmt.Printf("Received Key Down For: %v\n", pressedKey)

				// Check if the Space key is pressed to complete the word
				if pressedKey == "Space" || pressedKey == "Enter" {
					// Join wordSlice into a complete word
					word := strings.Join(wordSlice, "")
					fmt.Printf("Complete Word: %s\n", word)

					// Optionally do something with the complete word
					go goDoWorkWithWord(word)

					// Reset wordSlice for the next word
					wordSlice = nil
				} else {
					// Add pressed key to wordSlice
					wordSlice = append(wordSlice, pressedKey)
				}
			}
		}
	}
}

// todo mock
func goDoWorkWithWord(word string) {
	fmt.Printf("Full Word Is: %v\n", word)
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

func buildWordFromKeyPressStrings(keyPressString string) {
	panic("buildWordFromKeyPressStrings Not Implemented")
}
