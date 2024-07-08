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

			pressedKey := GetPressedKey(keyEvent)
			// Handle known key events
			handleKeyEvents(pressedKey, wordSlice)

		}
	}
}

func handleKeyEvents(pressedKey string, wordSlice []string) {
	// rethink this empty check, is there a better way?
	if pressedKey != "" {
		// TODO or method
		fmt.Printf("Received Key Down For: %v\n", pressedKey)

		//if we have a word, & if Space or Enter is pressed assume we have a Word in need to checking
		if pressedKey != "" && (pressedKey == "SPACE" || pressedKey == "RETURN") {

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

// todo test this first
func GetPressedKey(keyEvent types.KeyboardEvent) string {
	vkCode := strings.Split(keyEvent.VKCode.String(), "_")

	var keyEventType = getKeyEventType(keyEvent)
	// if not key down, or not key we don't care so return empty and handle later
	// rethink the empy return, is there a better way
	if keyEventType != "KeyDown" || len(vkCode) <= 0 {
		return ""
	}

	// so we have a char for a key down event, format it uppercase and return it
	var lastChar = strings.ToUpper(string(vkCode[len(vkCode)-1]))
	fmt.Printf("vkCode: %v,\n keyEvent: %v,\n lastChar: %v\n", vkCode, keyEvent, lastChar)
	return lastChar
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

// todo make this as fast as possible, routines, slices, pointers for less mem, somethign like that
// mocked up for now
func checkBlockedWordsForCurrentWord(word string) {
	for _, blockedWord := range BlockedWords {

		// SEEME checking key events as strings, so already capitalized, need to get perfectly right
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
