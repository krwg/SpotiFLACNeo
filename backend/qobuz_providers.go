package backend

type qobuzDownloadProvider interface {
	Name() string
	Attempts(trackID int64, quality string) []qobuzProviderAttempt
}

type qobuzProviderAttempt struct {
	Name     string
	ID       string
	Download func() (string, error)
}

type QobuzProviderWJHE struct {
	downloader *QobuzDownloader
}

func (p QobuzProviderWJHE) Name() string {
	return "QobuzProviderWJHE"
}

func (p QobuzProviderWJHE) Attempts(trackID int64, quality string) []qobuzProviderAttempt {
	return []qobuzProviderAttempt{
		{
			Name: p.Name(),
			ID:   GetQobuzWJHEStreamAPIURL(),
			Download: func() (string, error) {
				return p.downloader.DownloadFromWJHE(trackID, quality)
			},
		},
	}
}

type QobuzProviderMusicDL struct {
	downloader *QobuzDownloader
}

func (p QobuzProviderMusicDL) Name() string {
	return "QobuzProviderMusicDL"
}

func (p QobuzProviderMusicDL) Attempts(trackID int64, quality string) []qobuzProviderAttempt {
	return []qobuzProviderAttempt{
		{
			Name: p.Name(),
			ID:   GetQobuzMusicDLDownloadAPIURL(),
			Download: func() (string, error) {
				return p.downloader.DownloadFromMusicDL(trackID, quality)
			},
		},
	}
}

type QobuzProviderGDStudio struct {
	downloader *QobuzDownloader
}

func (p QobuzProviderGDStudio) Name() string {
	return "QobuzProviderGDStudio"
}

func (p QobuzProviderGDStudio) Attempts(trackID int64, quality string) []qobuzProviderAttempt {
	attempts := make([]qobuzProviderAttempt, 0, len(GetQobuzGDStudioAPIURLs()))
	for _, apiURL := range GetQobuzGDStudioAPIURLs() {
		currentAPIURL := apiURL
		attempts = append(attempts, qobuzProviderAttempt{
			Name: p.Name(),
			ID:   currentAPIURL,
			Download: func() (string, error) {
				return p.downloader.DownloadFromGDStudio(trackID, quality, currentAPIURL)
			},
		})
	}
	return attempts
}

func (q *QobuzDownloader) getQobuzDownloadProviders() []qobuzDownloadProvider {
	return []qobuzDownloadProvider{
		QobuzProviderWJHE{downloader: q},
		QobuzProviderGDStudio{downloader: q},
		QobuzProviderMusicDL{downloader: q},
	}
}

func moveQobuzAttemptIDsLast(providerIDs []string, lastIDs ...string) []string {
	if len(providerIDs) == 0 || len(lastIDs) == 0 {
		return append([]string(nil), providerIDs...)
	}

	lastIDSet := make(map[string]struct{}, len(lastIDs))
	for _, providerID := range lastIDs {
		lastIDSet[providerID] = struct{}{}
	}

	ordered := make([]string, 0, len(providerIDs))
	trailing := make([]string, 0, len(providerIDs))
	for _, providerID := range providerIDs {
		if _, ok := lastIDSet[providerID]; ok {
			trailing = append(trailing, providerID)
			continue
		}
		ordered = append(ordered, providerID)
	}

	return append(ordered, trailing...)
}
