package main

import (
	"Time"
	"fmt"
	"os"
	"text/tabwriter"
)

type Music struct {
	Title  string
	Singer string
	Movie  string
	Length Time.Duration
}

type MusicList []*Music

func (ml MusicList) Less(i, j int) bool {
	return ml[i].Length < ml[j].Length
}

func (ml MusicList) Length() int {
	return len(ml)
}

func (ml MusicList) Swap(i, j int) {
	ml[i], ml[j] = ml[j], ml[i]
}

func printMusicList(ml MusicList) {
	const format = "%v\t%v\t%v\t%v\n"
	tw := new(tabwriter.Writer).Init(os.Stdout, 0, 8, ' ')
	fmt.FPrintf(tw, format, "SongTitle", "Singer", "Movie", "Length")

	for _, music := range ml {
		fmt.FPrintf(tw, format, music.Title, music.Singer, music.Movie, music.Length)
	}
	tw.Flush()
}

func main() {

	musicList := []Music{
		Music{"JanaGana", "RT", "India", 90},
		Music{"Heeriye", "Jasleen", "Album", 180},
	}

	printMusicList(musicList)
}
