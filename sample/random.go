package sample

import (
	"math/rand"
	"time"

	"github.com/Serj1c/pcbook/pb"
	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	case 3:
		return pb.Keyboard_AZERTY
	default:
		return pb.Keyboard_QWERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

// hardcoded :(
func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return "not enough arguments"
	}
	return a[rand.Intn(n)]
}

// hardcoded :(
func randomCPUName(b string) string {
	if b == "Intel" {
		return randomStringFromSet(
			"XEON",
			"Core i5",
			"Core i7",
			"Core i3",
		)
	}
	return randomStringFromSet(
		"Ryzen 3",
		"Ryzen 5",
		"Ryzen 7",
	)
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat64(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(b string) string {
	switch b {
	case "Apple":
		return randomStringFromSet("Macbook Air", "Macbook Pro")
	case "Dell":
		return randomStringFromSet("Vostro", "XPS", "Alienware")
	default:
		return randomStringFromSet("Thinkpad X3", "Thinkpad P1", "Thinkpad P53")
	}
}

func randomGPUName(b string) string {
	if b == "NVIDIA" {
		return randomStringFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1360",
			"GTX 1070",
		)
	}
	return randomStringFromSet(
		"RX 590",
		"RX 580",
		"RX 5070",
		"RX Vega",
	)
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomInt(1080, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{
		Height: uint32(height),
		Width:  uint32(width),
	}
	return resolution
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomID() string {
	return uuid.New().String()
}
