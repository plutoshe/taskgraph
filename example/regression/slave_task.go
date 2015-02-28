package regression

import "github.com/taskgraph/taskgraph/factory"

type slaveTask struct {
	taskCommon
	totalIteration uint64
}

func (tk *slaveTask) SetEpoch(epoch uint64) {
	if epoch == tk.totalIteration {
		return
	}
	tk.setupGradientProcessor()
}

func (tk *slaveTask) setupGradientProcessor() {
	cp := factory.CreateComposer()
	cp.SetProcessor(&gradientProcessor{})

	for _, from := range tk.treeTopo.GetChildren() {
		cp.CreateInboundChannel(from, "gradient")
	}

	for _, from := range tk.treeTopo.GetParents() {
		cp.CreateInboundChannel(from, "parameter")
		cp.CreateOutboundChannel(from, "parameter")
	}

	cp.Compose()
}