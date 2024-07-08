package keyEventMappingUtils

import (
	"github.com/moutend/go-hook/pkg/types"
	// "os"
	"testing"
	// "time"
)

func TestGetPressedKey(t *testing.T) {

	// func getPressedKey(keyEvent types.KeyboardEvent) string {
	// {WM_KEYDOWN,{VK_C, 46, 0,744626625,0}}
	//
	// type KeyboardEvent struct {
	// Message Message
	// KBDLLHOOKSTRUCT
	// }
	// TODO - for wednesday, finish mocking this out, finish the happy path, and make one for not happy path
	var keyDown uintptr = 0x0100
	message := types.Message(keyDown)

	// Create a mock keyboard event
	mockKeyboardEvent := types.KeyboardEvent{
		Message: message,
	}

	got := GetPressedKey(mockKeyboardEvent)
	want := "C"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// gcc to comment and uncomment stuff yo
// TODO - finish test currently broken
// func TestTrackPressedKeyes(t *testing.T) {
//
// 	// Create channels to simulate input
// 	osSignalChan := make(chan os.Signal, 1)
// 	keyboardChan := make(chan types.KeyboardEvent, 1)
//
// 	// Create a goroutine to call TrackPressedKeys
// 	done := make(chan error)
// 	go func() {
// 		err := TrackPressedKeys(osSignalChan, keyboardChan)
// 		done <- err
// 	}()
//
// 	// type KeyboardEvent struct {
// 	// Message Message
// 	// KBDLLHOOKSTRUCT
// 	// }
// 	var keyDown uintptr = 0x0100
// 	message := types.Message(keyDown)
//
// 	// Create a mock keyboard event
// 	mockKeyboardEvent := types.KeyboardEvent{
// 		Message: message,
// 	}
//
// 	// Simulate sending a keyboard event
// 	keyboardChan <- mockKeyboardEvent
//
// 	// Simulate sending an OS signal (optional, depending on your method implementation)
// 	osSignalChan <- os.Interrupt
//
// 	// Close channels (optional, depending on your method implementation)
// 	close(keyboardChan)
// 	close(osSignalChan)
//
// 	// Wait for the goroutine to finish
// 	select {
// 	case err := <-done:
// 		if err != nil {
// 			t.Errorf("TrackPressedKeys returned an error: %v", err)
// 		}
// 	case <-time.After(1 * time.Second):
// 		t.Error("TrackPressedKeys timed out")
// 	}
//
// }
