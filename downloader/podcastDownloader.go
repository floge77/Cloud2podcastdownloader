package downloader

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const youtubeDownloader = "/usr/local/bin/youtube-dl"
var allowedProviders = []string{"soundcloud", "mixcloud", "youtube"}

type PodcastDownloader struct {
	standardOptions, lengthFilter, downloadArchivePath, channelName, provider string
	YoutubeDLCommand []string
}

func NewPodcastDownloader(minDurationInSeconds int, generalDownloadDirectory string, channelName string, playListToDownload string) *PodcastDownloader {
	pdl := PodcastDownloader{}

	var err error
	pdl.provider, err = pdl.getProviderFromChannelURL(playListToDownload)
	if err != nil {
		log.Fatalf("Could not get provider from: %v or provider not in allowed list: %v", playListToDownload, allowedProviders)
	}

	pdl.standardOptions = "-x -i --audio-format mp3 --embed-thumbnail --add-metadata "
	pdl.lengthFilter = "--match-filter \"duration>" + strconv.Itoa(minDurationInSeconds) + "\" "
	pdl.downloadArchivePath = "--download-archive " + generalDownloadDirectory + "/" + pdl.provider + "/" + channelName + "/archive.txt "
	pdl.channelName = "-o \"" + generalDownloadDirectory + "/" + pdl.provider + "/" + channelName + "/%(title)s__%(uploader)s__%(upload_date)s.%(ext)s\" "

	pdl.generateYoutubeDLCommand(playListToDownload)

	return &pdl
}

// youtube-dl is in Container available at: /usr/local/bin/youtube-dl
func (p *PodcastDownloader) DownloadTracks() {
	cmd := exec.Command(youtubeDownloader, p.YoutubeDLCommand...)
	fmt.Printf("Running command: %v", cmd)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
}

func (p *PodcastDownloader) generateYoutubeDLCommand(playListToDownload string) {
	cmdString := p.standardOptions + p.lengthFilter + p.downloadArchivePath + p.channelName + playListToDownload
	p.YoutubeDLCommand = strings.Split(cmdString, " ")
}


func (p *PodcastDownloader) getProviderFromChannelURL(channelURL string) (provider string, err error) {

	u, err := url.Parse(channelURL)
	provider = strings.Replace(u.Hostname(), "www.", "", -1)
	provider = strings.Split(provider, ".")[0]

	err = checkProviderIsAllowed(provider)
	return
}

func checkProviderIsAllowed(provider string) (err error) {
	providerNotFound := true
	for _, ap := range allowedProviders {
		if strings.Contains(provider, ap) {
			providerNotFound = false
		}
	}
	if providerNotFound {
		return errors.New("Provider not in: " + strings.Join(allowedProviders, ","))
	} else {
		return nil
	}
}