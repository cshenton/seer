package server

import (
	"context"
	"time"

	"github.com/chulabs/seer/seer"
	"github.com/chulabs/seer/stream"
	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/golang/protobuf/ptypes/timestamp"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

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
	err = srv.DB.CreateStream(in.Stream.Name, st)
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
	st, err := srv.DB.GetStream(in.Name)
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
	st, err := srv.DB.GetStream(in.Name)
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
	err = srv.DB.UpdateStream(in.Name, st)
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
	err = srv.DB.DeleteStream(in.Name)
	if err != nil {
		err = status.Error(codes.NotFound, err.Error())
		return nil, err
	}
	return &empty.Empty{}, nil
}

// ListStreams returns a paged set of streams.
func (srv *Server) ListStreams(c context.Context, in *seer.ListStreamsRequest) (s *seer.ListStreamsResponse, err error) {
	lst, err := srv.DB.ListStreams(int(in.PageNumber), int(in.PageSize))
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
	st, err := srv.DB.GetStream(in.Name)
	if err != nil {
		err = status.Error(codes.NotFound, err.Error())
		return nil, err
	}
	times, values, intervals, err := st.Forecast(int(in.N), []float64{0.8, 0.9, 0.95})
	if err != nil {
		err = status.Error(codes.InvalidArgument, err.Error())
		return nil, err
	}

	protoTimes := make([]*timestamp.Timestamp, len(times))
	for i := range times {
		protoTimes[i], _ = ptypes.TimestampProto(times[i])
	}

	protoIntervals := make([]*seer.Interval, len(intervals))
	for i := range intervals {
		protoIntervals[i] = &seer.Interval{
			Probability: intervals[i].Probability,
			LowerBound:  intervals[i].LowerBound,
			UpperBound:  intervals[i].UpperBound,
		}
	}
	f = &seer.Forecast{
		Times:     protoTimes,
		Values:    values,
		Intervals: protoIntervals,
	}
	return f, nil
}
