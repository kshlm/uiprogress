package uiprogress

import (
	"github.com/kshlm/uiprogress/util/strutil"
	"sync"
)

type Banner struct {
	text  string
	width int
	mtx   *sync.RWMutex
}

func NewBanner(text string, width int) *Banner {
	return &Banner{
		strutil.Resize(text, uint(width)),
		width,
		&sync.RWMutex{},
	}
}

func (b *Banner) String() string {
	b.mtx.RLock()
	defer b.mtx.RUnlock()
	return b.text
}

func (b *Banner) Set(text string) {
	b.mtx.Lock()
	b.text = strutil.Resize(text, uint(b.width))
	b.mtx.Unlock()
}

func (b *Banner) SetWidth(width int) {
	b.mtx.Lock()
	b.width = width
	b.text = strutil.Resize(b.text, uint(b.width))
	b.mtx.Unlock()
}
