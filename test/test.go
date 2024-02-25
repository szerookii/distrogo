package main

import (
	"fmt"

	"github.com/szerookii/distrogo"
)

func main() {
	distrokid := distrogo.NewDistroKid("your-bearer-here")

	releases, err := distrokid.GetReleases()

	if err != nil {
		panic(err)
	}

	for _, release := range releases {
		fmt.Printf("%s was released on %s and has %d tracks\n", release.Title, release.ReleaseDate, release.TrackCount)
	}

	release, err := distrokid.GetRelease(releases[0].ID)

	if err != nil {
		panic(err)
	}

	fmt.Printf("The first release is called %s and has %d tracks\n", release.Title, release.TrackCount)

	tracks, _ := distrokid.GetTracks(release.ID)

	for _, track := range tracks {
		fmt.Printf("Track %s has a duration of %d seconds\n", track.Title, track.DurationSeconds)
	}

	trackStats, _ := distrokid.GetTrackStats(tracks[0].ID)

	for name, stat := range trackStats.Stats {
		fmt.Printf("The track %s has %d plays on %s for platform %s\n", tracks[0].Title, stat.Data.Total, stat.Data.Period, name)
	}

	releaseStats, _ := distrokid.GetReleaseStats(release.ID)

	for name, stat := range releaseStats.Stats {
		fmt.Printf("The release %s has %d plays on %s for platform %s\n", release.Title, stat.Data.Total, stat.Data.Period, name)
	}
}
