package mirrornode

type GetTransactionsResponse struct {
	Links        *Links        `json:"links,omitempty"`
	Transactions []Transaction `json:"transactions,omitempty"`
}

// Transactions
//
// Transaction
type Transaction struct {
	BatchKey                 *Key                    `json:"batch_key"`
	Bytes                    *string                 `json:"bytes"`
	ChargedTxFee             *int64                  `json:"charged_tx_fee,omitempty"`
	ConsensusTimestamp       *string                 `json:"consensus_timestamp,omitempty"`
	EntityID                 *string                 `json:"entity_id"`
	MaxCustomFees            []CustomFeeLimit        `json:"max_custom_fees,omitempty"`
	MaxFee                   *string                 `json:"max_fee,omitempty"`
	MemoBase64               *string                 `json:"memo_base64"`
	Name                     *TransactionTypes       `json:"name,omitempty"`
	NftTransfers             []NftTransfer           `json:"nft_transfers,omitempty"`
	Node                     *string                 `json:"node"`
	Nonce                    *int64                  `json:"nonce,omitempty"`
	ParentConsensusTimestamp *string                 `json:"parent_consensus_timestamp"`
	Result                   *string                 `json:"result,omitempty"`
	Scheduled                *bool                   `json:"scheduled,omitempty"`
	StakingRewardTransfers   []StakingRewardTransfer `json:"staking_reward_transfers,omitempty"`
	TokenTransfers           []TokenTransfer         `json:"token_transfers,omitempty"`
	TransactionHash          *string                 `json:"transaction_hash,omitempty"`
	TransactionID            *string                 `json:"transaction_id,omitempty"`
	Transfers                []Transfer              `json:"transfers,omitempty"`
	ValidDurationSeconds     *string                 `json:"valid_duration_seconds,omitempty"`
	ValidStartTimestamp      *string                 `json:"valid_start_timestamp,omitempty"`
}

type Key struct {
	Type *Type   `json:"_type,omitempty"`
	Key  *string `json:"key,omitempty"`
}

// CustomFeeLimit
type CustomFeeLimit struct {
	AccountID           *string `json:"account_id"`
	Amount              *int64  `json:"amount,omitempty"`
	DenominatingTokenID *string `json:"denominating_token_id"`
}

type NftTransfer struct {
	IsApproval        bool    `json:"is_approval"`
	ReceiverAccountID *string `json:"receiver_account_id"`
	SenderAccountID   *string `json:"sender_account_id"`
	SerialNumber      int64   `json:"serial_number"`
	TokenID           *string `json:"token_id"`
}

// StakingRewardTransfers
//
// StakingRewardTransfer, A staking reward transfer
type StakingRewardTransfer struct {
	Account *string `json:"account"`
	// The number of tinybars awarded
	Amount int64 `json:"amount"`
}

type TokenTransfer struct {
	Account    *string `json:"account"`
	Amount     int64   `json:"amount"`
	IsApproval *bool   `json:"is_approval,omitempty"`
	TokenID    *string `json:"token_id"`
}

type Transfer struct {
	Account    *string `json:"account"`
	Amount     int64   `json:"amount"`
	IsApproval *bool   `json:"is_approval,omitempty"`
}

type Type string

const (
	EcdsaSecp256K1  Type = "ECDSA_SECP256K1"
	Ed25519         Type = "ED25519"
	ProtobufEncoded Type = "ProtobufEncoded"
)

type TransactionTypes string

const (
	Atomicbatch            TransactionTypes = "ATOMICBATCH"
	Consensuscreatetopic   TransactionTypes = "CONSENSUSCREATETOPIC"
	Consensusdeletetopic   TransactionTypes = "CONSENSUSDELETETOPIC"
	Consensussubmitmessage TransactionTypes = "CONSENSUSSUBMITMESSAGE"
	Consensusupdatetopic   TransactionTypes = "CONSENSUSUPDATETOPIC"
	Contractcall           TransactionTypes = "CONTRACTCALL"
	Contractcreateinstance TransactionTypes = "CONTRACTCREATEINSTANCE"
	Contractdeleteinstance TransactionTypes = "CONTRACTDELETEINSTANCE"
	Contractupdateinstance TransactionTypes = "CONTRACTUPDATEINSTANCE"
	Cryptoaddlivehash      TransactionTypes = "CRYPTOADDLIVEHASH"
	Cryptoapproveallowance TransactionTypes = "CRYPTOAPPROVEALLOWANCE"
	Cryptocreateaccount    TransactionTypes = "CRYPTOCREATEACCOUNT"
	Cryptodelete           TransactionTypes = "CRYPTODELETE"
	Cryptodeleteallowance  TransactionTypes = "CRYPTODELETEALLOWANCE"
	Cryptodeletelivehash   TransactionTypes = "CRYPTODELETELIVEHASH"
	Cryptotransfer         TransactionTypes = "CRYPTOTRANSFER"
	Cryptoupdateaccount    TransactionTypes = "CRYPTOUPDATEACCOUNT"
	Ethereumtransaction    TransactionTypes = "ETHEREUMTRANSACTION"
	Fileappend             TransactionTypes = "FILEAPPEND"
	Filecreate             TransactionTypes = "FILECREATE"
	Filedelete             TransactionTypes = "FILEDELETE"
	Fileupdate             TransactionTypes = "FILEUPDATE"
	Freeze                 TransactionTypes = "FREEZE"
	Node                   TransactionTypes = "NODE"
	Nodecreate             TransactionTypes = "NODECREATE"
	Nodedelete             TransactionTypes = "NODEDELETE"
	Nodestakeupdate        TransactionTypes = "NODESTAKEUPDATE"
	Nodeupdate             TransactionTypes = "NODEUPDATE"
	Schedulecreate         TransactionTypes = "SCHEDULECREATE"
	Scheduledelete         TransactionTypes = "SCHEDULEDELETE"
	Schedulesign           TransactionTypes = "SCHEDULESIGN"
	Systemdelete           TransactionTypes = "SYSTEMDELETE"
	Systemundelete         TransactionTypes = "SYSTEMUNDELETE"
	Tokenairdrop           TransactionTypes = "TOKENAIRDROP"
	Tokenassociate         TransactionTypes = "TOKENASSOCIATE"
	Tokenburn              TransactionTypes = "TOKENBURN"
	Tokencancelairdrop     TransactionTypes = "TOKENCANCELAIRDROP"
	Tokenclaimairdrop      TransactionTypes = "TOKENCLAIMAIRDROP"
	Tokencreation          TransactionTypes = "TOKENCREATION"
	Tokendeletion          TransactionTypes = "TOKENDELETION"
	Tokendissociate        TransactionTypes = "TOKENDISSOCIATE"
	Tokenfeescheduleupdate TransactionTypes = "TOKENFEESCHEDULEUPDATE"
	Tokenfreeze            TransactionTypes = "TOKENFREEZE"
	Tokengrantkyc          TransactionTypes = "TOKENGRANTKYC"
	Tokenmint              TransactionTypes = "TOKENMINT"
	Tokenpause             TransactionTypes = "TOKENPAUSE"
	Tokenreject            TransactionTypes = "TOKENREJECT"
	Tokenrevokekyc         TransactionTypes = "TOKENREVOKEKYC"
	Tokenunfreeze          TransactionTypes = "TOKENUNFREEZE"
	Tokenunpause           TransactionTypes = "TOKENUNPAUSE"
	Tokenupdate            TransactionTypes = "TOKENUPDATE"
	Tokenupdatenfts        TransactionTypes = "TOKENUPDATENFTS"
	Tokenwipe              TransactionTypes = "TOKENWIPE"
	Uncheckedsubmit        TransactionTypes = "UNCHECKEDSUBMIT"
	Unknown                TransactionTypes = "UNKNOWN"
	Utilprng               TransactionTypes = "UTILPRNG"
)
