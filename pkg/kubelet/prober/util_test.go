/*
Copyright 2022 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package prober

import (
	"testing"
	"time"

	v1 "k8s.io/api/core/v1"
	utilpointer "k8s.io/utils/pointer"
)

func TestProbeTimeDuration(t *testing.T) {
	tests := []struct {
		name             string
		probe            v1.Probe
		expectedDuration time.Duration
	}{{
		name: "only seconds",
		probe: v1.Probe{
			PeriodSeconds: 1,
		},
		expectedDuration: 1 * time.Second,
	}, {
		name: "zero milliseconds",
		probe: v1.Probe{
			PeriodSeconds:      4,
			PeriodMilliseconds: utilpointer.Int32(0),
		},
		expectedDuration: 4 * time.Second,
	}, {
		name: "seconds and positive milliseconds",
		probe: v1.Probe{
			PeriodSeconds:      1,
			PeriodMilliseconds: utilpointer.Int32(900),
		},
		expectedDuration: 1*time.Second + 900*time.Millisecond,
	}, {
		name: "seconds and negative milliseconds",
		probe: v1.Probe{
			PeriodSeconds:      1,
			PeriodMilliseconds: utilpointer.Int32(-900),
		},
		expectedDuration: 100 * time.Millisecond,
	}}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			duration := GetProbeTimeDuration(tt.probe.PeriodSeconds, tt.probe.PeriodMilliseconds)
			if duration != tt.expectedDuration {
				t.Errorf("incorrection duration, wanted: %v, got: %v", tt.expectedDuration, duration)
			}
		})
	}
}
