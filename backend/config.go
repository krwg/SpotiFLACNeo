package backend

import (
	"encoding/json"
<<<<<<< HEAD
=======
	"errors"
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
	"os"
	"path/filepath"
	"strings"
)

<<<<<<< HEAD
=======
const legacyTidalAPICacheFile = "tidal-api-urls.json"

func normalizeCustomTidalAPIValue(value interface{}) string {
	customAPI, _ := value.(string)
	customAPI = strings.TrimRight(strings.TrimSpace(customAPI), "/")
	if strings.HasPrefix(customAPI, "https://") {
		return customAPI
	}
	return ""
}

func sanitizeDownloaderValue(value interface{}, allowTidal bool) string {
	downloader, _ := value.(string)
	switch strings.TrimSpace(strings.ToLower(downloader)) {
	case "tidal":
		if allowTidal {
			return "tidal"
		}
		return "auto"
	case "qobuz":
		return "qobuz"
	case "amazon":
		return "amazon"
	default:
		return "auto"
	}
}

func sanitizeAutoOrderValue(value interface{}, allowTidal bool) string {
	autoOrder, _ := value.(string)
	allowed := map[string]struct{}{
		"qobuz":  {},
		"amazon": {},
	}
	fallback := "qobuz-amazon"
	if allowTidal {
		allowed["tidal"] = struct{}{}
		fallback = "tidal-qobuz-amazon"
	}

	seen := make(map[string]struct{})
	parts := make([]string, 0, 3)
	for _, rawPart := range strings.Split(strings.TrimSpace(strings.ToLower(autoOrder)), "-") {
		part := strings.TrimSpace(rawPart)
		if part == "" {
			continue
		}
		if _, ok := allowed[part]; !ok {
			continue
		}
		if _, ok := seen[part]; ok {
			continue
		}
		seen[part] = struct{}{}
		parts = append(parts, part)
	}

	if len(parts) < 2 {
		return fallback
	}

	return strings.Join(parts, "-")
}

func SanitizeSettingsMap(settings map[string]interface{}) map[string]interface{} {
	if settings == nil {
		return nil
	}

	sanitized := make(map[string]interface{}, len(settings))
	for key, value := range settings {
		sanitized[key] = value
	}

	customAPI := normalizeCustomTidalAPIValue(sanitized["customTidalApi"])
	sanitized["customTidalApi"] = customAPI
	allowTidal := customAPI != ""
	sanitized["downloader"] = sanitizeDownloaderValue(sanitized["downloader"], allowTidal)
	sanitized["autoOrder"] = sanitizeAutoOrderValue(sanitized["autoOrder"], allowTidal)

	return sanitized
}

func CleanupLegacyTidalPublicAPIState() error {
	appDir, err := EnsureAppDir()
	if err != nil {
		return err
	}

	cachePath := filepath.Join(appDir, legacyTidalAPICacheFile)
	if err := os.Remove(cachePath); err != nil && !errors.Is(err, os.ErrNotExist) {
		return err
	}

	return nil
}

func SanitizePersistedConfigSettings() error {
	configPath, err := GetConfigPath()
	if err != nil {
		return err
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return err
	}

	var settings map[string]interface{}
	if err := json.Unmarshal(data, &settings); err != nil {
		return err
	}

	sanitized := SanitizeSettingsMap(settings)
	payload, err := json.MarshalIndent(sanitized, "", "  ")
	if err != nil {
		return err
	}

	return os.WriteFile(configPath, payload, 0o644)
}

>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
func GetDefaultMusicPath() string {

	homeDir, err := os.UserHomeDir()
	if err != nil {

		return "C:\\Users\\Public\\Music"
	}

	return filepath.Join(homeDir, "Music")
}

func GetConfigPath() (string, error) {
	dir, err := EnsureAppDir()
	if err != nil {
		return "", err
	}

	return filepath.Join(dir, "config.json"), nil
}

func LoadConfigSettings() (map[string]interface{}, error) {
	configPath, err := GetConfigPath()
	if err != nil {
		return nil, err
	}

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		return nil, nil
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, err
	}

	var settings map[string]interface{}
	if err := json.Unmarshal(data, &settings); err != nil {
		return nil, err
	}

<<<<<<< HEAD
	return settings, nil
=======
	return SanitizeSettingsMap(settings), nil
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
}

func GetRedownloadWithSuffixSetting() bool {
	settings, err := LoadConfigSettings()
	if err != nil || settings == nil {
		return false
	}

	enabled, _ := settings["redownloadWithSuffix"].(bool)
	return enabled
}

func GetCustomTidalAPISetting() string {
	settings, err := LoadConfigSettings()
	if err != nil || settings == nil {
		return ""
	}

<<<<<<< HEAD
	customAPI, _ := settings["customTidalApi"].(string)
	customAPI = strings.TrimRight(strings.TrimSpace(customAPI), "/")
	if strings.HasPrefix(customAPI, "https://") {
		return customAPI
	}

	return ""
=======
	return normalizeCustomTidalAPIValue(settings["customTidalApi"])
>>>>>>> 0c3a7b70afc89d776b23941087a0a19a741988ea
}

func normalizeExistingFileCheckMode(value string) string {
	switch strings.TrimSpace(strings.ToLower(value)) {
	case "isrc", "upc":
		return "isrc"
	default:
		return "filename"
	}
}

func GetExistingFileCheckModeSetting() string {
	settings, err := LoadConfigSettings()
	if err != nil || settings == nil {
		return "filename"
	}

	rawMode, _ := settings["existingFileCheckMode"].(string)
	return normalizeExistingFileCheckMode(rawMode)
}

func GetLinkResolverSetting() string {
	settings, err := LoadConfigSettings()
	if err != nil || settings == nil {
		return linkResolverProviderDeezerSongLink
	}

	resolver, _ := settings["linkResolver"].(string)
	switch strings.TrimSpace(strings.ToLower(resolver)) {
	case "songlink", linkResolverProviderDeezerSongLink:
		return linkResolverProviderDeezerSongLink
	case "songstats":
		return linkResolverProviderSongstats
	case "":
		return linkResolverProviderDeezerSongLink
	default:
		return linkResolverProviderDeezerSongLink
	}
}

func GetLinkResolverAllowFallback() bool {
	settings, err := LoadConfigSettings()
	if err != nil || settings == nil {
		return true
	}

	allowFallback, ok := settings["allowResolverFallback"].(bool)
	if !ok {
		return true
	}

	return allowFallback
}
