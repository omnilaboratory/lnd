package bean

// Message struct
type Message struct {
	Type    int32  		 `json:"type"`
	Data    interface{}  `json:"data"`
}

type MsgType int

type ReplyMessage struct {
	Type   MsgType `json:"type"`
	Status bool         `json:"status"`
	From   string       `json:"from"`
	To     string       `json:"to"`
	Result interface{}  `json:"result"`
}

type Mnemonic struct {
	Mnemonic string `json:"mnemonic"`
}

type OpenChannelInfo struct {
	FundingPubKey      string `json:"funding_pubkey"`
	FunderAddressIndex int    `json:"funder_address_index"`
	IsPrivate          bool   `json:"is_private"`
}

type AcceptChannelInfo struct {
	TemporaryChannelId string `json:"temporary_channel_id"`
	FundingPubKey      string `json:"funding_pubkey"`
	FundeeAddressIndex int    `json:"fundee_address_index"`
	Approval           bool   `json:"approval"`
}