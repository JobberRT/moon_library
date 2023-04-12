package model

// DNS template structs
type (
	DNSTemplate struct {
		ID                 int `gorm:"primaryKey; autoIncrement"`
		CreatedAt          int
		UpdatedAt          int
		Name               string `gorm:"unique; not null"`
		QueryStrategy      string `gorm:"not null"`
		DisableDNSCache    bool   `gorm:"not null"`
		DisableDNSFallback bool   `gorm:"not null"`
		// Servers json string of dnsServers
		Servers string `gorm:"not null"`
		// Hosts json string of dnHosts
		Hosts string
		Nodes []*Node
	}
	dnsServers struct {
		Address   string   `json:"address"`
		Port      uint16   `json:"port"`
		Domains   []string `json:"domains"`
		ExpectIPs []string `json:"expect_ips"`
	}
	// dnsHosts KV, Domain: [ IPs ]
	dnHosts struct {
		Domain string   `json:"domain"`
		IPs    []string `json:"ips"`
	}
)

// Fallback template structs
type (
	FallbackTemplate struct {
		ID        int `gorm:"primaryKey; autoIncrement"`
		CreatedAt int
		UpdatedAt int
		Name      string `gorm:"unique; not null"`
		// Fallback json string of fallback
		Fallback string `gorm:"not null"`
	}
	fallback struct {
		Dest string   `json:"dest"`
		XVer uint8    `json:"xver"`
		SNI  string   `json:"sni"`
		ALPN []string `json:"alpn"`
		Path string   `json:"path"`
	}
)

// Inbound template structs
type (
	InboundTemplate struct {
		ID        int `gorm:"primaryKey; autoIncrement"`
		CreatedAt int
		UpdatedAt int
		Name      string `gorm:"unique; not null"`
		ListenOn  string `gorm:"not null"`
		Protocol  string `gorm:"not null"`
		Sniffing  bool   `gorm:"not null"`
		// InboundSetting json string of one of the InboundXXXs
		InboundSetting string `gorm:"not null"`
		Nodes          []*Node
	}
	inboundVLess struct {
		Flow       string `json:"flow"`
		FallbackID int    `json:"fallback"`
	}
	inboundVMess struct {
		DisableInsecure bool `json:"disableInsecure"`
	}
	inboundTrojan struct {
		FallbackTemplateID int `json:"fallback"`
	}
	inboundShadowsocks struct {
		Network string `json:"network"`
		Method  string `json:"method"`
	}
)

// Outbound template structs
type (
	OutboundTemplate struct {
		ID          int `gorm:"primaryKey; autoIncrement"`
		CreatedAt   int
		UpdatedAt   int
		Name        string `gorm:"unique; not null"`
		SendThrough string `gorm:"not null"`
		Protocol    string `gorm:"not null"`
		// OutboundSetting json string of one of the OutboundXXXs
		OutboundSetting string `gorm:"not null"`
		Nodes           []*Node
	}
	outboundFreedom struct {
		DomainStrategy string `json:"domainStrategy"`
	}
	outboundDNS struct {
		Network string `json:"network"`
		Address string `json:"address"`
		Port    uint16 `json:"port"`
	}
	outboundBlackhole struct {
		Response string `json:"response"`
	}
)

// Routing template structs
type (
	RoutingTemplate struct {
		ID             int `gorm:"primaryKey; autoIncrement"`
		CreatedAt      int
		UpdatedAt      int
		Name           string `gorm:"unique; not null"`
		DomainStrategy string `gorm:"not null"`
		DomainMatcher  string `gorm:"not null"`
		// Rule json string of rule
		Rule  string `gorm:"not null"`
		Nodes []*Node
	}
	rule struct {
		OutboundTag string `json:"outboundTag"`
		Network     string `json:"network"`
		DestPort    uint16 `json:"destPort"`
		SrcPort     uint16 `json:"srcPort"`
		Domains     string `json:"domains"`
		IPs         string `json:"ips"`
		Protocols   string `json:"protocols"`
	}
)

// StreamSetting template structs
type (
	StreamSettingTemplate struct {
		ID        int `gorm:"primaryKey; autoIncrement"`
		CreatedAt int
		UpdatedAt int
		Name      string `gorm:"unique;not null" json:"name"`
		Network   string `gorm:"not null" json:"network"`
		Security  string `gorm:"not null" json:"tls"`
		// NetworkSetting json of one of the NetworkXXX
		NetworkSetting string
		// NetworkSetting json of one of the SecurityXXX
		SecuritySetting string
		Nodes           []*Node
	}
	networkTCP struct {
		AcceptProxyProtocol bool   `json:"acceptProxyProtocol"`
		HeaderType          string `json:"headerType"`
	}
	networkKCP struct {
		Seed             string `json:"seed"`
		UplinkCapacity   uint16 `json:"uplinkCapacity"`
		DownLinkCapacity uint16 `json:"downlink_capacity"`
		ReadBufferSize   uint16 `json:"readBufferSize"`
		WriteBufferSize  uint16 `json:"writeBufferSize"`
		Header           string `json:"header"`
		Congestion       bool   `json:"congestion"`
	}
	networkWebsocket struct {
		Path                string `json:"path"`
		AcceptProxyProtocol bool   `json:"acceptProxyProtocol"`
	}
	networkHTTP struct {
		Path   string   `json:"path"`
		Method string   `json:"method"`
		Hosts  []string `json:"hosts"`
	}
	networkQUIC struct {
		Security string `json:"security"`
		Key      string `json:"key"`
		Header   string `json:"header"`
	}
	networkGRPC struct {
		ServiceName string `json:"serviceName"`
		MultiMode   bool   `json:"multiMode"`
	}
	securityTLS struct {
		Fingerprint      string   `json:"fingerprint"`
		RejectUnknownSNI string   `json:"rejectUnknownSNI"`
		AllowInsecure    string   `json:"allowInsecure"`
		ALPNs            []string ` json:"alpns"`
	}
	securityReality struct {
		Fingerprint string   `json:"fingerprint"`
		Dest        string   `json:"dest"`
		ServerNames []string `json:"serverNames"`
	}
)
