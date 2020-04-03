package entities

type Resource struct {
	index     int
	maxAmount int
	amount    int
}

func NewResource(index int) *Resource {
	return &Resource{
		index,
		20,
		0,
	}
}

func (r *Resource) add(amount int) {
	r.amount += amount

	if r.amount > r.maxAmount {
		r.amount = r.maxAmount
	}
}
