package bag

type StringBag interface {
	Add(string)
	Size() int
	IsEmpty() bool
}

func NewStringBag() StringBag {
	return &stringBag{}
}

type stringBag struct {
	StringBag

	bag []string
}

func (b *stringBag) Add(data string) {
	b.bag = append(b.bag, data)
}

func (b *stringBag) Size() int {
	return len(b.bag)
}

func (b *stringBag) IsEmpty() bool {
	return b.Size() == 0
}
