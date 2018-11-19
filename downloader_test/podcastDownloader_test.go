package musiccloud

import (
	"cloud2podcastdownloader/downloader"
	"testing"
)

func TestPodcastDownloaderGetYoutubeDLCommand(t *testing.T) {
	pdl := downloader.NewPodcastDownloader(1800, "/downloads", "channel", "https://youtube.com/playlist")

	expectedCmd := "/usr/local/bin/youtube-dl -x -i --audio-format mp3 --embed-thumbnail --add-metadata --match-filter \"duration>1800\" --download-archive /downloads/youtube/channel/archive.txt -o \"/downloads/youtube/channel/%(title)s__%(uploader)s__%(upload_date)s.%(ext)s\" https://youtube.com/playlist"
	if pdl.YoutubeDLCommand != expectedCmd {
		t.Errorf("For PodcastDownloader %v the cmd-string should be: \n %v \n but was: \n %v", pdl, expectedCmd, pdl.YoutubeDLCommand)
	}
}
