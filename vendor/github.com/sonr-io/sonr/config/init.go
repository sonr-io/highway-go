package config

import (
	"errors"
	"os"
	"runtime"

	"github.com/denisbrodbeck/machineid"
	"github.com/kataras/golog"
	"github.com/spf13/viper"
)

var (
	logger = golog.Default.Child("internal/io")

	// Path Manipulation Errors
	ErrDuplicateFilePathOption    = errors.New("Duplicate file path option")
	ErrPrefixSuffixSetWithReplace = errors.New("Prefix or Suffix set with Replace.")
	ErrSeparatorLength            = errors.New("Separator length must be 1.")
	ErrNoFileNameSet              = errors.New("File name was not set by options.")

	// Device ID Errors
	ErrEmptyDeviceID = errors.New("Device ID cannot be empty")
	ErrMissingEnvVar = errors.New("Cannot set EnvVariable with empty value")

	// Directory errors
	ErrDirectoryInvalid = errors.New("Directory Type is invalid")
	ErrDirectoryUnset   = errors.New("Directory path has not been set")
	ErrDirectoryJoin    = errors.New("Failed to join directory path")

	// Setup Variables
	isSaved  = false
	instance *SonrConfig
)

// GetString returns the configuration value for the given key as a string.
func GetString(key string) string {
	return viper.GetString(key)
}

// GetInt returns the configuration value for the given key as an int.
func GetInt(key string) int {
	return viper.GetInt(key)
}

// Set sets the configuration value for the given key.
func SetKey(key string, value interface{}) {
	viper.Set(key, value)
}

func Load() (*SonrConfig, error) {
	// Return the instance if it has been initialized.
	if instance != nil {
		return instance, nil
	}

	// Get the user home directory.
	hp, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// Get user cache directory.
	tp, err := os.UserCacheDir()
	if err != nil {
		return nil, err
	}

	// Get the user config directory.
	sp, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	wf, err := Folder(hp).CreateFolder(".sonr-wallet")
	if err != nil {
		return nil, err
	}

	id, err := machineid.ID()
	if err != nil {
		return nil, err
	}

	// Create the configuration object.
	config := &SonrConfig{
		HomeDir:         hp,
		CacheDir:        tp,
		ConfigDir:       sp,
		WalletDir:       wf.String(),
		DeviceId:        id,
		HighwayAddress:  viper.GetString("highway.address"),
		HighwayPort:     viper.GetInt("highway.port"),
		HighwayNetwork:  viper.GetString("highway.network"),
		LibP2PLowWater:  viper.GetInt("libp2p.lowWater"),
		LibP2PHighWater: viper.GetInt("libp2p.highWater"),
		LibP2PRendevouz: viper.GetString("libp2p.rendevouz"),
	}
	config.Save()
	instance = config
	return config, nil
}

// Arch returns the current architecture.
func Arch() string {
	return runtime.GOARCH
}

// ID returns the device ID.
func ID() string {
	return instance.DeviceId
}

// IsMobile returns true if the current platform is ANY mobile platform.
func IsMobile() bool {
	return runtime.GOOS == "android" || runtime.GOOS == "ios"
}

// IsDesktop returns true if the current platform is ANY desktop platform.
func IsDesktop() bool {
	return runtime.GOOS == "windows" || runtime.GOOS == "linux" || runtime.GOOS == "darwin"
}

// IsAndroid returns true if the current platform is android.
func IsAndroid() bool {
	return runtime.GOOS == "android"
}

// IsIOS returns true if the current platform is iOS.
func IsIOS() bool {
	return runtime.GOOS == "ios"
}

// IsWindows returns true if the current platform is windows.
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsLinux returns true if the current platform is linux.
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsMacOS returns true if the current platform is macOS.
func IsMacOS() bool {
	return runtime.GOOS == "darwin"
}

// Platform returns formatted GOOS for Text format.
// Returns: ["MacOS", "Windows", "Linux", "Android", "iOS"]
func Platform() string {
	switch runtime.GOOS {
	case "windows":
		return "Windows"
	case "darwin":
		return "MacOS"
	case "linux":
		return "Linux"
	case "android":
		return "Android"
	case "ios":
		return "iOS"
	default:
		return "Unknown"
	}
}
