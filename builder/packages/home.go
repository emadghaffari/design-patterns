package packages

// HomeEditionPCBuilder struct
type HomeEditionPCBuilder struct {
	pc PersonalComputer
}

// SetRAM method for HomeEditionPCBuilder struct
func (b *HomeEditionPCBuilder) SetRAM() PCBuilder {
	b.pc.RAMCap = 4
	return b
}

// SetHDD method for HomeEditionPCBuilder struct
func (b *HomeEditionPCBuilder) SetHDD() PCBuilder {
	b.pc.HDDCap = 500
	return b
}

// SetCPU method for HomeEditionPCBuilder struct
func (b *HomeEditionPCBuilder) SetCPU() PCBuilder {
	b.pc.CPU = "i3"
	return b
}

// SetOS method for HomeEditionPCBuilder struct
func (b *HomeEditionPCBuilder) SetOS() PCBuilder {
	b.pc.OS = "Windows 7 Home Edition"
	return b
}

// SetGPU method for HomeEditionPCBuilder struct
func (b *HomeEditionPCBuilder) SetGPU() PCBuilder {
	b.pc.GPU = "Intel Graphic Card"
	return b
}

// GetPC method for HomeEditionPCBuilder struct
func (b *HomeEditionPCBuilder) GetPC() PersonalComputer {
	return b.pc
}
