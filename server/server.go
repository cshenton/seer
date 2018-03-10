package server

import (
	"context"

	"github.com/chulabs/seer/seer"
	"github.com/chulabs/seer/store"
	"github.com/golang/protobuf/ptypes/empty"
)

// Server fulfills the protocol buffer's SeerServer interface.
type Server struct {
	db store.StreamStore
}

// New creates a database connection and returns a Server.
func New(path string) (srv *Server) {
	// instantiate database
	// return server struct
	return
}

// CreateStream creates the provided stream.
func (srv *Server) CreateStream(c context.Context, in *seer.CreateStreamRequest) (s *seer.Stream, err error) {
	// stream.New with provided data
	// srv.db.CreateStream
	// make and return message
	return
}

// GetStream retrieves and returns the requested stream.
func (srv *Server) GetStream(c context.Context, in *seer.GetStreamRequest) (s *seer.Stream, err error) {
	// srv.db.GetStream(name)
	// make and return message
	return
}

// UpdateStream applies an adaptive filter update using the provided events.
func (srv *Server) UpdateStream(c context.Context, in *seer.UpdateStreamRequest) (s *seer.Stream, err error) {
	// srv.db.GetStream(name)
	// s.Update(events)
	// srv.db.UpdateStream(name, s)
	// make and return message
	return
}

// DeleteStream removes the requested stream.
func (srv *Server) DeleteStream(c context.Context, in *seer.DeleteStreamRequest) (em *empty.Empty, err error) {
	// srv.db.DeleteStream(name)
	// make and return message
	return
}

// ListStreams returns a paged set of streams.
func (srv *Server) ListStreams(c context.Context, in *seer.ListStreamsRequest) (s *seer.ListStreamsResponse, err error) {
	// srv.db.ListStreams()
	// make and return message
	return
}

// GetForecast generates a forecast from a stream from its current time.
func (srv *Server) GetForecast(c context.Context, in *seer.GetForecastRequest) (f *seer.Forecast, err error) {
	// srv.db.GetStream(name)
	// s.Forecast()
	// make and return message
	return
}
