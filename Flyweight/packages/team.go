package packages

// Team struct
type Team struct {
	ID             uint64
	Name           string
	shield         []byte
	Players        []Player
	HistoricalData []HistoricalData
}
