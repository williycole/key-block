// wordProcessingUtils/wordProcessing.go
package wordProcessingUtils

type WordForProcess struct {
	WordSlice []string
	Word      string
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
