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

type Transaction struct {
	Data 		string `json:"data"`
	From 		string `json:"from"`
	Gas 		string `json:"gas"`
	GasPrice 	string `json:"gasPrice"`
	Input 		string `json:"input"`
	Nonce 		string `json:"nonce"`
	To 			string `json:"to"`
	Value 		string `json:"value"`
}

type CallInfo struct {
	Message 	string `json:"message"`
	Type 		string `json:"type"`
}

type ApproveTxParams struct {
	Meta 			Meta `json:"meta"`
	Transaction 	Transaction `json:"transaction"`
	Call_info 		[]CallInfo `json:"call_info"`
}

type ApproveTxResponse struct {
	Approved 		bool
	Password 		string
	Transaction 	Transaction
}


type ApproveSignDataParams struct {
	Address 		string `json:"address"`
	Raw_data		string `json:"raw_data"`
	Message			string `json:"message"`
	Hash 			string `json:"hash"`
	Meta 			Meta `json:"meta"`
}

type ApproveSignDataResponse struct {
	Approved 		bool
	Password 		string
}

type ApproveNewAccountParams struct {
	Meta 			Meta `json:"meta"`
}

type ApproveNewAccountResponse struct {
	Approved 		bool
	Password 		string
}

type ApproveListingResponse struct {
	Accounts 		[]ApproveListingAccount
}

type ApproveImportParams struct {
	Meta 			Meta `json:"meta"`
}

type ApproveImportResponse struct {
	Approved 			bool
	Old_Password 		string
	New_Password 		string
}

type ApproveExportParams struct {
	Meta 			Meta `json:"meta"`
	Address 		string `json:"address"`
}

type ApproveExportResponse struct {
	Approved 		bool
}

type ShowErrorParam struct {
	Text 			string
}