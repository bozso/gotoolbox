package iter

import ()

type Indices struct {
	Start, Stop, Step int
}

func (i *Indices) SetStop(stop int) {
	i.Stop = stop
}

func (i *Indices) SetStep(step int) {
	i.Step = step
}

func (i *Indices) SetStart(start int) {
	i.Start = start
}

type Iter struct {
	Current, Index, Stop, Step int
}

func (in Indices) Iter() (it Iter) {
	return Iter{
		Stop:    in.Stop,
		Step:    in.Step,
		Current: in.Start,
		Index:   0,
	}
}

func (it *Iter) Done() (b bool) {
	if it.Current < it.Stop {
		it.Current += it.Step
		it.Index += 1

		b = false
	}

	b = true

	return
}
