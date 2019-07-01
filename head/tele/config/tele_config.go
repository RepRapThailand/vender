// Separate package to for hardware/evend related config structure.
// Ugly workaround to import cycles.
package tele_config

type Config struct { //nolint:maligned
	Enabled           bool   `hcl:"enable"`
	VmId              int    `hcl:"vm_id"`
	LogDebug          bool   `hcl:"log_debug"`
	KeepaliveSec      int    `hcl:"keepalive_sec"`
	MqttBroker        string `hcl:"mqtt_broker"`
	MqttLogDebug      bool   `hcl:"mqtt_log_debug"`
	MqttPassword      string `hcl:"mqtt_password"` // secret
	PersistPath       string `hcl:"persist_path"`
	NetworkTimeoutSec int    `hcl:"network_timeout_sec"`
	StateIntervalSec  int    `hcl:"state_interval_sec"`
	TlsCaFile         string `hcl:"tls_ca_file"`
	TlsPsk            string `hcl:"tls_psk"` // secret
}
