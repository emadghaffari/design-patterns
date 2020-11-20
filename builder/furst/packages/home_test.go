package packages

import "testing"

// TestSetRAM func
func TestSetRAM(t *testing.T) {
	home := HomeEditionPCBuilder{}
	home.SetRAM()
	if home.pc.RAMCap != 4 {
		t.Error("The ram is wrong!")
	}
}

// TestSetHDD func
func TestSetHDD(t *testing.T) {
	home := HomeEditionPCBuilder{}
	home.SetHDD()
	if home.pc.HDDCap != 500 {
		t.Error("The ram is wrong!")
	}
}

// TestSetCPU func
func TestSetCPU(t *testing.T) {
	home := HomeEditionPCBuilder{}
	home.SetCPU()
	if home.pc.CPU != "i3" {
		t.Error("The ram is wrong!")
	}
}

// TestSetOS func
func TestSetOS(t *testing.T) {
	home := HomeEditionPCBuilder{}
	home.SetOS()
	if home.pc.OS != "Windows 7 Home Edition" {
		t.Error("The ram is wrong!")
	}
}

// TestSetGPU func
func TestSetGPU(t *testing.T) {
	home := HomeEditionPCBuilder{}
	home.SetGPU()
	if home.pc.GPU != "Intel Graphic Card" {
		t.Error("The ram is wrong!")
	}
}
