package requst

import (
	"github.com/panjf2000/ants/v2"
	"log"
	"sync"
)

var TaskWaitGroup sync.WaitGroup

type RequetsPool struct {
	Pool *ants.PoolWithFunc
}

func (pool *RequetsPool) Wait() {
	TaskWaitGroup.Wait()
}
func (pool *RequetsPool) Close() {
	pool.Pool.Release()
}

func (pool *RequetsPool) Done() {
	TaskWaitGroup.Done()
}
func InitRequestPool(threadNum int, str string) *RequetsPool {
	var pool *ants.PoolWithFunc
	var err error
	if str == "Sniper" || str == "Cluster bomb" {
		pool, err = ants.NewPoolWithFunc(threadNum, Newtemplate)
		if err != nil {
			log.Fatalf("Init Request Pool Error: %s\n", err.Error())
		}
	}
	if str == "Pitchfork" {
		pool, err = ants.NewPoolWithFunc(threadNum, NewZiptemplate)
		if err != nil {
			log.Fatalf("Init Request Pool Error: %s\n", err.Error())
		}
	}

	return &RequetsPool{
		Pool: pool,
	}
}
