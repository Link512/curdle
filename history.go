package main

import "strings"

type history struct {
	chars      map[byte]struct{}
	historyStr strings.Builder
}

func newHistory() *history {
	return &history{
		chars: map[byte]struct{}{},
	}
}

func (h *history) Add(c byte) {
	if _, found := h.chars[c]; found {
		return
	}
	h.chars[c] = struct{}{}
	h.historyStr.WriteByte(c)
	h.historyStr.WriteByte(' ')
}

func (h *history) String() string {
	return h.historyStr.String()
}
