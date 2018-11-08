package tview

type rect struct {
	x, y, width, height int
}

type primitivesScreen struct {
	allPrimitives map[uint64]rect
	primitives    [][][]uint64
	width         int
	height        int
}

func (p *primitivesScreen) SetPrimitive(primitiveID uint64, x, y, width, height int) {

}

func (p *primitivesScreen) RemovePrimitive(primitiveID uint64) {

}

func (p *primitivesScreen) Reset() {

}

func (p *primitivesScreen) Resize(toWidth, toHeight int) {

}

func (p *primitivesScreen) GetSize() (width, height int) {
	return p.width, p.height
}

func (p *primitivesScreen) GetMatchingPrimitives(x, y int) []uint64 {
	return nil
}
