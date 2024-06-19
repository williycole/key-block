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

// TODO
func buildWordFromKeyDownMap(pressedKey string) string {
	panic("method not implemented")
}

// TODO
func detectNewWordBySpaceKeyDown(pressedKey string) string {
	panic("method not implemented")
}

// TODO - probably let this be a map for prototyping but will be a db or cache in mem
// TODO - this will be done with go routines for checking a lot of words super fast
func checkBlockedWordsForCurrentWord(pressedKey string) string {
	panic("method not implemented")
}

// TODO - here for poc lets just log this for now, eventually we want to use this event to kick off a lot of things
func handleLogBlockedWordEvenDetected(pressedKey string) string {
	panic("method not implemented")
}
