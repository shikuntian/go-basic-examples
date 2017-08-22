package library

import (
	"errors"
)

type MusicManager struct {
	musics []MusicEntry
}

func newMusicManager() *MusicManager {
	return &MusicManager{make([]MusicEntry, 0)}
}

func (m *MusicManager) Len() int {
	return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *MusicEntry, err error) {
	if index < 0 || index >= len(m.musics) {
		return nil, errors.New("Index out of ranges")
	}
	return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *MusicEntry {
	if (len(m.musics) == 0) {
		return nil
	}
	for _, m := range m.musics {
		if (m == name) {
			return &m
		}
	}
	return nil
}

func (m *MusicManager) Add(music *MusicEntry) {
	m.musics = append(m.musics, music)
}

func (m *MusicManager) Remove(idx int) *MusicEntry {
	if (idx < 0 || idx > len(m.musics)) {
		return nil
	}
	removedMusic := &m.musics[idx]
	if idx < len(m.musics) - 1 {
		m.musics = append(m.musics[:idx - 1], m.musics[idx + 1:]...)
	} else if idx == 0 {
		m.musics = make([]MusicEntry, 0)
	} else {
		m.musics = m.musics[:idx - 1]
	}
	return removedMusic
}




