package keyEventMappingUtils

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/moutend/go-hook/pkg/types"
)

func TrackPressedKeys(osSignalChan <-chan os.Signal, keyboardChan <-chan types.KeyboardEvent) error {
	fmt.Println("...capturing keyboard input...")
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

				//if we have a word, & if Space or Enter is pressed assume we have a Word in need to checking
				if pressedKey != "" && (pressedKey == "Space" || pressedKey == "Enter") {
					// Join wordSlice into a complete word
					word := strings.Join(wordSlice, "")
					fmt.Printf("Complete Word: %s\n", word)

					go checkBlockedWordsForCurrentWord(word)
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

// todo make this as fast as possible, routines, slices, pointers for less mem, somethign like that
// mocked up for no
func checkBlockedWordsForCurrentWord(word string) {
	for _, blockedWord := range BlockedWords {

		if blockedWord == word {
			handleBlockedWordEvent(word)
		}
	}
}

// todo finish
func handleBlockedWordEvent(word string) {
	fmt.Printf("Blocked Word Found: %v, is not allowd\n", word)
	fmt.Printf("Kicking off blocked word events\n")
}

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
