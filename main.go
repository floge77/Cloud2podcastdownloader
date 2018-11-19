package main

import (
	"cloud2podcastdownloader/configReader"
	"cloud2podcastdownloader/downloader"
)

func main() {
	// in docker  will be located at /downloads
	config := configReader.GetConfig("/downloads/config.yaml")

	for _, podcast := range config.PodcastsToServe {
		youtubeDownloader := downloader.NewPodcastDownloader(config.MinSetLengthInSeconds, config.DownloadDirectory, podcast.Channel, podcast.PlaylistToDownloadURL)
		youtubeDownloader.DownloadTracks()
	}
}