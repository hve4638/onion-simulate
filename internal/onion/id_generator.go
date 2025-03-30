package onion

import "math/rand"

type IdGenerator struct {
	size   uint32
	ids    []uint32
	offset int
}

func MakeIdCounter(size uint32, seed int64) IdGenerator {
	ids := make([]uint32, size)

	for i := range ids {
		ids[i] = uint32(i)
	}
	r := rand.New(rand.NewSource(seed))
	r.Shuffle(len(ids), func(i, j int) {
		ids[i], ids[j] = ids[j], ids[i]
	})

	return IdGenerator{
		size:   size,
		ids:    ids,
		offset: 0,
	}
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
