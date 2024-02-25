package distrogo

import (
	"encoding/json"
	"fmt"
	"time"
)

type Track struct {
	DurationSeconds int        `json:"duration_seconds"`
	ISRC            string     `json:"isrc"`
	IsCoverSong     bool       `json:"iscoversong"`
	Links           TrackLinks `json:"links"`
	HasISRC         bool       `json:"hasisrc"`
	ID              int        `json:"id"`
	TrackNum        int        `json:"tracknum"`
	Title           string     `json:"title"`
}

type TrackStats struct {
	Artist          string                  `json:"artist"`
	DurationSeconds int                     `json:"duration_seconds"`
	ISRC            string                  `json:"isrc"`
	Stats           map[string]ServiceStats `json:"stats"`
	Links           TrackStatsLinks         `json:"links"`
	ID              int                     `json:"id"`
	Artwork         ReleaseArtwork          `json:"artwork"`
	TrackNum        int                     `json:"tracknum"`
	Title           string                  `json:"title"`
}

type ServiceStats struct {
	Data StatsData `json:"data"`
}

type StatsData struct {
	Points []StatsPoint `json:"points"`
	Total  int          `json:"total"`
	Period string       `json:"period"`
}

type StatsPoint struct {
	Date  time.Time `json:"date"`
	Count int       `json:"count"`
}

type TrackStatsLinks struct {
	Release Link `json:"release"`
}

type TrackLinks struct {
	Stats        Link `json:"stats"`
	TrackCredits Link `json:"trackcredits"`
	TrackLyrics  Link `json:"tracklyrics"`
}

func (d *DistroKid) GetTracks(releaseID int) ([]Track, error) {
	url := fmt.Sprintf("%s%s", BaseURL, fmt.Sprintf(TracksEndpoint, releaseID))
	data, err := d.DoRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	var tracksIdk struct {
		Data struct {
			Tracks []Track `json:"tracks"`
		} `json:"data"`
	}
	err = json.Unmarshal(data, &tracksIdk)

	if err != nil {
		return nil, err
	}

	return tracksIdk.Data.Tracks, nil
}

func (d *DistroKid) GetTrackStats(trackID int) (TrackStats, error) {
	url := fmt.Sprintf("%s%s", BaseURL, fmt.Sprintf(TrackStatsEndpoint, trackID))
	data, err := d.DoRequest("GET", url, nil)

	if err != nil {
		return TrackStats{}, err
	}

	var trackStatsIdk struct {
		Data TrackStats `json:"data"`
	}
	err = json.Unmarshal(data, &trackStatsIdk)

	if err != nil {
		return TrackStats{}, err
	}

	return trackStatsIdk.Data, nil
}
