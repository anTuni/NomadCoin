package p2p

type MessageKind int

type Message struct {
	Kind    MessageKind
	Payload []byte
}

const (
	MessageNewestBlock MessageKind = iota
	MessageAllBlocksRequest
	MessageAllBlocksResponse
)
