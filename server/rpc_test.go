package server_test

import (
	"context"
	"testing"
	"time"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/timestamp"

	"github.com/chulabs/seer/seer"
	"github.com/chulabs/seer/server"
	"github.com/chulabs/seer/stream"
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

}

func TestGetForecast(t *testing.T) {

}
