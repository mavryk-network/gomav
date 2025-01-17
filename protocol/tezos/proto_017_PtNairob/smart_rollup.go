package proto_017_PtNairob

import (
	mv "github.com/mavryk-network/gomav/v2"
	"github.com/mavryk-network/gomav/v2/encoding"
	"github.com/mavryk-network/gomav/v2/protocol/core"
	"github.com/mavryk-network/gomav/v2/protocol/tezos/proto_016_PtMumbai"
)

type SmartRollupOriginate = proto_016_PtMumbai.SmartRollupOriginate
type SmartRollupAddMessages = proto_016_PtMumbai.SmartRollupAddMessages
type SmartRollupCement = proto_016_PtMumbai.SmartRollupCement
type SmartRollupPublish = proto_016_PtMumbai.SmartRollupPublish
type SmartRollupRefute = proto_016_PtMumbai.SmartRollupRefute
type SmartRollupTimeout = proto_016_PtMumbai.SmartRollupTimeout
type SmartRollupExecuteOutboxMessage = proto_016_PtMumbai.SmartRollupExecuteOutboxMessage
type SmartRollupRecoverBond = proto_016_PtMumbai.SmartRollupRecoverBond

type SmartRollupOriginateContentsAndResult = proto_016_PtMumbai.SmartRollupOriginateContentsAndResult
type SmartRollupAddMessagesContentsAndResult = proto_016_PtMumbai.SmartRollupAddMessagesContentsAndResult
type SmartRollupPublishContentsAndResult = proto_016_PtMumbai.SmartRollupPublishContentsAndResult
type SmartRollupRefuteContentsAndResult = proto_016_PtMumbai.SmartRollupRefuteContentsAndResult
type SmartRollupTimeoutContentsAndResult = proto_016_PtMumbai.SmartRollupTimeoutContentsAndResult
type SmartRollupExecuteOutboxMessageContentsAndResult = proto_016_PtMumbai.SmartRollupExecuteOutboxMessageContentsAndResult
type SmartRollupRecoverBondContentsAndResult = proto_016_PtMumbai.SmartRollupRecoverBondContentsAndResult

type SmartRollupCementResultContents struct {
	ConsumedMilligas mv.BigUint                    `json:"consumed_milligas"`
	InboxLevel       int32                         `json:"inbox_level"`
	CommitmentHash   *mv.SmartRollupCommitmentHash `json:"commitment_hash"`
}

func (r *SmartRollupCementResultContents) GetConsumedMilligas() mv.BigUint { return r.ConsumedMilligas }

type SmartRollupCementResult interface {
	core.ManagerOperationResult
}

func init() {
	encoding.RegisterEnum(&encoding.Enum[SmartRollupCementResult]{
		Variants: encoding.Variants[SmartRollupCementResult]{
			0: (*core.OperationResultApplied[*SmartRollupCementResultContents])(nil),
			1: (*core.OperationResultFailed)(nil),
			2: (*core.OperationResultSkipped)(nil),
			3: (*core.OperationResultBacktracked[*SmartRollupCementResultContents])(nil),
		},
	})
}

//json:kind=OperationKind()
type SmartRollupCementContentsAndResult struct {
	SmartRollupCement
	Metadata proto_016_PtMumbai.ManagerMetadata[SmartRollupCementResult] `json:"metadata"`
}

func (*SmartRollupCementContentsAndResult) OperationContentsAndResult() {}
func (op *SmartRollupCementContentsAndResult) GetMetadata() any {
	return &op.Metadata
}
