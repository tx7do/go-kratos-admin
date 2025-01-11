package utils

import (
	"time"

	"google.golang.org/protobuf/types/known/durationpb"
)

func Float64ToDurationpb(d float64) *durationpb.Duration {
	duration := time.Duration(d * float64(time.Second))
	return durationpb.New(duration)
}

func SecondToDurationpb(seconds *float64) *durationpb.Duration {
	if seconds == nil {
		return nil
	}
	return durationpb.New(time.Duration(*seconds) * time.Second)
}

func DurationpbSecond(duration *durationpb.Duration) *float64 {
	if duration == nil {
		return nil
	}
	seconds := duration.AsDuration().Seconds()
	secondsInt64 := seconds
	return &secondsInt64
}
