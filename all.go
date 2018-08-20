package owparser

// All stats together in one struct
type All struct {
	*General
	QuickPlay   *Stats `json:"quickplay,omitempty"`
	Competitive *Stats `json:"competitive,omitempty"`
}

// NewAll Get all stats together
func (c *CareerProfile) NewAll() *All {
	all := new(All)
	all.General = c.NewGeneral()
	all.QuickPlay = c.NewStats(QuickPlay)
	all.Competitive = c.NewStats(Competitive)

	return all
}
