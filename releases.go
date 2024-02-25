package distrogo

import (
	"encoding/json"
	"fmt"
	"time"
)

type Release struct {
	Stores              []string        `json:"stores,omitempty"`
	Artist              string          `json:"artist,omitempty"`
	IsDeleted           bool            `json:"isdeleted,omitempty"`
	NumTracks           string          `json:"numtracks,omitempty"` // Consider changing to int if applicable
	ReleaseDate         time.Time       `json:"releasedate,omitempty"`
	CanDeleteFromStores bool            `json:"candeletefromstores,omitempty"`
	Tracks              []Track         `json:"tracks,omitempty"`
	StoreIcons          map[string]Icon `json:"store_icons,omitempty"`
	UploadDate          time.Time       `json:"uploaddate,omitempty"`
	IsPriority          bool            `json:"ispriority,omitempty"`
	UUID                string          `json:"uuid,omitempty"`
	ReleaseDateString   string          `json:"releasedate_string,omitempty"`
	Links               ReleaseLinks    `json:"links,omitempty"`
	TrackCount          int             `json:"trackcount,omitempty"`
	CanEdit             bool            `json:"canedit,omitempty"`
	UploadDateString    string          `json:"uploaddate_string,omitempty"`
	ID                  int             `json:"id,omitempty"`
	Artwork             ReleaseArtwork  `json:"artwork,omitempty"`
	ReleaseStatus       string          `json:"release_status,omitempty"`
	Single              bool            `json:"single,omitempty"`
	UPC                 string          `json:"upc,omitempty"`
	IsVerified          bool            `json:"isverified,omitempty"`
	RecordLabel         string          `json:"recordlabel,omitempty"`
	Title               string          `json:"title,omitempty"`
	Alerts              []interface{}   `json:"alerts,omitempty"` // can't find right type for this
}

type ReleaseStats struct {
	Artist  string                  `json:"artist"`
	Tracks  []Track                 `json:"tracks"`
	Stats   map[string]ServiceStats `json:"stats"`
	ID      int                     `json:"id"`
	Artwork ReleaseArtwork          `json:"artwork"`
	Title   string                  `json:"title"`
}

type Icon struct {
	Icon string `json:"icon,omitempty"`
	Alt  string `json:"alt,omitempty"`
}

type ReleaseLinks struct {
	ReleaseEdit Link `json:"Releaseedit"`
	Detail      Link `json:"detail"`
	Tracks      Link `json:"tracks"`
	HyperFollow Link `json:"hyperfollow"`
}

type Link struct {
	Href string `json:"href"`
	Type string `json:"type"`
	Rel  string `json:"rel"`
}

type ReleaseArtwork struct {
	Size300 string `json:"300"`
	Full    string `json:"full"`
	Size100 string `json:"100"`
}

func (d *DistroKid) GetReleases() ([]Release, error) {
	url := fmt.Sprintf("%s%s", BaseURL, ReleasesEndpoint)
	data, err := d.DoRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	var releasesIdk struct {
		Data struct {
			Items []Release `json:"items"`
		} `json:"data"`
	}
	err = json.Unmarshal(data, &releasesIdk)

	if err != nil {
		return nil, err
	}

	return releasesIdk.Data.Items, nil
}

func (d *DistroKid) GetRelease(id int) (Release, error) {
	url := fmt.Sprintf("%s%s", BaseURL, fmt.Sprintf(ReleaseEndpoint, id))
	data, err := d.DoRequest("GET", url, nil)

	if err != nil {
		return Release{}, err
	}

	var releaseResp struct {
		Data Release `json:"data"`
	}
	err = json.Unmarshal(data, &releaseResp)

	if err != nil {
		return Release{}, err
	}

	return releaseResp.Data, nil
}

func (d *DistroKid) GetReleaseStats(id int) (ReleaseStats, error) {
	url := fmt.Sprintf("%s%s", BaseURL, fmt.Sprintf(ReleaseStatsEndpoint, id))
	data, err := d.DoRequest("GET", url, nil)

	if err != nil {
		return ReleaseStats{}, err
	}

	var releaseStatsResp struct {
		Data ReleaseStats `json:"data"`
	}
	err = json.Unmarshal(data, &releaseStatsResp)

	if err != nil {
		return ReleaseStats{}, err
	}

	return releaseStatsResp.Data, nil
}
