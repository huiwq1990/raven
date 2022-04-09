//go:build !linux
// +build !linux

/*
 * Copyright 2022 The OpenYurt Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package network_engine

import (
	"github.com/openyurtio/raven/pkg/types"
	"net"
)

type vxlanAgent struct {
}

func (agent *vxlanAgent) Init(localCniIP net.IP, localGatewayPublicIP net.IP) {

}

func (agent *vxlanAgent) ConnectToGateway(gateway *types.Gateway) error {
	return nil
}

func (agent *vxlanAgent) MTU() (int, error) {
	return 0, nil
}

func (agent *vxlanAgent) Cleanup() {

}
