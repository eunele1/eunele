package dpds

type DotProcessor interface {
    Process(dt *DotTree, sourceDot *MetaDot, sourceDotRoute string)
}

type DotProcessorChannel struct {
	
}

func (dpc *DotProcessorChannel) Process(dt *DotTree, sourceDot *MetaDot, sourceDotRoute string) {
	return
}
