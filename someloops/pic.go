package someloops

// Pic to be called from "golang.org/x/tour/pic.Show()"
func Pic(dx, dy int) [][]uint8 {
	outer := make([][]uint8, dy)
	for y := 0; y < dy; y++ {
		inner := make([]uint8, dx)
		for x := 0; x < dx; x++ {
			inner[x] = uint8(x ^ y) // this defines the output picture
		}
		outer[y] = inner
	}
	return outer
}
