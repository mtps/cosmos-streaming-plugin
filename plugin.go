package main

import (
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/plugin"
	serverTypes "github.com/cosmos/cosmos-sdk/server/types"
	"github.com/cosmos/cosmos-sdk/store/types"
	"sync"
)

var _ plugin.Plugin = examplePlugin{}

// Plugin is the base interface for all kinds of cosmos-sdk plugins
// It will be included in interfaces of different Plugins
// --
// type Plugin interface {
//    // Name should return unique name of the plugin
//    Name() string
//
//    // Version returns current version of the plugin
//    Version() string
//
//    // Init is called once when the Plugin is being loaded
//    // The plugin is passed the AppOptions for configuration
//    // A plugin will not necessarily have a functional Init
//    Init(env serverTypes.AppOptions) error
//
//    // Closer interface for shutting down the plugin process
//    io.Closer
// }

// StateStreamingPlugin interface for plugins that load a baseapp.StreamingService implementation from a plugin onto a baseapp.BaseApp
// --
// type StateStreamingPlugin interface {
//    // Register configures and registers the plugin streaming service with the BaseApp
//    Register(bApp *baseapp.BaseApp, marshaller codec.BinaryCodec, keys map[string]*types.KVStoreKey) error
//
//    // Start starts the background streaming process of the plugin streaming service
//    Start(wg *sync.WaitGroup) error
//
//    // Plugin is the base Plugin interface
//    Plugin
//}

type examplePlugin struct {
}

// Name from Plugin interface.
func (e examplePlugin) Name() string {
	panic("todo")
}

// Version from Plugin interface.
func (e examplePlugin) Version() string {
	panic("todo")
}

// Init from Plugin interface.
func (e examplePlugin) Init(env serverTypes.AppOptions) error {
	panic("todo")
}

// Close from io.Closer interface via Plugin interface.
func (e examplePlugin) Close() error {
	panic("todo")
}

// Register from StateStreamPlugin interface.
func (e examplePlugin) Register(bApp *baseapp.BaseApp, marshaller codec.BinaryCodec, keys map[string]*types.KVStoreKey) error {
	exposeStoreKeys := make([]types.StoreKey, 0, len(keys))
	for _, storeKey := range keys {
		exposeStoreKeys = append(exposeStoreKeys, storeKey)
	}
	service, err := NewExampleStreamingService(exposeStoreKeys, marshaller)
	if err != nil {
		return err
	}
	bApp.SetStreamingService(service)
	return nil
}

// Start from StateStreamPlugin interface.
func (e examplePlugin) Start(wg *sync.WaitGroup) error {
	panic("todo")
}
