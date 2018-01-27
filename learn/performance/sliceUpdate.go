package perf

type counter struct{ a, b, c, d int }

func Update(val int) {
	ct := make([]counter, 1000)
	for i := range ct {
		ct[i].a += val
		ct[i].b += val
		ct[i].c += val
		ct[i].d += val
	}
}

func UpdateRef(val int) {
	ct := make([]counter, 1000)
	for i := range ct {
		v := &ct[i]
		v.a += val
		v.b += val
		v.c += val
		v.d += val
	}
}

func UpdateRef1(val int) {
	ct := make([]counter, 1000)
	for i := range ct {
		ct[i].update(val)

	}
}

func (c *counter) update(v int) {
	c.a += v
	c.b += v
	c.c += v
	c.d += v
}
