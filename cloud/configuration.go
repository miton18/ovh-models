package cloud

// GENERATED SDK for cloud API

// A load balancer configuration
type Configuration struct {
	Backends  []Backend  `json:"backends,omitempty"`
	Frontends []Frontend `json:"frontends,omitempty"`
	Version   int64      `json:"version,omitempty"`
}
