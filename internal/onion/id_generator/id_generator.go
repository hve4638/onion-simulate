package idgenerator

type IdGenerator struct {
	size   uint32
	ids    []uint32
	offset int
}

func (generator *IdGenerator) Next() uint32 {
	if !generator.HasNext() {
		panic("IdGenerator: out of range")
	}

	id := generator.ids[generator.offset]
	generator.offset++
	return id
}

func (generator *IdGenerator) HasNext() bool {
	return generator.offset < int(generator.size)
}
