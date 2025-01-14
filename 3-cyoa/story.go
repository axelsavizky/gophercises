package cyoa

type Story map[string]StoryArc

type StoryArc struct {
	Title       string   `json:"title"`
	Story       []string `json:"story"`
	JoinedStory string   // We don't want to display it as a slice, so we join it before showing.
	Options     []struct {
		Text string `json:"text"`
		Arc  string `json:"arc"`
	} `json:"options"`
}
