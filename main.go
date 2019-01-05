package main

import (
	"github.com/floge77/c2p/cloud2podcastdownloader/downloader"
	"github.com/floge77/c2p/cloud2podcastdownloader/configReader"
)

func main() {
	// in docker  will be located at /downloads
	config := configReader.GetConfig("/downloads/config.yaml")

	for _, podcast := range config.PodcastsToServe {
		youtubeDownloader := downloader.NewPodcastDownloader(config.MinSetLengthInSeconds, config.DownloadDirectory, podcast.Channel, podcast.PlaylistToDownloadURL)
		youtubeDownloader.DownloadTracks()
	}
}