package musicplayer

import "fmt"

type Player interface {
    Play(source string)
}

type WAVPlayer struct {
    progress int
}

func (player *WAVPlayer) Play(source string) {
    fmt.Println("WAVPlayer play: ", source)
}

type MP3Player struct {
    progress int
}

func (player *MP3Player) Play(source string) {
    fmt.Println("MP3Player play: ", source)
}

func Play(source string, mtype string) {
    var player Player
    switch mtype {
        case "wav":
            player = &WAVPlayer{12}
        case "mp3":
            player = &MP3Player{24}
        default:
            fmt.Println("invalid type")
            return
    }
    player.Play(source)
}
