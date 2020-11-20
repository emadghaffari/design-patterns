package packages

// LaptopBuilder struct
type LaptopBuilder struct {
	pc System
}

// SetRAM method for LaptopBuilder struct
func (b *LaptopBuilder) SetRAM(ram int) Devicebuilder {
	b.pc.RAMCap = ram
	return b
}

// SetHDD method for LaptopBuilder struct
func (b *LaptopBuilder) SetHDD(hard int) Devicebuilder {
	b.pc.HDDCap = hard
	return b
}

// SetCPU method for LaptopBuilder struct
func (b *LaptopBuilder) SetCPU(cpu string) Devicebuilder {
	b.pc.CPU = cpu
	return b
}

// SetOS method for LaptopBuilder struct
func (b *LaptopBuilder) SetOS(os string) Devicebuilder {
	b.pc.OS = os
	return b
}

// SetGPU method for LaptopBuilder struct
func (b *LaptopBuilder) SetGPU(gpu string) Devicebuilder {
	b.pc.GPU = gpu
	return b
}

// GetSystem method for LaptopBuilder struct
func (b *LaptopBuilder) GetSystem() System {
	return b.pc
}
