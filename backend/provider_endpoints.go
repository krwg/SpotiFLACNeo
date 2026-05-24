package backend

<<<<<<< HEAD
const amazonMusicAPIBaseURL = "https://amazon.spotbye.qzz.io"
const qobuzMusicDLDownloadAPIURL = "https://www.musicdl.me/api/qobuz/download"

var defaultQobuzStreamAPIBaseURLs = []string{
	"https://dab.yeet.su/api/stream?trackId=",
	"https://dabmusic.xyz/api/stream?trackId=",
}

func GetQobuzStreamAPIBaseURLs() []string {
	return append([]string(nil), defaultQobuzStreamAPIBaseURLs...)
=======
import (
	"net/url"
	"strings"
)

const amazonMusicAPIBaseURL = "https://amazon.spotbye.qzz.io"

const (
	qobuzWJHEBaseURL           = "https://music.wjhe.top"
	qobuzWJHESearchAPIURL      = qobuzWJHEBaseURL + "/api/music/qobuz/search"
	qobuzWJHEStreamAPIURL      = qobuzWJHEBaseURL + "/api/music/qobuz/url"
	qobuzMusicDLDownloadAPIURL = "https://www.musicdl.me/api/qobuz/download"
	qobuzGDStudioAPIURLXYZ     = "https://music.gdstudio.xyz/api.php"
	qobuzGDStudioAPIURLORG     = "https://music.gdstudio.org/api.php"
	qobuzGDStudioVersion       = "2026.5.10"
)

var defaultQobuzDownloadProviderURLs = []string{
	qobuzWJHEStreamAPIURL,
	qobuzGDStudioAPIURLXYZ,
	qobuzGDStudioAPIURLORG,
	qobuzMusicDLDownloadAPIURL,
}

func GetQobuzDownloadProviderURLs() []string {
	return append([]string(nil), defaultQobuzDownloadProviderURLs...)
}

func GetQobuzWJHESearchAPIURL() string {
	return qobuzWJHESearchAPIURL
}

func GetQobuzWJHEStreamAPIURL() string {
	return qobuzWJHEStreamAPIURL
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
}

func GetQobuzMusicDLDownloadAPIURL() string {
	return qobuzMusicDLDownloadAPIURL
}

<<<<<<< HEAD
=======
func GetQobuzGDStudioAPIURLs() []string {
	return []string{qobuzGDStudioAPIURLXYZ, qobuzGDStudioAPIURLORG}
}

func GetQobuzGDStudioPrimaryAPIURL() string {
	return qobuzGDStudioAPIURLXYZ
}

func GetQobuzGDStudioFallbackAPIURL() string {
	return qobuzGDStudioAPIURLORG
}

func GetQobuzGDStudioSignatureHost(apiURL string) string {
	parsed, err := url.Parse(strings.TrimSpace(apiURL))
	if err != nil || strings.TrimSpace(parsed.Host) == "" {
		return ""
	}
	return strings.TrimSpace(parsed.Host)
}

func GetQobuzGDStudioVersion() string {
	return qobuzGDStudioVersion
}

func IsQobuzWJHEProviderURL(raw string) bool {
	candidate := strings.TrimSpace(raw)
	return candidate == qobuzWJHEStreamAPIURL || strings.HasPrefix(candidate, qobuzWJHEStreamAPIURL+"?")
}

func IsQobuzMusicDLProviderURL(raw string) bool {
	return strings.EqualFold(strings.TrimSpace(raw), qobuzMusicDLDownloadAPIURL)
}

func IsQobuzGDStudioProviderURL(raw string) bool {
	candidate := strings.TrimSpace(raw)
	for _, apiURL := range GetQobuzGDStudioAPIURLs() {
		if candidate == apiURL || strings.HasPrefix(candidate, apiURL+"?") {
			return true
		}
	}
	return false
}

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
func GetAmazonMusicAPIBaseURL() string {
	return amazonMusicAPIBaseURL
}
