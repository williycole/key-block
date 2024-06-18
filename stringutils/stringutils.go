package stringutils

import (
	"fmt"
	"os"
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
		case k := <-keyboardChan:
			// filter out dupes by only checking key down events
			var keyEventType = getKeyEventType(k)
			if keyEventType == "KeyDown" {
				// fmt.Println("KeyDown yyyyyoooo")
				keyName := VirtualKeyMap[uint16(k.VKCode)]
				if keyName == "" {
					keyName = fmt.Sprintf("Unknown (0x%X)", k.VKCode)
				}
				fmt.Printf("Received %v %v\n", k.Message, keyName)
				fmt.Printf("Received %v\n", keyName)
			}
			continue
		}
	}
}

// TODO
// func convertVkCodeToString(pressedKey string) string {
// 	// 	fmt.Println("here I am")
// 	return pressedKey
// }

// TODO
// func trimKeyDownString(pressedKey string) string {
// 	// 	fmt.Println("here I am")
// 	return pressedKey
// }

// TODO - or hanlde the fact that it recieves up down and we only need one
func getOnlyKeyDown(pressedKey string) string {
	// 	fmt.Println("here I am")
	return pressedKey
}

// TODO
func detectNewWordBySpaceKeyDown(pressedKey string) string {
	// 	fmt.Println("here I am")
	return pressedKey
}

// TODO
func buildWordFromKeyDownMap(pressedKey string) string {
	// 	fmt.Println("here I am")
	return pressedKey
}

// TODO - probably let this be a map for prototyping but will be a db or cache in mem
// TODO - this will be done with go routines for checking a lot of words super fast
func checkBlockedWordsForCurrentWord(pressedKey string) string {
	return pressedKey
}

// TODO - here for poc lets just log this for now, eventually we want to use this event to kick off a lot of things
func handleLogBlockedWordEvenDetected(pressedKey string) string {
	// 	fmt.Println("here I am")
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
