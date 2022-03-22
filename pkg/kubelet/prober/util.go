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

import "time"

// MinProbeTimeDuration sets the minimum value
const MinProbeTimeDuration = 100 * time.Millisecond

// GetProbeTimeDuration combines second and millisecond time increments into a single time.Duration
func GetProbeTimeDuration(seconds int32, milliseconds *int32) time.Duration {
	if milliseconds != nil {
		return time.Duration(seconds)*time.Second + time.Duration(*milliseconds)*time.Millisecond
	}
	return time.Duration(seconds) * time.Second
}
