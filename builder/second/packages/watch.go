package packages

// WatchBuilder struct
type WatchBuilder struct {
	pc System
}

// SetRAM method for WatchBuilder struct
func (b *WatchBuilder) SetRAM(ram int) Devicebuilder {
	b.pc.RAMCap = ram
	return b
}

// SetHDD method for WatchBuilder struct
func (b *WatchBuilder) SetHDD(hard int) Devicebuilder {
	b.pc.HDDCap = hard
	return b
}

// SetCPU method for WatchBuilder struct
func (b *WatchBuilder) SetCPU(cpu string) Devicebuilder {
	b.pc.CPU = cpu
	return b
}

// SetOS method for WatchBuilder struct
func (b *WatchBuilder) SetOS(os string) Devicebuilder {
	b.pc.OS = os
	return b
}

// SetGPU method for WatchBuilder struct
func (b *WatchBuilder) SetGPU(gpu string) Devicebuilder {
	b.pc.GPU = gpu
	return b
}

// GetSystem method for WatchBuilder struct
func (b *WatchBuilder) GetSystem() System {
	return b.pc
}
