package idgenerator

import "math/rand"

func NewIdGenerator(size uint32, seed int64) IdGenerator {
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
