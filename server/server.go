package server

import (
	"context"
	"time"

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
	st, err := srv.db.GetStream(in.Name)
	if err != nil {
		err = status.Error(codes.NotFound, err.Error())
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

// UpdateStream applies an adaptive filter update using the provided events.
func (srv *Server) UpdateStream(c context.Context, in *seer.UpdateStreamRequest) (s *seer.Stream, err error) {
	st, err := srv.db.GetStream(in.Name)
	if err != nil {
		err = status.Error(codes.NotFound, err.Error())
		return nil, err
	}
	t := make([]time.Time, len(in.Event.Times))
	for i := range t {
		ts, _ := ptypes.Timestamp(in.Event.Times[i])
		t[i] = ts
	}

	err = st.Update(in.Event.Values, t)
	if err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		return nil, err
	}
	err = srv.db.UpdateStream(in.Name, st)
	if err != nil {
		err = status.Error(codes.NotFound, err.Error())
		return nil, err
	}

	ts, _ := ptypes.TimestampProto(st.Time)
	s = &seer.Stream{
		Name:          st.Config.Name,
		Period:        st.Config.Period,
		LastEventTime: ts,
		Domain:        seer.Domain(st.Config.Domain),
		Min:           st.Config.Min,
		Max:           st.Config.Max,
	}
	return s, nil
}

// DeleteStream removes the requested stream.
func (srv *Server) DeleteStream(c context.Context, in *seer.DeleteStreamRequest) (em *empty.Empty, err error) {
	err = srv.db.DeleteStream(in.Name)
	if err != nil {
		err = status.Error(codes.NotFound, err.Error())
		return nil, err
	}
	return &empty.Empty{}, nil
}

// ListStreams returns a paged set of streams.
func (srv *Server) ListStreams(c context.Context, in *seer.ListStreamsRequest) (s *seer.ListStreamsResponse, err error) {
	lst, err := srv.db.ListStreams(int(in.PageNumber), int(in.PageSize))
	if err != nil {
		err = status.Error(codes.NotFound, err.Error())
		return nil, err
	}
	ls := make([]*seer.Stream, len(lst))
	for i := range lst {
		st := lst[i]
		ts, _ := ptypes.TimestampProto(st.Time)
		ls[i] = &seer.Stream{
			Name:          st.Config.Name,
			Period:        st.Config.Period,
			LastEventTime: ts,
			Domain:        seer.Domain(st.Config.Domain),
			Min:           st.Config.Min,
			Max:           st.Config.Max,
		}
	}
	s = &seer.ListStreamsResponse{
		Streams: ls,
	}
	return s, nil
}

// GetForecast generates a forecast from a stream from its current time.
func (srv *Server) GetForecast(c context.Context, in *seer.GetForecastRequest) (f *seer.Forecast, err error) {
	// srv.db.GetStream(name)
	// s.Forecast()
	// make and return message
	return
}
