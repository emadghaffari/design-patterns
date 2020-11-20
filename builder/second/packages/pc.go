package packages

// PCBuilder struct
type PCBuilder struct {
	pc System
}

// SetRAM method for PCBuilder struct
func (b *PCBuilder) SetRAM(ram int) Devicebuilder {
	b.pc.RAMCap = ram
	return b
}

// SetHDD method for PCBuilder struct
func (b *PCBuilder) SetHDD(hard int) Devicebuilder {
	b.pc.HDDCap = hard
	return b
}

// SetCPU method for PCBuilder struct
func (b *PCBuilder) SetCPU(cpu string) Devicebuilder {
	b.pc.CPU = cpu
	return b
}

// SetOS method for PCBuilder struct
func (b *PCBuilder) SetOS(os string) Devicebuilder {
	b.pc.OS = os
	return b
}

// SetGPU method for PCBuilder struct
func (b *PCBuilder) SetGPU(gpu string) Devicebuilder {
	b.pc.GPU = gpu
	return b
}

// GetSystem method for PCBuilder struct
func (b *PCBuilder) GetSystem() System {
	return b.pc
}
