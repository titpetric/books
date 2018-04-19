package factory

import (
	"fmt"
	"time"
)

/* Log query statistics to stdout */
type DatabaseProfilerStdout struct {
}

func (*DatabaseProfilerStdout) Pre(query string, args ...interface{}) *DatabaseProfilerContext {
	return &DatabaseProfilerContext{
		Query: query,
		Args:  fmt.Sprintf("%#v", args),
		Time:  time.Now(),
	}
}

func (*DatabaseProfilerStdout) Post(p *DatabaseProfilerContext) {
	duration := time.Since(p.Time).Seconds()
	fmt.Printf("[%.4fs] %s (%s)\n", duration, p.Query, p.Args)
}

func (*DatabaseProfilerStdout) Flush() {
}

/* Log query statistics to memory */
type DatabaseProfilerMemory struct {
	Log []string
}

func (*DatabaseProfilerMemory) Pre(query string, args ...interface{}) *DatabaseProfilerContext {
	return &DatabaseProfilerContext{
		Query: query,
		Args:  fmt.Sprintf("%#v", args),
		Time:  time.Now(),
	}
}

func (this *DatabaseProfilerMemory) Post(p *DatabaseProfilerContext) {
	duration := time.Since(p.Time).Seconds()
	this.Log = append(this.Log, fmt.Sprintf("[%.4fs] %s (%s)", duration, p.Query, p.Args))
}

func (this *DatabaseProfilerMemory) Flush() {
	count := len(this.Log)
	for _, line := range this.Log[:count] {
		fmt.Println(line)
	}
	this.Log = this.Log[count:]
}
