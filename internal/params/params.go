package params

type ApproveListingAccount struct {
	Address 		string `json:"address"`
	Type 			string `json:"type"`
	Url 			string `json:"url"`
}

type ApproveListingParams struct {
	Accounts		[]ApproveListingAccount `json:"accounts"`
	Meta 			Meta `json:"meta"`
}

type Meta struct {
	Transport 		string `json:"scheme"`
	Remote 			string `json:"remote"`
	Local 			string `json:"local"`
}

type OnSignerStartupInfo struct {
	Extapi_http    string `json:"extapi_http"`
	Extapi_ipc     string `json:"extapi_ipc"`
	Extapi_version string `json:"extapi_version"`
	Intapi_version string `json:"intapi_version"`
}

type OnSignerStartupParam struct {
	Info *OnSignerStartupInfo `json:"info"`
}

type OnSignerStartupReply struct {
}


type ApproveSignDataParams struct {
	Address 		string `json:"address"`
	Raw_data		string `json:"raw_data"`
	Message			string `json:"message"`
	Hash 			string `json:"hash"`
	Meta 			Meta `json:"meta"`
}

