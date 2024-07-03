package shutdown

import (
	"context"
	"testing"
	"time"
)

func TestShutdown_SetReady(t *testing.T) {
	type fields struct {
		ready      int
		inProgress bool
	}
	type args struct {
		ready bool
	}

	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "test setting ready to false when already false",
			fields: fields{ready: 0, inProgress: false},
			args:   args{false},
			want:   false,
		},
		{
			name:   "test setting ready to false when was true and in progress",
			fields: fields{ready: 0, inProgress: true},
			args:   args{false},
			want:   false,
		},
		{
			name:   "test setting ready to true when was false and in progress",
			fields: fields{ready: 1, inProgress: true},
			args:   args{true},
			want:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, cancel := context.WithCancel(context.Background())
			s := NewShutdown(cancel)
			s.ready.Store(int32(tt.fields.ready))
			s.inProgress.Store(tt.fields.inProgress)
			r := s.SetReady(tt.args.ready)

			if tt.want == false {
				if r {
					t.Errorf("SetReady() = %v, want %v", r, tt.want)
				}
			} else if tt.want == true && tt.fields.inProgress {
				wasReady := false

				select {
				// max timeout for test to fail
				case <-time.After(5 * time.Second):
				case <-s.readyChan:
					wasReady = true
				}
				if !wasReady {
					t.Errorf("SetReady() = %v, want %v", wasReady, tt.want)
				}
			}
		})
	}
}
