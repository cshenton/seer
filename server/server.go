package server

import (
	"context"

	"github.com/chulabs/seer/seer"
	google_protobuf "github.com/golang/protobuf/ptypes/empty"
)

// Seer fulfills the protocol buffer's SeerServer interface.
type Seer struct{}

// CreateStream creates the provided stream.
func (w *Seer) CreateStream(c context.Context, in *seer.CreateStreamRequest) (s *seer.Stream, err error) {
	return
}

// GetStream retrieves and returns the requested stream.
func (w *Seer) GetStream(c context.Context, in *seer.GetStreamRequest) (s *seer.Stream, err error) {
	return
}

// UpdateStream applies an adaptive filter update using the provided events.
func (w *Seer) UpdateStream(c context.Context, in *seer.UpdateStreamRequest) (s *seer.Stream, err error) {
	return
}

// DeleteStream removes the requested stream.
func (w *Seer) DeleteStream(c context.Context, in *seer.DeleteStreamRequest) (em *google_protobuf.Empty, err error) {
	return
}

// ListStreams returns a paged set of streams.
func (w *Seer) ListStreams(c context.Context, in *seer.ListStreamsRequest) (s *seer.ListStreamsResponse, err error) {
	return
}

// GetForecast generates a forecast from a stream from its current time.
func (w *Seer) GetForecast(c context.Context, in *seer.GetForecastRequest) (f *seer.Forecast, err error) {
	return
}
