package rgb

import (
	"math"
)

// Temperature will run RGB function
func (r *ActiveRGB) Temperature(temperature float32) {
	buf := map[int][]byte{}
	color := getTemperatureColor(float64(temperature), 0, 100)
	modify := ModifyBrightness(color)
	for j := 0; j < r.LightChannels; j++ {
		buf[j] = []byte{
			byte(modify.Red),
			byte(modify.Green),
			byte(modify.Blue),
		}
		if r.ContainsPump && r.HasLCD {
			if j > 15 && j < 20 {
				buf[j] = []byte{0, 0, 0}
			}
		}
	}
	r.Output = SetColor(buf)
}

func lerp(start, end, t float64) float64 {
	return start + t*(end-start)
}

func getTemperatureColor(temp, minTemp, maxTemp float64) Color {
	// Normalize the temperature between 0 and 1
	normalized := (temp - minTemp) / (maxTemp - minTemp)
	normalized = math.Max(0, math.Min(1, normalized)) // Clamp value between 0 and 1

	var r, g, b float64

	if normalized <= 0.5 {
		// Transition from green (0,255,0) to orange (255,165,0)
		t := normalized / 0.5
		r = lerp(0, 255, t)   // Red increases from 0 to 255
		g = lerp(255, 165, t) // Green decreases from 255 to 165
		b = 0                 // Blue remains 0
	} else {
		// Transition from orange (255,165,0) to red (255,0,0)
		t := (normalized - 0.5) / 0.5
		r = 255             // Red stays at 255
		g = lerp(165, 0, t) // Green decreases from 165 to 0
		b = 0               // Blue remains 0
	}

	return Color{Red: r, Green: g, Blue: b, Brightness: 1}
}
