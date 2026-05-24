package backend

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
<<<<<<< HEAD
	"crypto/sha256"
=======
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
<<<<<<< HEAD
=======
	"strconv"
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
	"strings"
	"sync"
	"time"
)

type QobuzDownloader struct {
	client *http.Client
<<<<<<< HEAD
	appID  string
}

type QobuzSearchResponse struct {
	Query  string `json:"query"`
	Tracks struct {
		Limit  int          `json:"limit"`
		Offset int          `json:"offset"`
		Total  int          `json:"total"`
		Items  []QobuzTrack `json:"items"`
	} `json:"tracks"`
=======
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
}

type QobuzTrack struct {
	ID                  int64   `json:"id"`
	Title               string  `json:"title"`
	Version             string  `json:"version"`
	Duration            int     `json:"duration"`
	TrackNumber         int     `json:"track_number"`
	MediaNumber         int     `json:"media_number"`
	ISRC                string  `json:"isrc"`
	Copyright           string  `json:"copyright"`
	MaximumBitDepth     int     `json:"maximum_bit_depth"`
	MaximumSamplingRate float64 `json:"maximum_sampling_rate"`
	Hires               bool    `json:"hires"`
	HiresStreamable     bool    `json:"hires_streamable"`
	ReleaseDateOriginal string  `json:"release_date_original"`
	Performer           struct {
		Name string `json:"name"`
		ID   int64  `json:"id"`
	} `json:"performer"`
	Album struct {
		Title string `json:"title"`
		ID    string `json:"id"`
		Image struct {
			Small     string `json:"small"`
			Thumbnail string `json:"thumbnail"`
			Large     string `json:"large"`
		} `json:"image"`
		Artist struct {
			Name string `json:"name"`
			ID   int64  `json:"id"`
		} `json:"artist"`
		Label struct {
			Name string `json:"name"`
		} `json:"label"`
	} `json:"album"`
}

<<<<<<< HEAD
type QobuzStreamResponse struct {
	URL string `json:"url"`
}

=======
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
type qobuzMusicDLRequest struct {
	URL     string `json:"url"`
	Quality string `json:"quality"`
}

type qobuzMusicDLResponse struct {
	Success     bool   `json:"success"`
	Type        string `json:"type"`
	URLType     string `json:"url_type"`
	TrackID     string `json:"track_id"`
	Quality     string `json:"quality_label"`
	DownloadURL string `json:"download_url"`
	Message     string `json:"message"`
	Error       string `json:"error"`
}

<<<<<<< HEAD
const qobuzMusicDLProbeTrackID int64 = 341032040
=======
type qobuzPublicSearchResponse struct {
	Tracks struct {
		Total int          `json:"total"`
		Items []QobuzTrack `json:"items"`
	} `json:"tracks"`
}

const qobuzProbeTrackID int64 = 341032040
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea

var (
	qobuzMusicDLDebugKeyOnce sync.Once
	qobuzMusicDLDebugKey     string
	qobuzMusicDLDebugKeyErr  error
<<<<<<< HEAD
=======
	qobuzStreamingURLPattern = regexp.MustCompile(`https?://[^\s"'<>\\)]+`)
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
)

var qobuzMusicDLDebugKeySeedParts = [][]byte{
	{0x73, 0x70, 0x6f, 0x74, 0x69, 0x66},
	{0x6c, 0x61, 0x63, 0x3a, 0x71, 0x6f},
	{0x62, 0x75, 0x7a, 0x3a, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x64, 0x6c, 0x3a, 0x76, 0x31},
}

var qobuzMusicDLDebugKeyAAD = []byte{
	0x71, 0x6f, 0x62, 0x75, 0x7a, 0x7c, 0x6d, 0x75, 0x73, 0x69, 0x63, 0x64,
	0x6c, 0x7c, 0x64, 0x65, 0x62, 0x75, 0x67, 0x7c, 0x76, 0x31,
}

var qobuzMusicDLDebugKeyNonce = []byte{
	0x91, 0x2a, 0x5c, 0x77, 0x0f, 0x33, 0xa8, 0x14, 0x62, 0x9d, 0xce, 0x41,
}

var qobuzMusicDLDebugKeyCiphertext = []byte{
	0xf3, 0x4a, 0x83, 0x45, 0x24, 0xb6, 0x22, 0xaf, 0xd6, 0xc3, 0x6e, 0x2d,
	0x56, 0xd1, 0xbb, 0x0b, 0xe9, 0x1b, 0x4f, 0x1c, 0x5f, 0x41, 0x55, 0xc2,
	0xc6, 0xdf, 0xad, 0x21, 0x58, 0xfe, 0xd5, 0xb8, 0x2d, 0x29, 0xf9, 0x9e,
	0x6f, 0xd6,
}

var qobuzMusicDLDebugKeyTag = []byte{
	0x69, 0x0c, 0x42, 0x70, 0x14, 0x83, 0xff, 0x14, 0xc8, 0xbe, 0x17, 0x00,
	0x69, 0xb1, 0xfe, 0xbb,
}

func NewQobuzDownloader() *QobuzDownloader {
	return &QobuzDownloader{
		client: &http.Client{
			Timeout: 60 * time.Second,
		},
<<<<<<< HEAD
		appID: qobuzDefaultAPIAppID,
=======
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
	}
}

func previewQobuzResponseBody(body []byte, maxLen int) string {
	preview := strings.TrimSpace(string(body))
	if len(preview) > maxLen {
		return preview[:maxLen] + "..."
	}
	return preview
}

func buildQobuzOpenTrackURL(trackID int64) string {
	return fmt.Sprintf("https://open.qobuz.com/track/%d", trackID)
}

func getQobuzMusicDLDebugKey() (string, error) {
	qobuzMusicDLDebugKeyOnce.Do(func() {
		hasher := sha256.New()
		for _, part := range qobuzMusicDLDebugKeySeedParts {
			hasher.Write(part)
		}

		block, err := aes.NewCipher(hasher.Sum(nil))
		if err != nil {
			qobuzMusicDLDebugKeyErr = err
			return
		}

		gcm, err := cipher.NewGCM(block)
		if err != nil {
			qobuzMusicDLDebugKeyErr = err
			return
		}

		sealed := make([]byte, 0, len(qobuzMusicDLDebugKeyCiphertext)+len(qobuzMusicDLDebugKeyTag))
		sealed = append(sealed, qobuzMusicDLDebugKeyCiphertext...)
		sealed = append(sealed, qobuzMusicDLDebugKeyTag...)

		plaintext, err := gcm.Open(nil, qobuzMusicDLDebugKeyNonce, sealed, qobuzMusicDLDebugKeyAAD)
		if err != nil {
			qobuzMusicDLDebugKeyErr = err
			return
		}

		qobuzMusicDLDebugKey = string(plaintext)
	})

	if qobuzMusicDLDebugKeyErr != nil {
		return "", qobuzMusicDLDebugKeyErr
	}

	return qobuzMusicDLDebugKey, nil
}

<<<<<<< HEAD
func (q *QobuzDownloader) searchByISRC(isrc string) (*QobuzTrack, error) {
	if strings.HasPrefix(isrc, "qobuz_") {
		trackID := strings.TrimPrefix(isrc, "qobuz_")
		resp, err := doQobuzSignedRequest(http.MethodGet, "track/get", url.Values{"track_id": {trackID}}, q.client)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch track: %w", err)
=======
func firstNonEmptyQobuzValue(values ...string) string {
	for _, value := range values {
		trimmed := strings.TrimSpace(value)
		if trimmed != "" {
			return trimmed
		}
	}
	return ""
}

func normalizeQobuzSearchValue(value string) string {
	replacer := strings.NewReplacer(
		"&", " and ",
		"feat.", " ",
		"ft.", " ",
		"/", " ",
		"-", " ",
		"_", " ",
	)
	normalized := strings.ToLower(strings.TrimSpace(value))
	normalized = replacer.Replace(normalized)
	return strings.Join(strings.Fields(normalized), " ")
}

func qobuzTrackDisplayArtist(track QobuzTrack) string {
	return firstNonEmptyQobuzValue(track.Performer.Name, track.Album.Artist.Name)
}

func qobuzTrackSupportsHiRes(track QobuzTrack) bool {
	if track.Hires || track.HiresStreamable {
		return true
	}
	return track.MaximumBitDepth >= 24 || track.MaximumSamplingRate > 48
}

func scoreQobuzSearchCandidate(track QobuzTrack, spotifyTrackName string, spotifyArtistName string, spotifyAlbumName string) int {
	score := 0

	titleNeedle := normalizeQobuzSearchValue(spotifyTrackName)
	titleHaystack := normalizeQobuzSearchValue(track.Title)
	switch {
	case titleNeedle != "" && titleHaystack == titleNeedle:
		score += 1000
	case titleNeedle != "" && (strings.Contains(titleHaystack, titleNeedle) || strings.Contains(titleNeedle, titleHaystack)):
		score += 500
	}

	artistNeedle := normalizeQobuzSearchValue(spotifyArtistName)
	artistHaystack := normalizeQobuzSearchValue(qobuzTrackDisplayArtist(track))
	switch {
	case artistNeedle != "" && artistHaystack == artistNeedle:
		score += 300
	case artistNeedle != "" && artistHaystack != "" && (strings.Contains(artistHaystack, artistNeedle) || strings.Contains(artistNeedle, artistHaystack)):
		score += 180
	}

	albumNeedle := normalizeQobuzSearchValue(spotifyAlbumName)
	albumHaystack := normalizeQobuzSearchValue(track.Album.Title)
	switch {
	case albumNeedle != "" && albumHaystack == albumNeedle:
		score += 150
	case albumNeedle != "" && albumHaystack != "" && (strings.Contains(albumHaystack, albumNeedle) || strings.Contains(albumNeedle, albumHaystack)):
		score += 90
	}

	if qobuzTrackSupportsHiRes(track) {
		score += 40
	} else if track.MaximumBitDepth >= 16 {
		score += 20
	}

	return score
}

func mapQobuzWJHEQuality(quality string) (int, string) {
	switch strings.TrimSpace(quality) {
	case "27", "7":
		return 2000, "flac"
	case "", "6":
		return 1000, "flac"
	default:
		return 320, "mp3"
	}
}

func buildQobuzWJHEDownloadURL(trackID int64, quality string) string {
	wjheQuality, wjheFormat := mapQobuzWJHEQuality(quality)
	params := url.Values{
		"ID":      {strconv.FormatInt(trackID, 10)},
		"quality": {strconv.Itoa(wjheQuality)},
		"format":  {wjheFormat},
	}
	return GetQobuzWJHEStreamAPIURL() + "?" + params.Encode()
}

func qobuzURLLooksStreamable(raw string) bool {
	candidate := strings.TrimSpace(raw)
	if candidate == "" {
		return false
	}

	parsed, err := url.Parse(candidate)
	if err != nil {
		return false
	}

	return (parsed.Scheme == "http" || parsed.Scheme == "https") && parsed.Host != ""
}

func findQobuzStreamingURLInPayload(payload interface{}) string {
	switch value := payload.(type) {
	case string:
		candidate := strings.ReplaceAll(strings.TrimSpace(value), `\/`, `/`)
		if qobuzURLLooksStreamable(candidate) {
			return candidate
		}
	case []interface{}:
		for _, item := range value {
			if url := findQobuzStreamingURLInPayload(item); url != "" {
				return url
			}
		}
	case map[string]interface{}:
		for _, key := range []string{"download_url", "url", "play_url", "stream_url", "link", "file"} {
			if nested, ok := value[key]; ok {
				if url := findQobuzStreamingURLInPayload(nested); url != "" {
					return url
				}
			}
		}
		for _, nested := range value {
			if url := findQobuzStreamingURLInPayload(nested); url != "" {
				return url
			}
		}
	}

	return ""
}

func extractQobuzStreamingURL(body []byte) string {
	trimmed := strings.TrimSpace(string(body))
	if trimmed == "" {
		return ""
	}

	var directResp struct {
		URL         string `json:"url"`
		DownloadURL string `json:"download_url"`
		Data        struct {
			URL         string `json:"url"`
			DownloadURL string `json:"download_url"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &directResp); err == nil {
		for _, candidate := range []string{
			directResp.DownloadURL,
			directResp.URL,
			directResp.Data.DownloadURL,
			directResp.Data.URL,
		} {
			if qobuzURLLooksStreamable(candidate) {
				return candidate
			}
		}
	}

	var genericPayload interface{}
	if err := json.Unmarshal(body, &genericPayload); err == nil {
		if streamURL := findQobuzStreamingURLInPayload(genericPayload); streamURL != "" {
			return streamURL
		}
	}

	if openIdx := strings.Index(trimmed, "("); openIdx >= 0 {
		if closeIdx := strings.LastIndex(trimmed, ")"); closeIdx > openIdx+1 {
			callbackBody := strings.TrimSpace(trimmed[openIdx+1 : closeIdx])
			if streamURL := extractQobuzStreamingURL([]byte(callbackBody)); streamURL != "" {
				return streamURL
			}
		}
	}

	for _, match := range qobuzStreamingURLPattern.FindAllString(trimmed, -1) {
		candidate := strings.ReplaceAll(match, `\/`, `/`)
		if qobuzURLLooksStreamable(candidate) {
			return candidate
		}
	}

	return ""
}

func newQobuzNoRedirectClient(base *http.Client) *http.Client {
	if base == nil {
		return &http.Client{
			Timeout: 20 * time.Second,
			CheckRedirect: func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			},
		}
	}

	cloned := *base
	if cloned.Timeout == 0 {
		cloned.Timeout = 20 * time.Second
	}
	cloned.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}
	return &cloned
}

func (q *QobuzDownloader) searchByISRC(isrc string, spotifyTrackName string, spotifyArtistName string, spotifyAlbumName string) (*QobuzTrack, error) {
	if strings.HasPrefix(isrc, "qobuz_") {
		trackID := strings.TrimSpace(strings.TrimPrefix(isrc, "qobuz_"))
		resp, err := doQobuzSignedRequest(http.MethodGet, "track/get", url.Values{"track_id": {trackID}}, q.client)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch track from Qobuz public API: %w", err)
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
<<<<<<< HEAD
			return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
=======
			body, _ := io.ReadAll(io.LimitReader(resp.Body, 512))
			return nil, fmt.Errorf("Qobuz public API track/get returned status %d: %s", resp.StatusCode, previewQobuzResponseBody(body, 256))
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
		}

		var trackResp QobuzTrack
		if err := json.NewDecoder(resp.Body).Decode(&trackResp); err != nil {
<<<<<<< HEAD
			return nil, fmt.Errorf("failed to decode response: %w", err)
=======
			return nil, fmt.Errorf("failed to decode Qobuz public track/get response: %w", err)
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
		}

		return &trackResp, nil
	}

<<<<<<< HEAD
	resp, err := doQobuzSignedRequest(http.MethodGet, "track/search", url.Values{
		"query": {isrc},
		"limit": {"1"},
	}, q.client)
	if err != nil {
		return nil, fmt.Errorf("failed to search track: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API returned status %d", resp.StatusCode)
	}

	var searchResp QobuzSearchResponse

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	if len(body) == 0 {
		return nil, fmt.Errorf("API returned empty response")
	}

	if err := json.Unmarshal(body, &searchResp); err != nil {

		bodyStr := string(body)
		if len(bodyStr) > 200 {
			bodyStr = bodyStr[:200] + "..."
		}
		return nil, fmt.Errorf("failed to decode response: %w (response: %s)", err, bodyStr)
	}

	if len(searchResp.Tracks.Items) == 0 {
		return nil, fmt.Errorf("track not found for ISRC: %s", isrc)
	}

	return &searchResp.Tracks.Items[0], nil
}

func buildQobuzAPIURL(apiBase string, trackID int64, quality string) string {
	return fmt.Sprintf("%s%d&quality=%s", apiBase, trackID, quality)
}

func (q *QobuzDownloader) DownloadFromStandard(apiBase string, trackID int64, quality string) (string, error) {
	apiURL := buildQobuzAPIURL(apiBase, trackID, quality)
	req, err := NewRequestWithDefaultHeaders(http.MethodGet, apiURL, nil)
	if err != nil {
		return "", err
	}

	resp, err := q.client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("status %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	if len(body) == 0 {
		return "", fmt.Errorf("empty body")
	}

	var streamResp QobuzStreamResponse
	if err := json.Unmarshal(body, &streamResp); err == nil && streamResp.URL != "" {
		return streamResp.URL, nil
	}

	var nestedResp struct {
		Data struct {
			URL string `json:"url"`
		} `json:"data"`
	}
	if err := json.Unmarshal(body, &nestedResp); err == nil && nestedResp.Data.URL != "" {
		return nestedResp.Data.URL, nil
	}

	return "", fmt.Errorf("invalid response")
=======
	queries := []string{strings.TrimSpace(isrc)}
	if fallbackQuery := strings.TrimSpace(strings.Join([]string{spotifyTrackName, spotifyArtistName}, " ")); fallbackQuery != "" {
		queries = append(queries, fallbackQuery)
	}

	var lastErr error
	for _, query := range queries {
		if strings.TrimSpace(query) == "" {
			continue
		}

		var searchResp qobuzPublicSearchResponse
		if err := doQobuzSignedJSONRequest("track/search", url.Values{
			"query": {strings.TrimSpace(query)},
			"limit": {"10"},
		}, &searchResp); err != nil {
			lastErr = fmt.Errorf("failed to search Qobuz public API: %w", err)
			continue
		}

		if searchResp.Tracks.Total == 0 || len(searchResp.Tracks.Items) == 0 {
			lastErr = fmt.Errorf("track not found for query: %s", query)
			continue
		}

		bestIndex := 0
		bestScore := -1
		for idx, candidate := range searchResp.Tracks.Items {
			score := scoreQobuzSearchCandidate(candidate, spotifyTrackName, spotifyArtistName, spotifyAlbumName)
			if idx == 0 || score > bestScore {
				bestIndex = idx
				bestScore = score
			}
		}

		selected := searchResp.Tracks.Items[bestIndex]
		return &selected, nil
	}

	if lastErr == nil {
		lastErr = fmt.Errorf("track not found for ISRC: %s", isrc)
	}
	return nil, lastErr
}

func (q *QobuzDownloader) DownloadFromWJHE(trackID int64, quality string) (string, error) {
	apiURL := buildQobuzWJHEDownloadURL(trackID, quality)
	client := newQobuzNoRedirectClient(q.client)

	req, err := NewRequestWithDefaultHeaders(http.MethodHead, apiURL, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create WJHE request: %w", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to reach WJHE: %w", err)
	}

	if resp.StatusCode == http.StatusMethodNotAllowed || resp.StatusCode == http.StatusNotImplemented {
		resp.Body.Close()
		req, err = NewRequestWithDefaultHeaders(http.MethodGet, apiURL, nil)
		if err != nil {
			return "", fmt.Errorf("failed to create WJHE fallback request: %w", err)
		}
		resp, err = client.Do(req)
		if err != nil {
			return "", fmt.Errorf("failed to reach WJHE with GET fallback: %w", err)
		}
	}
	defer resp.Body.Close()

	if location := strings.TrimSpace(resp.Header.Get("Location")); qobuzURLLooksStreamable(location) {
		return location, nil
	}

	body, err := io.ReadAll(io.LimitReader(resp.Body, 128*1024))
	if err != nil {
		return "", fmt.Errorf("failed to read WJHE response: %w", err)
	}

	if streamURL := extractQobuzStreamingURL(body); streamURL != "" {
		return streamURL, nil
	}

	if resp.Request != nil && resp.Request.URL != nil {
		if streamURL := strings.TrimSpace(resp.Request.URL.String()); streamURL != "" && streamURL != apiURL && qobuzURLLooksStreamable(streamURL) {
			return streamURL, nil
		}
	}

	if resp.StatusCode < http.StatusOK || resp.StatusCode >= http.StatusBadRequest {
		return "", fmt.Errorf("WJHE returned status %d: %s", resp.StatusCode, previewQobuzResponseBody(body, 256))
	}

	return "", fmt.Errorf("WJHE response did not include a stream URL")
}

func qobuzGDStudioPaddedVersion() string {
	parts := strings.Split(GetQobuzGDStudioVersion(), ".")
	for idx, part := range parts {
		part = strings.TrimSpace(part)
		if len(part) == 1 {
			part = "0" + part
		}
		parts[idx] = part
	}
	return strings.Join(parts, "")
}

func qobuzGDStudioEscapedValue(value string) string {
	return strings.ReplaceAll(url.QueryEscape(strings.TrimSpace(value)), "+", "%20")
}

func (q *QobuzDownloader) getQobuzGDStudioTS9(apiURL string) string {
	fallback := strconv.FormatInt(time.Now().UnixMilli(), 10)
	if len(fallback) >= 9 {
		fallback = fallback[:9]
	}

	client := q.client
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}

	signatureHost := GetQobuzGDStudioSignatureHost(apiURL)
	if signatureHost == "" {
		return fallback
	}

	req, err := NewRequestWithDefaultHeaders(http.MethodGet, fmt.Sprintf("https://%s/time", signatureHost), nil)
	if err != nil {
		return fallback
	}

	resp, err := client.Do(req)
	if err != nil {
		return fallback
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 64))
	if err != nil {
		return fallback
	}

	timestamp := strings.TrimSpace(string(body))
	if len(timestamp) >= 9 {
		return timestamp[:9]
	}

	return fallback
}

func buildQobuzGDStudioSignature(apiURL string, value string, ts9 string) string {
	signatureHost := GetQobuzGDStudioSignatureHost(apiURL)
	signatureBase := fmt.Sprintf("%s|%s|%s|%s", signatureHost, qobuzGDStudioPaddedVersion(), ts9, qobuzGDStudioEscapedValue(value))
	sum := md5.Sum([]byte(signatureBase))
	digest := hex.EncodeToString(sum[:])
	return strings.ToUpper(digest[len(digest)-8:])
}

func mapQobuzGDStudioBitrate(quality string) string {
	switch strings.TrimSpace(quality) {
	case "27", "7":
		return "999"
	case "", "6":
		return "740"
	default:
		return "320"
	}
}

func (q *QobuzDownloader) DownloadFromGDStudio(trackID int64, quality string, apiURL string) (string, error) {
	apiURL = strings.TrimSpace(apiURL)
	if apiURL == "" {
		apiURL = GetQobuzGDStudioPrimaryAPIURL()
	}

	signatureHost := GetQobuzGDStudioSignatureHost(apiURL)
	if signatureHost == "" {
		return "", fmt.Errorf("GDStudio API URL is invalid: %s", apiURL)
	}

	trackIDString := strconv.FormatInt(trackID, 10)
	ts9 := q.getQobuzGDStudioTS9(apiURL)
	payload := url.Values{
		"types":  {"url"},
		"id":     {trackIDString},
		"source": {"qobuz"},
		"br":     {mapQobuzGDStudioBitrate(quality)},
		"s":      {buildQobuzGDStudioSignature(apiURL, trackIDString, ts9)},
	}

	req, err := NewRequestWithDefaultHeaders(http.MethodPost, apiURL, strings.NewReader(payload.Encode()))
	if err != nil {
		return "", fmt.Errorf("failed to create GDStudio request: %w", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Set("Origin", fmt.Sprintf("https://%s", signatureHost))
	req.Header.Set("Referer", fmt.Sprintf("https://%s/", signatureHost))

	resp, err := q.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to reach GDStudio: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(io.LimitReader(resp.Body, 256*1024))
	if err != nil {
		return "", fmt.Errorf("failed to read GDStudio response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("GDStudio returned status %d: %s", resp.StatusCode, previewQobuzResponseBody(body, 256))
	}

	streamURL := extractQobuzStreamingURL(body)
	if streamURL == "" {
		return "", fmt.Errorf("GDStudio response did not include a stream URL: %s", previewQobuzResponseBody(body, 256))
	}

	return streamURL, nil
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
}

func (q *QobuzDownloader) DownloadFromMusicDL(trackID int64, quality string) (string, error) {
	if strings.TrimSpace(quality) == "" {
		quality = "6"
	}

	debugKey, err := getQobuzMusicDLDebugKey()
	if err != nil {
		return "", fmt.Errorf("failed to decrypt MusicDL debug key: %w", err)
	}

	payload, err := json.Marshal(qobuzMusicDLRequest{
		URL:     buildQobuzOpenTrackURL(trackID),
		Quality: strings.TrimSpace(quality),
	})
	if err != nil {
		return "", fmt.Errorf("failed to encode MusicDL request: %w", err)
	}

	req, err := NewRequestWithDefaultHeaders(http.MethodPost, GetQobuzMusicDLDownloadAPIURL(), bytes.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("failed to create MusicDL request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Debug-Key", debugKey)

	resp, err := q.client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to reach MusicDL: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read MusicDL response: %w", err)
	}

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("MusicDL returned status %d: %s", resp.StatusCode, previewQobuzResponseBody(body, 256))
	}

	var downloadResp qobuzMusicDLResponse
	if err := json.Unmarshal(body, &downloadResp); err != nil {
		return "", fmt.Errorf("failed to decode MusicDL response: %w (%s)", err, previewQobuzResponseBody(body, 256))
	}

	if !downloadResp.Success {
		message := strings.TrimSpace(downloadResp.Error)
		if message == "" {
			message = strings.TrimSpace(downloadResp.Message)
		}
		if message == "" {
			message = "MusicDL reported failure"
		}
		return "", fmt.Errorf("%s", message)
	}

	downloadURL := strings.TrimSpace(downloadResp.DownloadURL)
	if downloadURL == "" {
		return "", fmt.Errorf("MusicDL response did not include a download_url")
	}

	return downloadURL, nil
}

<<<<<<< HEAD
func CheckQobuzMusicDLStatus(client *http.Client) bool {
=======
func CheckQobuzMusicDLStatusDetailed(client *http.Client) error {
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
	if client == nil {
		client = &http.Client{Timeout: 4 * time.Second}
	}

<<<<<<< HEAD
	downloader := &QobuzDownloader{client: client, appID: qobuzDefaultAPIAppID}
	_, err := downloader.DownloadFromMusicDL(qobuzMusicDLProbeTrackID, "27")
	return err == nil
=======
	downloader := &QobuzDownloader{client: client}
	_, err := downloader.DownloadFromMusicDL(qobuzProbeTrackID, "27")
	return err
}

func CheckQobuzMusicDLStatus(client *http.Client) bool {
	return CheckQobuzMusicDLStatusDetailed(client) == nil
}

func CheckQobuzWJHEStatusDetailed(client *http.Client) error {
	if client == nil {
		client = &http.Client{Timeout: 4 * time.Second}
	}

	downloader := &QobuzDownloader{client: client}
	_, err := downloader.DownloadFromWJHE(qobuzProbeTrackID, "27")
	return err
}

func CheckQobuzWJHEStatus(client *http.Client) bool {
	return CheckQobuzWJHEStatusDetailed(client) == nil
}

func CheckQobuzGDStudioAPIStatusDetailed(client *http.Client, apiURL string) error {
	if client == nil {
		client = &http.Client{Timeout: 4 * time.Second}
	}

	downloader := &QobuzDownloader{client: client}
	_, err := downloader.DownloadFromGDStudio(qobuzProbeTrackID, "27", apiURL)
	return err
}

func CheckQobuzGDStudioAPIStatus(client *http.Client, apiURL string) bool {
	return CheckQobuzGDStudioAPIStatusDetailed(client, apiURL) == nil
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
}

func (q *QobuzDownloader) GetDownloadURL(trackID int64, quality string, allowFallback bool) (string, error) {
	qualityCode := quality
	if qualityCode == "" || qualityCode == "5" {
		qualityCode = "6"
	}

	fmt.Printf("Getting download URL for track ID: %d with requested quality: %s\n", trackID, qualityCode)

	downloadFunc := func(qual string) (string, error) {
<<<<<<< HEAD
		type Provider struct {
			Name string
			API  string
			Func func() (string, error)
		}

		providerMap := make(map[string]Provider)
		providerIDs := []string{GetQobuzMusicDLDownloadAPIURL()}

		providerMap[GetQobuzMusicDLDownloadAPIURL()] = Provider{
			Name: "MusicDL",
			API:  GetQobuzMusicDLDownloadAPIURL(),
			Func: func() (string, error) {
				return q.DownloadFromMusicDL(trackID, qual)
			},
		}

		for _, api := range GetQobuzStreamAPIBaseURLs() {
			currentAPI := api
			providerIDs = append(providerIDs, currentAPI)
			providerMap[currentAPI] = Provider{
				Name: "Standard(" + currentAPI + ")",
				API:  currentAPI,
				Func: func() (string, error) {
					return q.DownloadFromStandard(currentAPI, trackID, qual)
				},
			}
		}

		orderedProviderIDs := prioritizeProviders("qobuz", providerIDs)
		primaryProviderID := GetQobuzMusicDLDownloadAPIURL()
		if len(orderedProviderIDs) > 1 && orderedProviderIDs[0] != primaryProviderID {
			reordered := []string{primaryProviderID}
			for _, providerID := range orderedProviderIDs {
				if providerID == primaryProviderID {
					continue
				}
				reordered = append(reordered, providerID)
			}
			orderedProviderIDs = reordered
		}
		var lastErr error
		for _, providerID := range orderedProviderIDs {
			p, ok := providerMap[providerID]
=======
		attemptMap := make(map[string]qobuzProviderAttempt)
		attemptIDs := make([]string, 0, len(GetQobuzDownloadProviderURLs()))
		for _, provider := range q.getQobuzDownloadProviders() {
			for _, attempt := range provider.Attempts(trackID, qual) {
				attemptMap[attempt.ID] = attempt
				attemptIDs = append(attemptIDs, attempt.ID)
			}
		}

		orderedProviderIDs := prioritizeProviders("qobuz", attemptIDs)
		orderedProviderIDs = moveQobuzAttemptIDsLast(orderedProviderIDs, GetQobuzMusicDLDownloadAPIURL())
		var lastErr error
		for _, providerID := range orderedProviderIDs {
			attempt, ok := attemptMap[providerID]
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
			if !ok {
				continue
			}

<<<<<<< HEAD
			fmt.Printf("Trying Provider: %s (Quality: %s)...\n", p.Name, qual)

			url, err := p.Func()
			if err == nil {
				fmt.Printf("✓ Success\n")
				recordProviderSuccess("qobuz", p.API)
=======
			fmt.Printf("Trying Provider: %s (Quality: %s)...\n", attempt.Name, qual)

			url, err := attempt.Download()
			if err == nil {
				fmt.Printf("✓ Success\n")
				recordProviderSuccess("qobuz", attempt.ID)
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
				return url, nil
			}

			fmt.Printf("Provider failed: %v\n", err)
<<<<<<< HEAD
			recordProviderFailure("qobuz", p.API)
=======
			recordProviderFailure("qobuz", attempt.ID)
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
			lastErr = err
		}
		return "", lastErr
	}

	url, err := downloadFunc(qualityCode)
	if err == nil {
		return url, nil
	}

	currentQuality := qualityCode

	if currentQuality == "27" && allowFallback {
		fmt.Printf("⚠ Download with quality 27 failed, trying fallback to 7 (24-bit Standard)...\n")
		url, err := downloadFunc("7")
		if err == nil {
			fmt.Println("✓ Success with fallback quality 7")
			return url, nil
		}

		currentQuality = "7"
	}

	if currentQuality == "7" && allowFallback {
		fmt.Printf("⚠ Download with quality 7 failed, trying fallback to 6 (16-bit Lossless)...\n")
		url, err := downloadFunc("6")
		if err == nil {
			fmt.Println("✓ Success with fallback quality 6")
			return url, nil
		}
	}

	return "", fmt.Errorf("all APIs and fallbacks failed. Last error: %v", err)
}

func (q *QobuzDownloader) DownloadFile(url, filepath string) error {
	fmt.Println("Starting file download...")

	downloadClient := &http.Client{
		Timeout: 5 * time.Minute,
	}

	req, err := NewRequestWithDefaultHeaders(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("failed to create download request: %w", err)
	}

	resp, err := downloadClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download file: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("download failed with status %d", resp.StatusCode)
	}

	fmt.Printf("Creating file: %s\n", filepath)
	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	fmt.Println("Downloading...")

	pw := NewProgressWriter(out)
	_, err = io.Copy(pw, resp.Body)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	fmt.Printf("\rDownloaded: %.2f MB (Complete)\n", float64(pw.GetTotal())/(1024*1024))
	return nil
}

func (q *QobuzDownloader) DownloadCoverArt(coverURL, filepath string) error {
	if coverURL == "" {
		return fmt.Errorf("no cover URL provided")
	}

	req, err := NewRequestWithDefaultHeaders(http.MethodGet, coverURL, nil)
	if err != nil {
		return fmt.Errorf("failed to create cover request: %w", err)
	}

	resp, err := q.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to download cover: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("cover download failed with status %d", resp.StatusCode)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("failed to create cover file: %w", err)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func buildQobuzFilename(title, artist, album, albumArtist, releaseDate string, trackNumber, discNumber int, format string, includeTrackNumber bool, position int, useAlbumTrackNumber bool, extra ...string) string {
	var filename string
	isrc := ""
	if len(extra) > 0 {
		isrc = SanitizeOptionalFilename(extra[0])
	}

	numberToUse := position
	if useAlbumTrackNumber && trackNumber > 0 {
		numberToUse = trackNumber
	}

	year := ""
	if len(releaseDate) >= 4 {
		year = releaseDate[:4]
	}

	if strings.Contains(format, "{") {
		filename = format
		filename = strings.ReplaceAll(filename, "{title}", title)
		filename = strings.ReplaceAll(filename, "{artist}", artist)
		filename = strings.ReplaceAll(filename, "{album}", album)
		filename = strings.ReplaceAll(filename, "{album_artist}", albumArtist)
		filename = strings.ReplaceAll(filename, "{year}", year)
		filename = strings.ReplaceAll(filename, "{date}", SanitizeFilename(releaseDate))
		filename = strings.ReplaceAll(filename, "{isrc}", isrc)

		if discNumber > 0 {
			filename = strings.ReplaceAll(filename, "{disc}", fmt.Sprintf("%d", discNumber))
		} else {
			filename = strings.ReplaceAll(filename, "{disc}", "")
		}

		if numberToUse > 0 {
			filename = strings.ReplaceAll(filename, "{track}", fmt.Sprintf("%02d", numberToUse))
		} else {

			filename = regexp.MustCompile(`\{track\}\.\s*`).ReplaceAllString(filename, "")
			filename = regexp.MustCompile(`\{track\}\s*-\s*`).ReplaceAllString(filename, "")
			filename = regexp.MustCompile(`\{track\}\s*`).ReplaceAllString(filename, "")
		}
	} else {

		switch format {
		case "artist-title":
			filename = fmt.Sprintf("%s - %s", artist, title)
		case "title":
			filename = title
		default:
			filename = fmt.Sprintf("%s - %s", title, artist)
		}

		if includeTrackNumber && position > 0 {
			filename = fmt.Sprintf("%02d. %s", numberToUse, filename)
		}
	}

	return filename + ".flac"
}

func (q *QobuzDownloader) DownloadTrack(spotifyID, outputDir, quality, filenameFormat string, includeTrackNumber bool, position int, spotifyTrackName, spotifyArtistName, spotifyAlbumName, spotifyAlbumArtist, spotifyReleaseDate string, useAlbumTrackNumber bool, spotifyCoverURL string, embedMaxQualityCover bool, spotifyTrackNumber, spotifyDiscNumber, spotifyTotalTracks int, spotifyTotalDiscs int, spotifyCopyright, spotifyPublisher, spotifyComposer, metadataSeparator, spotifyURL string, allowFallback bool, useFirstArtistOnly bool, useSingleGenre bool, embedGenre bool) (string, error) {
	var isrc string
	if spotifyID != "" {
		linkClient := NewSongLinkClient()
		resolvedISRC, err := linkClient.GetISRCDirect(spotifyID)
		if err != nil {
			return "", fmt.Errorf("failed to get ISRC: %v", err)
		}
		isrc = resolvedISRC
	} else {
		return "", fmt.Errorf("spotify ID is required for Qobuz download")
	}

	return q.DownloadTrackWithISRC(isrc, outputDir, quality, filenameFormat, includeTrackNumber, position, spotifyTrackName, spotifyArtistName, spotifyAlbumName, spotifyAlbumArtist, spotifyReleaseDate, useAlbumTrackNumber, spotifyCoverURL, embedMaxQualityCover, spotifyTrackNumber, spotifyDiscNumber, spotifyTotalTracks, spotifyTotalDiscs, spotifyCopyright, spotifyPublisher, spotifyComposer, metadataSeparator, spotifyURL, allowFallback, useFirstArtistOnly, useSingleGenre, embedGenre)
}

func (q *QobuzDownloader) DownloadTrackWithISRC(isrc, outputDir, quality, filenameFormat string, includeTrackNumber bool, position int, spotifyTrackName, spotifyArtistName, spotifyAlbumName, spotifyAlbumArtist, spotifyReleaseDate string, useAlbumTrackNumber bool, spotifyCoverURL string, embedMaxQualityCover bool, spotifyTrackNumber, spotifyDiscNumber, spotifyTotalTracks int, spotifyTotalDiscs int, spotifyCopyright, spotifyPublisher, spotifyComposer, metadataSeparator, spotifyURL string, allowFallback bool, useFirstArtistOnly bool, useSingleGenre bool, embedGenre bool) (string, error) {
	fmt.Printf("Fetching track info for ISRC: %s\n", isrc)

	metaChan := make(chan Metadata, 1)
	if embedGenre && isrc != "" {
		go func() {
			if ShouldSkipMusicBrainzMetadataFetch() {
				fmt.Println("Skipping MusicBrainz metadata fetch because status check is offline.")
				metaChan <- Metadata{}
			} else {
				fmt.Println("Fetching MusicBrainz metadata...")
				if fetchedMeta, err := FetchMusicBrainzMetadata(isrc, spotifyTrackName, spotifyArtistName, spotifyAlbumName, useSingleGenre, embedGenre); err == nil {
					fmt.Println("✓ MusicBrainz metadata fetched")
					metaChan <- fetchedMeta
				} else {
					fmt.Printf("Warning: Failed to fetch MusicBrainz metadata: %v\n", err)
					metaChan <- Metadata{}
				}
			}
		}()
	} else {
		close(metaChan)
	}

	if outputDir != "." {
		if err := os.MkdirAll(outputDir, 0755); err != nil {
			return "", fmt.Errorf("failed to create output directory: %w", err)
		}
	}

<<<<<<< HEAD
	track, err := q.searchByISRC(isrc)
=======
	track, err := q.searchByISRC(isrc, spotifyTrackName, spotifyArtistName, spotifyAlbumName)
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
	if err != nil {
		return "", err
	}

	artists := spotifyArtistName
	trackTitle := spotifyTrackName
	albumTitle := spotifyAlbumName

	fmt.Printf("Found track: %s - %s\n", artists, trackTitle)
	fmt.Printf("Album: %s\n", albumTitle)

	qualityInfo := "Standard"
	if track.Hires {
<<<<<<< HEAD
		qualityInfo = fmt.Sprintf("Hi-Res (%d-bit / %.1f kHz)", track.MaximumBitDepth, track.MaximumSamplingRate)
=======
		if track.MaximumBitDepth > 0 && track.MaximumSamplingRate > 0 {
			qualityInfo = fmt.Sprintf("Hi-Res (%d-bit / %.1f kHz)", track.MaximumBitDepth, track.MaximumSamplingRate)
		} else if track.MaximumBitDepth > 0 {
			qualityInfo = fmt.Sprintf("Hi-Res available (%d-bit)", track.MaximumBitDepth)
		} else {
			qualityInfo = "Hi-Res available"
		}
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
	}
	fmt.Printf("Quality: %s\n", qualityInfo)

	fmt.Println("Getting download URL...")
	downloadURL, err := q.GetDownloadURL(track.ID, quality, allowFallback)
	if err != nil {
		return "", fmt.Errorf("failed to get download URL: %w", err)
	}

	if downloadURL == "" {
		return "", fmt.Errorf("received empty download URL")
	}

	urlPreview := downloadURL
	if len(downloadURL) > 60 {
		urlPreview = downloadURL[:60] + "..."
	}
	fmt.Printf("Download URL obtained: %s\n", urlPreview)

	safeArtist := sanitizeFilename(artists)
	safeAlbumArtist := sanitizeFilename(spotifyAlbumArtist)

	if useFirstArtistOnly {
		safeArtist = sanitizeFilename(GetFirstArtist(artists))
		safeAlbumArtist = sanitizeFilename(GetFirstArtist(spotifyAlbumArtist))
	}

	safeTitle := sanitizeFilename(trackTitle)
	safeAlbum := sanitizeFilename(albumTitle)

	filename := buildQobuzFilename(safeTitle, safeArtist, safeAlbum, safeAlbumArtist, spotifyReleaseDate, spotifyTrackNumber, spotifyDiscNumber, filenameFormat, includeTrackNumber, position, useAlbumTrackNumber, isrc)
	filepath := filepath.Join(outputDir, filename)
	filepath, alreadyExists := ResolveOutputPathForDownload(filepath, GetRedownloadWithSuffixSetting())
	if alreadyExists {
		fmt.Printf("File already exists: %s (%.2f MB)\n", filepath, float64(mustFileSize(filepath))/(1024*1024))
		return "EXISTS:" + filepath, nil
	}

	fmt.Printf("Downloading FLAC file to: %s\n", filepath)
	if err := q.DownloadFile(downloadURL, filepath); err != nil {
		return "", fmt.Errorf("failed to download file: %w", err)
	}

	fmt.Printf("Downloaded: %s\n", filepath)

	coverPath := ""

	if spotifyCoverURL != "" {
		coverPath = filepath + ".cover.jpg"
		coverClient := NewCoverClient()
		if err := coverClient.DownloadCoverToPath(spotifyCoverURL, coverPath, embedMaxQualityCover); err != nil {
			fmt.Printf("Warning: Failed to download Spotify cover: %v\n", err)
			coverPath = ""
		} else {
			defer os.Remove(coverPath)
			fmt.Println("Spotify cover downloaded")
		}
	}

	var mbMeta Metadata
	if isrc != "" {
		mbMeta = <-metaChan
	}

	fmt.Println("Embedding metadata and cover art...")

	trackNumberToEmbed := spotifyTrackNumber
	if trackNumberToEmbed == 0 {
		trackNumberToEmbed = 1
	}

	upc := ""
	if identifiers, err := GetSpotifyTrackIdentifiersDirect(spotifyURL); err == nil || identifiers.ISRC != "" || identifiers.UPC != "" {
		if strings.TrimSpace(isrc) == "" && strings.TrimSpace(identifiers.ISRC) != "" {
			isrc = strings.TrimSpace(identifiers.ISRC)
		}
		upc = strings.TrimSpace(identifiers.UPC)
	}

	metadata := Metadata{
		Title:       trackTitle,
		Artist:      artists,
		Album:       albumTitle,
		AlbumArtist: spotifyAlbumArtist,
		Date:        spotifyReleaseDate,
		TrackNumber: trackNumberToEmbed,
		TotalTracks: spotifyTotalTracks,
		DiscNumber:  spotifyDiscNumber,
		TotalDiscs:  spotifyTotalDiscs,
		URL:         spotifyURL,
		Comment:     spotifyURL,
		Copyright:   spotifyCopyright,
		Publisher:   spotifyPublisher,
		Composer:    spotifyComposer,
		Separator:   metadataSeparator,
		Description: "https://github.com/spotbye/SpotiFLAC",
		ISRC:        isrc,
		UPC:         upc,
		Genre:       mbMeta.Genre,
	}

	if err := EmbedMetadata(filepath, metadata, coverPath); err != nil {
		return "", fmt.Errorf("failed to embed metadata: %w", err)
	}

	fmt.Println("Metadata embedded successfully!")
	return filepath, nil
}
