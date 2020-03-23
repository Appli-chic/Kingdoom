package screens

type Screen interface {
	HandleEvents() bool
	Update()
	Render()
}
