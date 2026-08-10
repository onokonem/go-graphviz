package gvc

import (
	_ "unsafe"

	"github.com/onokonem/go-graphviz/cdt"
	"github.com/onokonem/go-graphviz/cgraph"
	"github.com/onokonem/go-graphviz/internal/wasm"
)

//go:linkname toGraph github.com/onokonem/go-graphviz/cgraph.toGraph
func toGraph(*wasm.Graph) *cgraph.Graph

//go:linkname toGraphWasm github.com/onokonem/go-graphviz/cgraph.toGraphWasm
func toGraphWasm(*cgraph.Graph) *wasm.Graph

//go:linkname toNode github.com/onokonem/go-graphviz/cgraph.toNode
func toNode(*wasm.Node) *cgraph.Node

//go:linkname toEdge github.com/onokonem/go-graphviz/cgraph.toEdge
func toEdge(*wasm.Edge) *cgraph.Edge

//go:linkname toDictLink github.com/onokonem/go-graphviz/cdt.toLink
func toDictLink(*wasm.DictLink) *cdt.Link

//go:linkname toDictLinkWasm github.com/onokonem/go-graphviz/cdt.toLinkWasm
func toDictLinkWasm(*cdt.Link) *wasm.DictLink
