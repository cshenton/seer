package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/cshenton/seer/seer"
	"github.com/cshenton/seer/server"
	"github.com/cshenton/seer/stream"
)

func setUp(t *testing.T) (srv *server.Server) {
	srv, err := server.New(testPath(t))
	if err != nil {
		t.Fatal("unexpected error in server.New:", err)
	}

	names := []string{"sales", "visits", "usage"}

	for _, n := range names {
		s, _ := stream.New(n, 3600, 0, 0, 0)
		err := srv.DB.CreateStream(n, s)
		if err != nil {
			t.Fatal("unexpected error in CreateStream:", err)
		}
	}

	return srv
}

func TestCreateStream(t *testing.T) {
	srv := setUp(t)

	tt := []struct {
		name   string
		period float64
		min    float64
		max    float64
		domain int
	}{
		{"simple hourly", 3600, 0, 0, 0},
		{"positive daily", 86400, 0, 0, 1},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			in := &seer.CreateStreamRequest{
				Stream: &seer.Stream{
					Name:   tc.name,
					Period: tc.period,
					Min:    tc.min,
					Max:    tc.max,
					Domain: seer.Domain(tc.domain),
				},
			}
			s, err := srv.CreateStream(context.Background(), in)
			if err != nil {
				t.Fatal("unexpected error in CreateStream:", err)
			}
			if s.Name != tc.name {
				t.Errorf("expected name %v, but got %v", tc.name, s.Name)
			}
		})
	}

}

func TestCreateStreamErrs(t *testing.T) {
	srv := setUp(t)

	tt := []string{"sales", "s"}

	for _, name := range tt {
		t.Run(name, func(t *testing.T) {
			in := &seer.CreateStreamRequest{
				Stream: &seer.Stream{
					Name:   name,
					Period: 86400,
				},
			}
			s, err := srv.CreateStream(context.Background(), in)
			if err == nil {
				t.Error("expected error, but it was nil")
			}
			if s != nil {
				t.Error("expected nil response, but got", s)
			}

		})
	}
}

func TestGetStream(t *testing.T) {
	srv := setUp(t)

	tt := []string{"sales", "visits", "usage"}

	for _, name := range tt {
		t.Run(name, func(t *testing.T) {
			in := &seer.GetStreamRequest{
				Name: name,
			}
			s, err := srv.GetStream(context.Background(), in)
			if err != nil {
				t.Fatal("unexpected error in GetStream:", err)
			}
			if s.Name != name {
				t.Errorf("expected name %v, but got %v", name, s.Name)
			}
		})
	}
}

func TestGetStreamErrs(t *testing.T) {
	srv := setUp(t)

	name := "notastream"
	in := &seer.GetStreamRequest{
		Name: name,
	}
	s, err := srv.GetStream(context.Background(), in)
	if err == nil {
		t.Fatal("expected error, but it was nil")
	}
	if s != nil {
		t.Error("expected nil response, but got", s)
	}
}

func TestDeleteStream(t *testing.T) {
	srv := setUp(t)

	tt := []string{"sales", "visits", "usage"}

	for _, name := range tt {
		t.Run(name, func(t *testing.T) {
			in := &seer.DeleteStreamRequest{
				Name: name,
			}
			_, err := srv.DeleteStream(context.Background(), in)
			if err != nil {
				t.Error("unexpected error in DeleteStream:", err)
			}

			gin := &seer.GetStreamRequest{
				Name: name,
			}
			_, err = srv.GetStream(context.Background(), gin)
			if err == nil {
				t.Error("expected error on get after delete, but it was nil")
			}
		})
	}
}

func TestDeleteStreamErrs(t *testing.T) {
	srv := setUp(t)

	name := "notastream"

	in := &seer.DeleteStreamRequest{
		Name: name,
	}
	_, err := srv.DeleteStream(context.Background(), in)
	if err == nil {
		t.Error("expected error, but it was nil")
	}
}

func TestUpdateStream(t *testing.T) {
	srv := setUp(t)

	tt := []struct {
		name   string
		values []float64
		times  []time.Time
	}{
		{"sales", []float64{3.14}, []time.Time{time.Now()}},
		{
			"visits", []float64{3.14, 4.43},
			[]time.Time{
				time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC),
				time.Date(2016, 1, 1, 1, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			times := make([]*timestamp.Timestamp, len(tc.times))
			for i := range times {
				times[i], _ = ptypes.TimestampProto(tc.times[i])
			}
			in := &seer.UpdateStreamRequest{
				Name: tc.name,
				Event: &seer.Event{
					Values: tc.values,
					Times:  times,
				},
			}
			s, err := srv.UpdateStream(context.Background(), in)
			if err != nil {
				t.Error("unpexpected error in UpdateStream:", err)
			}
			if s.Name != tc.name {
				t.Errorf("expected name %v, but got %v", tc.name, s.Name)
			}
		})
	}
}

func TestUpdateStreamErrs(t *testing.T) {
	srv := setUp(t)

	tt := []struct {
		name   string
		values []float64
		times  []time.Time
	}{
		{"notastream", []float64{3.14}, []time.Time{time.Now()}},
		{"visits", []float64{3.14, 4.43}, []time.Time{time.Now()}},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			times := make([]*timestamp.Timestamp, len(tc.times))
			for i := range times {
				times[i], _ = ptypes.TimestampProto(tc.times[i])
			}
			in := &seer.UpdateStreamRequest{
				Name: tc.name,
				Event: &seer.Event{
					Values: tc.values,
					Times:  times,
				},
			}
			s, err := srv.UpdateStream(context.Background(), in)
			if err == nil {
				t.Error("expected error but it was nil")
			}
			if s != nil {
				t.Error("expected nil stream, but it was", s)
			}
		})
	}
}

func TestListStreams(t *testing.T) {
	srv := setUp(t)

	var pageNum int32 = 1
	var pageSize int32 = 2

	in := &seer.ListStreamsRequest{
		PageSize:   pageSize,
		PageNumber: pageNum,
	}
	s, err := srv.ListStreams(context.Background(), in)
	if err != nil {
		t.Error("unexpected error in ListStreams:", err)
	}
	if len(s.Streams) != int(pageSize) {
		t.Errorf("expected %v streams, but got %v", pageSize, len(s.Streams))
	}
	if s.Streams[0].Name != "sales" {
		t.Errorf("expected first name %v, but got %v", "sales", s.Streams[0].Name)
	}
	if s.Streams[1].Name != "usage" {
		t.Errorf("expected first name %v, but got %v", "usage", s.Streams[1].Name)
	}
}

func TestListStreamsErrs(t *testing.T) {
	srv := setUp(t)

	var pageNum int32 = 10
	var pageSize int32 = 5

	in := &seer.ListStreamsRequest{
		PageSize:   pageSize,
		PageNumber: pageNum,
	}
	s, err := srv.ListStreams(context.Background(), in)
	if err == nil {
		t.Error("expected error, but it was nil")
	}
	if s != nil {
		t.Error("expected nil streamlist, but got ", s)
	}
}

func TestGetForecast(t *testing.T) {
	srv := setUp(t)

	tt := []struct {
		name string
		n    int
	}{
		{"visits", 1},
		{"usage", 35},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			tm, _ := ptypes.TimestampProto(time.Now())
			uin := &seer.UpdateStreamRequest{
				Name: tc.name,
				Event: &seer.Event{
					Values: []float64{3.15},
					Times:  []*timestamp.Timestamp{tm},
				},
			}
			_, err := srv.UpdateStream(context.Background(), uin)
			if err != nil {
				t.Fatal("unexpected error in UpdateStream:", err)
			}
			in := &seer.GetForecastRequest{
				Name: tc.name,
				N:    int32(tc.n),
			}
			f, err := srv.GetForecast(context.Background(), in)
			if err != nil {
				t.Fatal("unexpected error in GetForecast:", err)
			}
			if len(f.Intervals) != 3 {
				t.Errorf("expected %v intervals, but got %v", 3, len(f.Intervals))
			}
			if len(f.Times) != tc.n {
				t.Errorf("expected %v times, but got %v", tc.n, len(f.Times))
			}
			if len(f.Values) != tc.n {
				t.Errorf("expected %v values, but got %v", tc.n, len(f.Values))
			}
		})
	}
}

func TestGetForecastErrs(t *testing.T) {
	srv := setUp(t)

	tt := []struct {
		name string
		n    int32
	}{
		{"notastream", 10},
		{"sales", -5},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			in := &seer.GetForecastRequest{
				Name: tc.name,
				N:    tc.n,
			}
			f, err := srv.GetForecast(context.Background(), in)
			if err == nil {
				t.Error("expected error, but it was nil")
			}
			if f != nil {
				t.Error("expected nil forecast, but got", f)
			}
		})
	}
}
