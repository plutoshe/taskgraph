package taskgraph

import "./filesystem"

// This interface is used by mapreduce application during taskgraph configuration phase.
type MapreduceBootstrap interface {
	Bootstrap
	InitWithMapreduceConfig(MapreduceConfig)
}

type MapreduceConfig struct {
	//defined the task num
	MapperNum  uint64
	ShuffleNum uint64
	ReducerNum uint64

	//filesystem
	FilesystemClient filesystem.Client
	//final result output path
	OutputDir string
	//temporary result output path
	InterDir string

	//emit function
	MapperFunc  func(MapreduceTask, string)
	ReducerFunc func(MapreduceTask, string, []string)

	//store the work, appname, and etcdurls
	UserDefined bool
	WorkDir     map[string][]Work
	AppName     string
	EtcdURLs    []string

	//optional, define the buffer size
	ReaderBufferSize int
	WriterBufferSize int
}

type Work struct {
	Config map[string]string
}
