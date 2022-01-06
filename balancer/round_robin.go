// Copyright GoFrame Author(https://goframe.org). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/katyusha.

package balancer

import "google.golang.org/grpc/balancer/roundrobin"

// RoundRobin just uses grpc Round Robin balancer.
// No need making such wheel ourselves.
const RoundRobin = roundrobin.Name
