package packages

// System struct Product to build
type System struct {
	RAMCap int    `json:"ram"`
	HDDCap int    `json:"hard"`
	CPU    string `json:"cpu"`
	OS     string `json:"os"`
	GPU    string `json:"gpu"`
}

// Builder struct for build a device
type Builder struct {
	b Devicebuilder
}

// Devicebuilder interface for build new device
type Devicebuilder interface {
	SetRAM(int) Devicebuilder
	SetHDD(int) Devicebuilder
	SetCPU(string) Devicebuilder
	SetOS(string) Devicebuilder
	SetGPU(string) Devicebuilder
	GetSystem() System
}

// SetBuilder method
func (m *Builder) SetBuilder(builder Devicebuilder) {
	m.b = builder
}

// Build method
func (m *Builder) Build() Devicebuilder {
	return m.b
}
