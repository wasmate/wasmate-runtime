package wmr

import "github.com/gofiber/fiber/v2"

func (WMR *WMR) NewHTTPServer(fc fiber.Config) *fiber.App {
	// NewHTTPServer creates a new HTTP server instance using the provided fiber.Config.
	WMR.HTTPServer = fiber.New(fc)
	return WMR.HTTPServer
}

func (WMR *WMR) HTTPServerListen(addr string) error {
	// HTTPServerListen starts listening for incoming HTTP requests on the specified address.
	// It returns an error if the server fails to start or encounters any issues during the listening process.
	return WMR.HTTPServer.Listen(addr)
}
