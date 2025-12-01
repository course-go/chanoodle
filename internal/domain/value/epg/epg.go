package epg

// EPG represents an electronic program guide.
// It contains a list of channels sorted by priority
// with their respective program events.
type EPG struct {
	Channels []Channel
}
