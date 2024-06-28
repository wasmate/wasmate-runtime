package wmr

import (
	"context"
	"sync"
	"sync/atomic"

	"github.com/gofiber/fiber/v2"
	"github.com/wasmate/wasmate-runtime/pkg/confer"
)

// WMR is the core data structure of WMR and is used to extend and integrate other structures and behaviors.
type WMR struct {
	Ctx        context.Context // Global context of WMR
	OfflineCtx context.Context // OfflineCtx is used to propagate WMR offline signals
	Confer     *confer.Confer  // Confer entity of WMR

	ShutdownWG     sync.WaitGroup // Wait for all cleanup operations to complete on WMR shutdown
	InFlowProtocol atomic.Value   // Inbound traffic protocol

	HTTPServer *fiber.App

	WMRPlugin *WMRPlugins

	sync.RWMutex
}

// NewWMR creates a new instance of WMR with the specified context, offline context, and confer.
// It ensures that only one instance of WMR is created and returns the existing instance if it already exists.
func NewWMR(ctx, offlineCtx context.Context, conf *confer.Confer) *WMR {

	_WMR := &WMR{Ctx: ctx, Confer: conf, OfflineCtx: offlineCtx}

	_WMR.WMRPlugin = NewWMRPlugins()
	return _WMR
}
