package mapreduce

import (
	"../../taskgraph"
)

type MapreduceTaskBuilder struct{}

func (t *MapreduceTaskBuilder) GetTask(taskID uint64) taskgraph.MapreduceTask {
	return &mapreduceTask{}
}
