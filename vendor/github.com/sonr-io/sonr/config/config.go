package config

import "github.com/spf13/viper"

type SonrConfig struct {
	HighwayAddress       string   `json:"highway_address"`
	HighwayPort          int      `json:"highway_port"`
	HighwayNetwork       string   `json:"highway_network"`
	HighwayDID           string   `json:"highway_did"`
	IPFSPort             int      `json:"ipfs_port"`
	IPFSPath             string   `json:"ipfs_path"`
	LibP2PLowWater       int      `json:"libp2p_low_water"`
	LibP2PHighWater      int      `json:"libp2p_high_water"`
	LibP2PRendevouz      string   `json:"libp2p_rendevouz"`
	LibP2PBootstrapPeers []string `json:"libp2p_bootstrap_peers"`
	HomeDir              string   `json:"home_dir"`
	CacheDir             string   `json:"cache_dir"`
	ConfigDir            string   `json:"config_dir"`
	WalletDir            string   `json:"wallet_dir"`
	DeviceId             string   `json:"device_id"`
	PublicIP             string   `json:"public_ip"`
	PrivateIP            string   `json:"private_ip"`
	AccountName          string   `json:"account_name"`
}

func (sc *SonrConfig) Save() (*SonrConfig, error) {
	viper.Set("highway.address", sc.HighwayAddress)
	viper.Set("highway.port", sc.HighwayPort)
	viper.Set("highway.network", sc.HighwayNetwork)
	viper.Set("highway.did", sc.HighwayDID)
	viper.Set("ipfs.port", sc.IPFSPort)
	viper.Set("ipfs.path", sc.IPFSPath)
	viper.Set("libp2p.lowWater", sc.LibP2PLowWater)
	viper.Set("libp2p.highWater", sc.LibP2PHighWater)
	viper.Set("libp2p.rendevouz", sc.LibP2PRendevouz)
	viper.Set("libp2p.bootstrap_peers", sc.LibP2PBootstrapPeers)
	viper.Set("home_dir", sc.HomeDir)
	viper.Set("cache_dir", sc.CacheDir)
	viper.Set("config_dir", sc.ConfigDir)
	viper.Set("wallet_dir", sc.WalletDir)
	viper.Set("device_id", sc.DeviceId)
	viper.Set("public_ip", sc.PublicIP)
	viper.Set("private_ip", sc.PrivateIP)
	viper.Set("account_name", sc.AccountName)
	err := viper.WriteConfig()
	if err != nil {
		return nil, err
	}
	return sc, nil
}

// Return the config dir path as a Folder.
func (sc *SonrConfig) ConfigFolder() Folder {
	return Folder(sc.ConfigDir)
}

// Return the home dir path as a Folder.
func (sc *SonrConfig) HomeFolder() Folder {
	return Folder(sc.HomeDir)
}

// Return the cache dir path as a Folder.
func (sc *SonrConfig) CacheFolder() Folder {
	return Folder(sc.CacheDir)
}

// Create or return the wallet directory as a Folder.
func (sc *SonrConfig) WalletFolder() Folder {
	return Folder(sc.WalletDir)
}
