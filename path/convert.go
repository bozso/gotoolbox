package path

import ()

type IndexedFrom interface {
	GetFrom(ii int) From
}

type From interface {
	FromPath(p Pather) error
}

type Transformer interface {
	Transform(Like) (Like, error)
}

type Valids []Valid

func (v Valids) Filter(filt Filterer) (filtered Valids, err error) {
	filtered = make(Valids, 0, len(v))

	for _, p := range v {
		keep, err := filt.Filter(p)
		if err != nil {
			break
		}
		if keep {
			filtered = append(filtered, p)
		}
	}
	return
}

func (v Valids) Transform(trans Transformer) (transformed Valids, err error) {
	return

}
