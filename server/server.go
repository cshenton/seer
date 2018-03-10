package server

import (
	"context"

	"github.com/golang/protobuf/ptypes"

	"github.com/chulabs/seer/seer"
	"github.com/chulabs/seer/store"
	"github.com/chulabs/seer/stream"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
	st, err := stream.New(
		in.Stream.Name,
		in.Stream.Period,
		in.Stream.Min,
		in.Stream.Max,
		int(in.Stream.Domain),
	)
	if err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		return nil, err
	}
	err = srv.db.CreateStream(in.Stream.Name, st)
	if err != nil {
		err = status.Error(codes.AlreadyExists, err.Error())
		return nil, err
	}
	t, _ := ptypes.TimestampProto(st.Time)
	s = &seer.Stream{
		Name:          st.Config.Name,
		Period:        st.Config.Period,
		LastEventTime: t,
		Domain:        seer.Domain(st.Config.Domain),
		Min:           st.Config.Min,
		Max:           st.Config.Max,
	}
	return s, nil
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
