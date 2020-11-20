package packages

// GameEditionPCBuilder struct edition pc builder
type GameEditionPCBuilder struct {
	pc PersonalComputer
}

// SetRAM method for GameEditionPCBuilder struct
func (b *GameEditionPCBuilder) SetRAM() PCBuilder {
	b.pc.RAMCap = 4
	return b
}

// SetHDD method for GameEditionPCBuilder struct
func (b *GameEditionPCBuilder) SetHDD() PCBuilder {
	b.pc.HDDCap = 500
	return b
}

// SetCPU method for GameEditionPCBuilder struct
func (b *GameEditionPCBuilder) SetCPU() PCBuilder {
	b.pc.CPU = "i3"
	return b
}

// SetOS method for GameEditionPCBuilder struct
func (b *GameEditionPCBuilder) SetOS() PCBuilder {
	b.pc.OS = "Windows 7 Home Edition"
	return b
}

// SetGPU method for GameEditionPCBuilder struct
func (b *GameEditionPCBuilder) SetGPU() PCBuilder {
	b.pc.GPU = "Intel Graphic Card"
	return b
}

// GetPC method for GameEditionPCBuilder struct
func (b *GameEditionPCBuilder) GetPC() PersonalComputer {
	return b.pc
}
