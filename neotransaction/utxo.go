package neotransaction

import "github.com/incognito-core-libs/neo-go-sdk/neoutils"

// UTXO Neo交易中的未花费输出
type UTXO struct {
	TxHash neoutils.HASH256
	Index  uint16
	TxOutput
}
