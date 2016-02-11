package dpds

import (
	"sync"
)

type DotTree struct {
	dotMap             map[uint64]*MetaDot // Map of meta dots
	nameIdMap          map[string]uint64   // Map of name->ids
	waiter            *sync.WaitGroup
	routeMap           map[string]*MetaDot // Map of routes to attached dot.
}

type DotTreeFactory struct {
	dt     *DotTree // Dot Tree interface
}

func (dtf *DotTreeFactory) GetInstance() *DotTree  {
	return nil
}

var DtFactory *DotTreeFactory = new(DotTreeFactory)

func (dt *DotTree) GetDot(route string) *MetaDot {
	return nil
}


func (dt *DotTree) GenerateRoutes() map[string]*MetaDot {
	return nil
}

func (dt *DotTree) ToJSON() string {
	return ""
}
