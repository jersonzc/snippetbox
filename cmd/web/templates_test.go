package main

import (
	"testing"
	"time"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name   string
		myTime time.Time
		want   string
	}{
		{
			name:   "UTC",
			myTime: time.Date(2022, 3, 17, 10, 15, 0, 0, time.UTC),
			want:   "17 Mar 2022 at 10:15",
		},
		{
			name:   "No date",
			myTime: time.Time{},
			want:   "",
		},
		{
			name:   "CET",
			myTime: time.Date(2022, 3, 17, 10, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want:   "17 Mar 2022 at 09:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			date := humanDate(tt.myTime)
			if date != tt.want {
				t.Errorf("got %s, want %s", date, tt.want)
			}
		})
	}
}
