// This file is automatically generated by qtc from "notify_api.qtpl".
// See https://github.com/valyala/quicktemplate for details.

// // Code generated. DO NOT EDIT.

//line notify_api.qtpl:2
package delphirpc

//line notify_api.qtpl:4
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line notify_api.qtpl:4
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line notify_api.qtpl:4
func (x *NotifyServicesSrc) StreamGoFile(qw422016 *qt422016.Writer) {
	//line notify_api.qtpl:4
	qw422016.N().S(`
package notify

import (
	`)
	//line notify_api.qtpl:8
	for imp := range x.goImports {
		//line notify_api.qtpl:8
		qw422016.N().S(`"`)
		//line notify_api.qtpl:8
		qw422016.E().S(imp)
		//line notify_api.qtpl:8
		qw422016.N().S(`"
    `)
		//line notify_api.qtpl:9
	}
	//line notify_api.qtpl:9
	qw422016.N().S(`
)

type msg int

const (
    msg`)
	//line notify_api.qtpl:15
	qw422016.E().S(x.services[0].serviceName)
	//line notify_api.qtpl:15
	qw422016.N().S(` msg = iota
    `)
	//line notify_api.qtpl:16
	for _, m := range x.services[1:] {
		//line notify_api.qtpl:16
		qw422016.N().S(`msg`)
		//line notify_api.qtpl:16
		qw422016.E().S(m.serviceName)
		//line notify_api.qtpl:16
		qw422016.N().S(`
    `)
		//line notify_api.qtpl:17
	}
	//line notify_api.qtpl:17
	qw422016.N().S(`
)

`)
	//line notify_api.qtpl:20
	for _, m := range x.services {
		//line notify_api.qtpl:20
		qw422016.N().S(`func `)
		//line notify_api.qtpl:20
		qw422016.E().S(m.serviceName)
		//line notify_api.qtpl:20
		qw422016.N().S(`(log *structlog.Logger, arg `)
		//line notify_api.qtpl:20
		qw422016.E().S(m.goType)
		//line notify_api.qtpl:20
		qw422016.N().S(`) {
    if log != nil {
        log.Debug(elco.PeerWindowClassName, "`)
		//line notify_api.qtpl:22
		qw422016.E().S(m.serviceName)
		//line notify_api.qtpl:22
		qw422016.N().S(`", arg, "MSG", msg`)
		//line notify_api.qtpl:22
		qw422016.E().S(m.serviceName)
		//line notify_api.qtpl:22
		qw422016.N().S(`)
    }
    W.`)
		//line notify_api.qtpl:24
		qw422016.E().S(m.notifyFunc)
		//line notify_api.qtpl:24
		qw422016.N().S(`( uintptr(msg`)
		//line notify_api.qtpl:24
		qw422016.E().S(m.serviceName)
		//line notify_api.qtpl:24
		qw422016.N().S(`), `)
		//line notify_api.qtpl:24
		qw422016.N().S(m.instructionArg)
		//line notify_api.qtpl:24
		qw422016.N().S(` )
}
`)
		//line notify_api.qtpl:26
		if m.notifyFunc == "NotifyStr" {
			//line notify_api.qtpl:26
			qw422016.N().S(`
func `)
			//line notify_api.qtpl:27
			qw422016.E().S(m.serviceName)
			//line notify_api.qtpl:27
			qw422016.N().S(`f(log *structlog.Logger, format string, a ...interface{}) {
    if log != nil {
        log.Debug(elco.PeerWindowClassName, "`)
			//line notify_api.qtpl:29
			qw422016.E().S(m.serviceName)
			//line notify_api.qtpl:29
			qw422016.N().S(`", fmt.Sprintf(format,a...), "MSG", msg`)
			//line notify_api.qtpl:29
			qw422016.E().S(m.serviceName)
			//line notify_api.qtpl:29
			qw422016.N().S(`)
    }
    W.Notifyf( uintptr(msg`)
			//line notify_api.qtpl:31
			qw422016.E().S(m.serviceName)
			//line notify_api.qtpl:31
			qw422016.N().S(`), format, a... )
}`)
			//line notify_api.qtpl:32
		}
		//line notify_api.qtpl:32
		qw422016.N().S(`
`)
		//line notify_api.qtpl:33
	}
	//line notify_api.qtpl:33
	qw422016.N().S(`



`)
//line notify_api.qtpl:37
}

//line notify_api.qtpl:37
func (x *NotifyServicesSrc) WriteGoFile(qq422016 qtio422016.Writer) {
	//line notify_api.qtpl:37
	qw422016 := qt422016.AcquireWriter(qq422016)
	//line notify_api.qtpl:37
	x.StreamGoFile(qw422016)
	//line notify_api.qtpl:37
	qt422016.ReleaseWriter(qw422016)
//line notify_api.qtpl:37
}

//line notify_api.qtpl:37
func (x *NotifyServicesSrc) GoFile() string {
	//line notify_api.qtpl:37
	qb422016 := qt422016.AcquireByteBuffer()
	//line notify_api.qtpl:37
	x.WriteGoFile(qb422016)
	//line notify_api.qtpl:37
	qs422016 := string(qb422016.B)
	//line notify_api.qtpl:37
	qt422016.ReleaseByteBuffer(qb422016)
	//line notify_api.qtpl:37
	return qs422016
//line notify_api.qtpl:37
}