package tui

import (
	"sync"

	"github.com/charmbracelet/bubbles/list"
)

type ItemGenerator struct {
	titles     []string
	descs      []string
	titleIndex int
	descIndex  int
	mtx        *sync.Mutex
}

func (r *ItemGenerator) reset() {
	r.mtx = &sync.Mutex{}

}

func (r *ItemGenerator) next() item {
	if r.mtx == nil {
		r.reset()
	}

	r.mtx.Lock()
	defer r.mtx.Unlock()

	i := item{
		title:   r.titles[r.titleIndex],
		command: r.descs[r.descIndex],
	}

	r.titleIndex++
	if r.titleIndex >= len(r.titles) {
		r.titleIndex = 0
	}

	r.descIndex++
	if r.descIndex >= len(r.descs) {
		r.descIndex = 0
	}

	return i
}

func (r *ItemGenerator) read(dic map[string]string) []list.Item {
	var items []list.Item
	for k, v := range dic {
		items = append(items, item{
			title:   k,
			command: v,
		})
	}
	return items
}
