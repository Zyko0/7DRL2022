package graphics

type Renderer struct {
	cameraYOffset float64
}

func NewRenderer() *Renderer {
	return &Renderer{}
}

func (r *Renderer) Update(playerY float64) {
	r.cameraYOffset = playerY
}
