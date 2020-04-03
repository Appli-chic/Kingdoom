package entities

type Inventory struct {
	resources map[int]*Resource
}

func NewInventory() *Inventory {
	return &Inventory{
		make(map[int]*Resource),
	}
}

func (i *Inventory) addResource(index int, amount int) {
	if _, ok := i.resources[index]; !ok {
		// Create the resource in the map
		i.resources[index] = NewResource(index)
		i.resources[index].add(amount)
	} else {
		// Add amount to the existing resource in the inventory
		i.resources[index].add(amount)
	}
}
