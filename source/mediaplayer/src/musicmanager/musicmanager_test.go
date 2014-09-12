package musicmanager

import "testing"

func TestOps(t *testing.T) {
    mm := NewMusicManager()
    if mm == nil {
        t.Error("create manager failed")
    }
    if mm.Len() != 0 {
        t.Error("create manager failed, musics not empty")
    }

    music := &Music{"123", "My Heart will go on", "jj", "http://www.music.baidu.com", "mp3"}
    mm.Add(music)
    if mm.Len() != 1 {
        t.Error("add music failed")
    }

    music = mm.Find("My Heart will go on")
    if music == nil {
        t.Error("music not found")
    }

    music = mm.Remove(0)
    if mm.Len() != 0 {
        t.Error("Remove music failed")
    }
}
