// Code generated by gen.go. DO NOT EDIT.
// To avoid your changes from being overwritten, delete this header.

package centres

import (
	"bnk.to/core/tools/services"

	pb "bnk.to/core/api/v1/centres"
)

// Server is the implementation of CentresService.
type Server struct {
	pb.UnsafeCentresServiceServer
	services.Common
}

// NewServer creates a new Server.
func NewServer(common services.Common) *Server {
	return &Server{
		Common: common,
	}
}
