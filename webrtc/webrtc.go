package webrtc

type ICEServer struct {
	URLs       []string `json:"urls"`
	UserName   string   `json:"username,omitempty"`
	Credential string   `json:"credential,omitempty"`
}

type RTCConfig struct {
	LifetimeDuration   string      `json:"lifetimeDuration"`
	IceServers         []ICEServer `json:"iceServers"`
	BlockStatus        string      `json:"blockStatus"`
	IceTransportPolicy string      `json:"iceTransportPolicy"`
}

func generateCredential(sharedSecret string, user string) (string, string) {

}
