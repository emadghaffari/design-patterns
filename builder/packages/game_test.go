package packages

import "testing"

// TestGameSetRAM func
func TestGameSetRAM(t *testing.T) {
	home := GameEditionPCBuilder{}
	home.SetRAM()
	if home.pc.RAMCap != 4 {
		t.Error("The ram is wrong!")
	}
}

// TestGameSetHDD func
func TestGameSetHDD(t *testing.T) {
	home := GameEditionPCBuilder{}
	home.SetHDD()
	if home.pc.HDDCap != 500 {
		t.Error("The ram is wrong!")
	}
}

// TestGameSetCPU func
func TestGameSetCPU(t *testing.T) {
	home := GameEditionPCBuilder{}
	home.SetCPU()
	if home.pc.CPU != "i3" {
		t.Error("The ram is wrong!")
	}
}

// TestGameSetOS func
func TestGameSetOS(t *testing.T) {
	home := GameEditionPCBuilder{}
	home.SetOS()
	if home.pc.OS != "Windows 7 Home Edition" {
		t.Error("The ram is wrong!")
	}
}

// TestGameSetGPU func
func TestGameSetGPU(t *testing.T) {
	home := GameEditionPCBuilder{}
	home.SetGPU()
	if home.pc.GPU != "Intel Graphic Card" {
		t.Error("The ram is wrong!")
	}
}
