package musicmanager

import "errors"

type Music struct {
    Id string
    Name string
    Artist string
    Source string
    Type string
}

type MusicManager struct {
    musics []Music
}

func NewMusicManager() *MusicManager {
    return &MusicManager{make([]Music, 0)}
}

func (m *MusicManager) Len() int {
    return len(m.musics)
}

func (m *MusicManager) Get(index int) (music *Music, err error) {
    if index < 0 || index >= len(m.musics) {
        return nil, errors.New("Index out of range")
    }
    return &m.musics[index], nil
}

func (m *MusicManager) Find(name string) *Music {
    if len(m.musics) == 0 {
        return nil
    }
    for _, music := range m.musics {
        if music.Name == name {
            return &music
        }
    }
    return nil
}

func (m *MusicManager) Add(music *Music) {
    m.musics = append(m.musics, *music)
}

func (m *MusicManager) Remove(index int) *Music {
    if index < 0 || index >= len(m.musics) {
        return nil
    }
    removeMusic := &m.musics[index]
    if len(m.musics) == 1 {
        m.musics = make([]Music, 0)
    } else if index == 0 {
        m.musics = m.musics[1:]
    } else if index == len(m.musics) -1 {
        m.musics = m.musics[:index-1]
    } else {
        m.musics = append(m.musics[0:index-1], m.musics[:index+1]...)
    }
    return removeMusic
}
