package wmr

import (
	"context"
	"fmt"
	"sync"

	extism "github.com/extism/go-sdk"
)

// WMRPlugins represents a collection of wasm plugins.
type WMRPlugins struct {
	RWMRutex sync.RWMutex
	Plugins  map[string]*WMRPlugin
}

// WMRPlugin Struct: actor pool
type WMRPlugin struct {
	Name   string
	Plugin extism.Plugin
	Mutex  sync.Mutex
}

func newWMRPlugin(pluginName string, plugin extism.Plugin) *WMRPlugin {
	return &WMRPlugin{Name: pluginName, Plugin: plugin}
}

// NewWMRPlugins creates a new instance of WMRPlugins.
func NewWMRPlugins() (newWMRPlugins *WMRPlugins) {
	newWMRPlugins = &WMRPlugins{
		Plugins: make(map[string]*WMRPlugin),
	}

	return
}

// AddPlugin adds a new wasm plugin to the WMRPlugins collection.
// If the plugin with the given name already exists, it will return an error unless the force flag is set to true.
// If force is true, the existing plugin will be replaced with the new plugin.
func (WMRPlugins *WMRPlugins) AddPlugin(pluginName string, plugin extism.Plugin, force bool) (err error) {
	WMRPlugins.RWMRutex.Lock()
	defer WMRPlugins.RWMRutex.Unlock()

	if _, ok := WMRPlugins.Plugins[pluginName]; !ok {
		WMRp := newWMRPlugin(pluginName, plugin)
		WMRPlugins.Plugins[pluginName] = WMRp
		return nil
	}

	// If plugin is extism, judge if force replace is required
	if !force {
		// If force is not enabled, return an error
		return fmt.Errorf("failed to add the wasm plugin(%s), the plug-in already exists", pluginName)
	}

	// If force is true, then force replace the wasm plugin by name.
	WMRp := newWMRPlugin(pluginName, plugin)
	WMRPlugins.Plugins[pluginName] = WMRp
	return nil
}

// GetPluginByName returns the wasm plugin with the given name from the WMRPlugins collection.
// It acquires a read lock on the RWMRutex to ensure thread safety.
// If the plugin with the given name exists, it returns the plugin and a nil error.
// If the plugin with the given name does not exist, it returns an empty extism.Plugin and an error indicating that the plugin does not exist.
func (WMRPlugins *WMRPlugins) GetPluginByName(pluginName string) (*WMRPlugin, error) {
	WMRPlugins.RWMRutex.RLock()
	defer WMRPlugins.RWMRutex.RUnlock()

	if plugin, ok := WMRPlugins.Plugins[pluginName]; ok {
		return plugin, nil
	}

	return nil, fmt.Errorf("no wasm plugin named %s", pluginName)
}

// Create a new wasm plugin using the extism.NewPlugin function
func (WMRPlugins *WMRPlugins) NewWMRPlugin(ctx context.Context,
	manifest extism.Manifest,
	config extism.PluginConfig,
	functions []extism.HostFunction) (*extism.Plugin, error) {

	return extism.NewPlugin(ctx, manifest, config, nil)
}
