package byparr

const (
	CmdRequestGet  string = "request.get"
	StatusOK       string = "ok"
	MessageSuccess string = "Success"
)

type Request struct {
	Cmd        string `json:"cmd"`
	URL        string `json:"url"`
	MaxTimeout int    `json:"maxTimeout"`
}

type Response struct {
	Status   string `json:"status"`
	Message  string `json:"message"`
	Solution struct {
		URL     string `json:"url"`
		Status  int    `json:"status"`
		Cookies []struct {
			Domain   string `json:"domain"`
			Expiry   int    `json:"expiry"`
			HTTPOnly bool   `json:"httpOnly"`
			Name     string `json:"name"`
			Path     string `json:"path"`
			SameSite string `json:"sameSite"`
			Secure   bool   `json:"secure"`
			Value    string `json:"value"`
			Size     int    `json:"size"`
			Session  bool   `json:"session"`
			Expires  int64  `json:"expires"`
		} `json:"cookies"`
		UserAgent string `json:"userAgent"`
		Headers   struct {
		} `json:"headers"`
	} `json:"solution"`
	StartTimestamp int64  `json:"startTimestamp"`
	EndTimestamp   int64  `json:"endTimestamp"`
	Version        string `json:"version"`
}
