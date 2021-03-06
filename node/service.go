// Copyright © 2018 Kowala SEZC <info@kowala.tech>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package node

import (
	"errors"
	"reflect"

	"github.com/kowala-tech/equilibrium/node/event"
	"github.com/kowala-tech/equilibrium/node/p2p"
)

var (
	errServiceUnknown = errors.New("unknown service")
)

// Context is a collection of service independent options inherited from
// the protocol stack, that is passed to all constructors to be optionally used;
// as well as utility methods to operate on the service environment.
type Context struct {
	cfg      *Config
	Services map[reflect.Type]Service // Index of the already constructed services
	EventMux *event.TypeMux           // Event multiplexer used for decoupled notifications
	//AccountManager  *accounts.Manager        // Account manager created by the node.
	//ConsensusEngine consensus.Engine         // Consensus Engine used by the node.
	//Signer types.Signer // Signer used by the node
	// Currency string (???)
}

// Service retrieves a currently running service registered of a specific type.
func (ctx *Context) Service(service interface{}) error {
	element := reflect.ValueOf(service).Elem()
	if running, ok := ctx.Services[element.Type()]; ok {
		element.Set(reflect.ValueOf(running))
		return nil
	}
	return errServiceUnknown
}

// DataDir retrieves the current data directory.
func (ctx *Context) DataDir() string {
	return ctx.cfg.DataDir
}

// Constructor is the function signature of the constructors needed to be
// registered for service instantiation.
type Constructor func(ctx *Context) (Service, error)

// Service is an individual protocol that can be registered into a node.
type Service interface {
	Start(server *p2p.Host) error
	Stop() error
}
