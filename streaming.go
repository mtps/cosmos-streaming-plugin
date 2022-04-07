package main

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
	"sync"
)

// IntermediateWriter is used so that we do not need to update the underlying io.Writer inside the StoreKVPairWriteListener
// everytime we begin writing to Kafka topic(s)
type IntermediateWriter struct {
	outChan chan<- []byte
}

// NewIntermediateWriter create an instance of an intermediateWriter that sends to the provided channel
func NewIntermediateWriter(outChan chan<- []byte) *IntermediateWriter {
	return &IntermediateWriter{
		outChan: outChan,
	}
}

// Write satisfies io.Writer
func (iw *IntermediateWriter) Write(b []byte) (int, error) {
	iw.outChan <- b
	return len(b), nil
}

type exampleStreamingService struct {
	keys      []types.StoreKey
	listeners map[types.StoreKey][]types.WriteListener
}

var _ baseapp.StreamingService = exampleStreamingService{}

func NewExampleStreamingService(storeKeys []types.StoreKey, c codec.BinaryCodec) (*exampleStreamingService, error) {
	listenChan := make(chan []byte)
	iw := NewIntermediateWriter(listenChan)
	listener := types.NewStoreKVPairWriteListener(iw, c)
	listeners := make(map[types.StoreKey][]types.WriteListener, len(storeKeys))
	// in this case, we are using the same listener for each Store
	for _, key := range storeKeys {
		listeners[key] = append(listeners[key], listener)
	}
	return &exampleStreamingService{keys: storeKeys, listeners: listeners}, nil
}

func (e exampleStreamingService) Stream(wg *sync.WaitGroup) error {
	panic("implement me")
}

func (e exampleStreamingService) Listeners() map[types.StoreKey][]types.WriteListener {
	return e.listeners
}

func (e exampleStreamingService) ListenBeginBlock(ctx sdk.Context, req abci.RequestBeginBlock, res abci.ResponseBeginBlock) error {
	panic("implement me")
}

func (e exampleStreamingService) ListenEndBlock(ctx sdk.Context, req abci.RequestEndBlock, res abci.ResponseEndBlock) error {
	panic("implement me")
}

func (e exampleStreamingService) ListenDeliverTx(ctx sdk.Context, req abci.RequestDeliverTx, res abci.ResponseDeliverTx) error {
	panic("implement me")
}

func (e exampleStreamingService) Close() error {
	panic("implement me")
}
