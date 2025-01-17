package proto_015_PtLimaPt

import (
	"math/big"
	"strconv"

	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/core/expression"
)

//json:kind=OperationKind()
type ZkRollupOrigination struct {
	ManagerOperation
	PublicParameters mv.Bytes            `mv:"dyn" json:"public_parameters"`
	CircuitsInfo     []*CircuitsInfoElem `mv:"dyn" json:"circuits_info"`
	InitState        mv.Bytes            `mv:"dyn" json:"init_state"`
	NbOps            int32               `json:"nb_ops"`
}

func (*ZkRollupOrigination) OperationKind() string { return "zk_rollup_origination" }

type CircuitsInfoElem struct {
	Value string          `mv:"dyn" json:"value"`
	Tag   CircuitsInfoTag `json:"tag"`
}

type CircuitsInfoTag uint8

func (c CircuitsInfoTag) String() string {
	switch c {
	case CircuitsInfoPublic:
		return "public"
	case CircuitsInfoPrivate:
		return "private"
	case CircuitsInfoFee:
		return "fee"
	default:
		return strconv.FormatInt(int64(c), 10)
	}
}

func (c CircuitsInfoTag) MarshalText() (text []byte, err error) { return []byte(c.String()), nil }

const (
	CircuitsInfoPublic CircuitsInfoTag = iota
	CircuitsInfoPrivate
	CircuitsInfoFee
)

//json:kind=OperationKind()
type ZkRollupPublish struct {
	ManagerOperation
	ZkRollup *mv.ZkRollupAddress `json:"zk_rollup"`
	Op       []*ZkRollupOpElem   `mv:"dyn" json:"op"`
}

func (*ZkRollupPublish) OperationKind() string { return "zk_rollup_publish" }

type ZkRollupOpElem struct {
	Op     ZkRollupOp                 `json:"op"`
	Ticket mv.Option1[ZkRollupTicket] `json:"ticket"`
}

type ZkRollupOp struct {
	OpCode   int32               `json:"op_code"`
	Price    ZkRollupPrice       `json:"price"`
	L1Dst    mv.PublicKeyHash    `json:"l1_dst"`
	RollupID *mv.ZkRollupAddress `json:"rollup_id"`
	Payload  mv.Bytes            `mv:"dyn" json:"payload"`
}

type ZkRollupPrice struct {
	ID     *mv.ScriptExprHash `json:"id"`
	Amount mv.BigInt          `json:"amount"`
}

type ZkRollupTicket struct {
	Contents expression.Expression `json:"contents"`
	Ty       expression.Expression `json:"ty"`
	Ticketer core.ContractID       `json:"ticketer"`
}

//json:kind=OperationKind()
type ZkRollupOriginationContentsAndResult struct {
	ZkRollupOrigination
	Metadata ManagerMetadata[ZkRollupPublishResult] `json:"metadata"`
}

func (*ZkRollupOriginationContentsAndResult) OperationContentsAndResult() {}
func (op *ZkRollupOriginationContentsAndResult) GetMetadata() any {
	return &op.Metadata
}

type ZkRollupPublishResultContents struct {
	BalanceUpdates
	ConsumedMilligas mv.BigUint `json:"consumed_milligas"`
	Size             mv.BigInt  `json:"size"`
}

func (r *ZkRollupPublishResultContents) GetConsumedMilligas() mv.BigUint { return r.ConsumedMilligas }
func (r *ZkRollupPublishResultContents) EstimateStorageSize(constants core.Constants) *big.Int {
	return r.Size.Int()
}

type ZkRollupPublishResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[ZkRollupPublishResult]{
		Variants: encoding.Variants[ZkRollupPublishResult]{
			0: (*core.OperationResultApplied[*ZkRollupPublishResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*ZkRollupPublishResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type ZkRollupPublishContentsAndResult struct {
	ZkRollupPublish
	Metadata ManagerMetadata[ZkRollupPublishResult] `json:"metadata"`
}

func (*ZkRollupPublishContentsAndResult) OperationContentsAndResult() {}
func (op *ZkRollupPublishContentsAndResult) GetMetadata() any {
	return &op.Metadata
}
