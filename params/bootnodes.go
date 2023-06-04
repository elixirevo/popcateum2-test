// Copyright 2015 The go-popcateum Authors
// This file is part of the go-popcateum library.
//
// The go-popcateum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-popcateum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-popcateum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/popcateum/go-popcateum/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Popcateum network.
var MainnetBootnodes = []string{
	// gcp-usa-1
	"enode://a299b362d54257afa15c495dced7443ed57333fa00f10c09cdb03f1c5f047e766224c1848dc27b858c430cefabb074601e1220b27a2c2bfcdde929a4a8dd158c@104.155.179.26:60606",

	// gcp-eu-1
	"enode://204d316801dc50fe305f1e54c4202cc990ca70dcb48ceaf4e0a5d88caadb0bbd02179c298e626a68dda228dc7c95b2e3d5b13fcc81d2ab0f44ea653b4f20ef5f@35.187.34.75:60606",

	// gcp-au-1
	"enode://be6d1d32592d4a8efdd80e979226423648d06f8c730a274b8b466c5ea7bba633b458591a6cfb16d9ba808c13911cefde1e45f087e24d0f05765f64e680c09f2a@34.129.173.78:60606",

	// gcp-jo-1
	"enode://903d9c6369a23e991027a5ca9ca7dc25c0f9ed2980dca43b0849e13d7cdb3a1d947df91afb9cc4880141e748800722cecf6ad98820cd53a4664254bec173241b@34.64.121.103:60606",
}

// SepoliaBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Sepolia test network.
var SepoliaBootnodes = []string{}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{}

var V5Bootnodes = []string{}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/popcateum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	case SepoliaGenesisHash:
		net = "sepolia"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
