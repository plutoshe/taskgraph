package taskgraph

// This interface is used by mapreduce application during taskgraph configuration phase.

type Work struct {
	Config map[string]string
}
