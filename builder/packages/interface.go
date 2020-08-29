package packages

// PersonalComputer struct Product to build
type PersonalComputer struct {
	RAMCap int    `json:"ram"`
	HDDCap int    `json:"hard"`
	CPU    string `json:"cpu"`
	OS     string `json:"os"`
	GPU    string `json:"gpu"`
}

//Manufacturer object which aware of build process for builder type
type Manufacturer struct {
	b PCBuilder
}

// PCBuilder interface Each builder should implement this interface
type PCBuilder interface {
	SetRAM() PCBuilder
	SetHDD() PCBuilder
	SetCPU() PCBuilder
	SetOS() PCBuilder
	SetGPU() PCBuilder
	GetPC() PersonalComputer
}

// SetBuilder method for Manufacturer
func (m *Manufacturer) SetBuilder(builder PCBuilder) {
	m.b = builder
}

// ConstructPC method for Manufacturer
func (m *Manufacturer) ConstructPC() {
	m.b.SetCPU().SetHDD().SetRAM().SetGPU().SetOS()
}
