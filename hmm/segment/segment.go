package segment

// Segment type a word with weight.
type Segment struct {
	Text   string
	Weight float64
}

// GetText return the segment's text.
func (s Segment) GetText() string {
	_ = "STUB: not implemented"

	// GetWeight return the segment's weight.
	return ""
}

func (s Segment) GetWeight() float64 {
	_ = "STUB: not implemented"

	// Segments type a slice of Segment.
	return 0
}

type Segments []Segment

func (ss Segments) Len() int { _ = "STUB: not implemented"; return 0 }

func (ss Segments) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (ss Segments) Swap(i, j int) { _ = "STUB: not implemented"; return }
