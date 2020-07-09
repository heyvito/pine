package pine

import "os"

type Writer interface {
	Bug(msg string, params ...interface{})
	Disk(msg string, params ...interface{})
	Error(msg string, params ...interface{})
	Finish(msg string, params ...interface{})
	Info(msg string, params ...interface{})
	Lock(msg string, params ...interface{})
	Plug(msg string, params ...interface{})
	Secure(msg string, params ...interface{})
	Sleep(msg string, params ...interface{})
	Spawn(msg string, params ...interface{})
	Success(msg string, params ...interface{})
	Terminate(msg string, params ...interface{})
	Thread(msg string, params ...interface{})
	Timing(msg string, params ...interface{})
	WTF(msg string, params ...interface{})
	Warn(msg string, params ...interface{})
	Fatal(msg string, params ...interface{})
}

func (w *PineWriter) Bug(msg string, params ...interface{}) {
	w.parent.write(Bug, w.name, nil, msg, params...)
}

func (w *PineWriter) BugExtra(extra, msg string, params ...interface{}) {
	w.parent.write(Bug, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) Bug(msg string, params ...interface{}) {
	w.parent.BugExtra(w.extra, msg, params...)
}

func (w *PineWriter) Disk(msg string, params ...interface{}) {
	w.parent.write(Disk, w.name, nil, msg, params...)
}

func (w *PineWriter) DiskExtra(extra, msg string, params ...interface{}) {
	w.parent.write(Disk, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) Disk(msg string, params ...interface{}) {
	w.parent.DiskExtra(w.extra, msg, params...)
}

func (w *PineWriter) Error(msg string, params ...interface{}) {
	w.parent.write(Error, w.name, nil, msg, params...)
}

func (w *PineWriter) ErrorExtra(extra, msg string, params ...interface{}) {
	w.parent.write(Error, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) Error(msg string, params ...interface{}) {
	w.parent.ErrorExtra(w.extra, msg, params...)
}

func (w *PineWriter) Finish(msg string, params ...interface{}) {
	w.parent.write(Finish, w.name, nil, msg, params...)
}

func (w *PineWriter) FinishExtra(extra, msg string, params ...interface{}) {
	w.parent.write(Finish, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) Finish(msg string, params ...interface{}) {
	w.parent.FinishExtra(w.extra, msg, params...)
}

func (w *PineWriter) Info(msg string, params ...interface{}) {
	w.parent.write(Info, w.name, nil, msg, params...)
}

func (w *PineWriter) InfoExtra(extra, msg string, params ...interface{}) {
	w.parent.write(Info, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) Info(msg string, params ...interface{}) {
	w.parent.InfoExtra(w.extra, msg, params...)
}

func (w *PineWriter) Lock(msg string, params ...interface{}) {
	w.parent.write(Lock, w.name, nil, msg, params...)
}

func (w *PineWriter) LockExtra(extra, msg string, params ...interface{}) {
	w.parent.write(Lock, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) Lock(msg string, params ...interface{}) {
	w.parent.LockExtra(w.extra, msg, params...)
}

func (w *PineWriter) Plug(msg string, params ...interface{}) {
	w.parent.write(Plug, w.name, nil, msg, params...)
}

func (w *PineWriter) PlugExtra(extra, msg string, params ...interface{}) {
	w.parent.write(Plug, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) Plug(msg string, params ...interface{}) {
	w.parent.PlugExtra(w.extra, msg, params...)
}

func (w *PineWriter) Secure(msg string, params ...interface{}) {
	w.parent.write(Secure, w.name, nil, msg, params...)
}

func (w *PineWriter) SecureExtra(extra, msg string, params ...interface{}) {
	w.parent.write(Secure, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) Secure(msg string, params ...interface{}) {
	w.parent.SecureExtra(w.extra, msg, params...)
}

func (w *PineWriter) Sleep(msg string, params ...interface{}) {
	w.parent.write(Sleep, w.name, nil, msg, params...)
}

func (w *PineWriter) SleepExtra(extra, msg string, params ...interface{}) {
	w.parent.write(Sleep, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) Sleep(msg string, params ...interface{}) {
	w.parent.SleepExtra(w.extra, msg, params...)
}

func (w *PineWriter) Spawn(msg string, params ...interface{}) {
	w.parent.write(Spawn, w.name, nil, msg, params...)
}

func (w *PineWriter) SpawnExtra(extra, msg string, params ...interface{}) {
	w.parent.write(Spawn, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) Spawn(msg string, params ...interface{}) {
	w.parent.SpawnExtra(w.extra, msg, params...)
}

func (w *PineWriter) Success(msg string, params ...interface{}) {
	w.parent.write(Success, w.name, nil, msg, params...)
}

func (w *PineWriter) SuccessExtra(extra, msg string, params ...interface{}) {
	w.parent.write(Success, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) Success(msg string, params ...interface{}) {
	w.parent.SuccessExtra(w.extra, msg, params...)
}

func (w *PineWriter) Terminate(msg string, params ...interface{}) {
	w.parent.write(Terminate, w.name, nil, msg, params...)
}

func (w *PineWriter) TerminateExtra(extra, msg string, params ...interface{}) {
	w.parent.write(Terminate, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) Terminate(msg string, params ...interface{}) {
	w.parent.TerminateExtra(w.extra, msg, params...)
}

func (w *PineWriter) Thread(msg string, params ...interface{}) {
	w.parent.write(Thread, w.name, nil, msg, params...)
}

func (w *PineWriter) ThreadExtra(extra, msg string, params ...interface{}) {
	w.parent.write(Thread, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) Thread(msg string, params ...interface{}) {
	w.parent.ThreadExtra(w.extra, msg, params...)
}

func (w *PineWriter) Timing(msg string, params ...interface{}) {
	w.parent.write(Timing, w.name, nil, msg, params...)
}

func (w *PineWriter) TimingExtra(extra, msg string, params ...interface{}) {
	w.parent.write(Timing, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) Timing(msg string, params ...interface{}) {
	w.parent.TimingExtra(w.extra, msg, params...)
}

func (w *PineWriter) WTF(msg string, params ...interface{}) {
	w.parent.write(WTF, w.name, nil, msg, params...)
}

func (w *PineWriter) WTFExtra(extra, msg string, params ...interface{}) {
	w.parent.write(WTF, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) WTF(msg string, params ...interface{}) {
	w.parent.WTFExtra(w.extra, msg, params...)
}

func (w *PineWriter) Warn(msg string, params ...interface{}) {
	w.parent.write(Warn, w.name, nil, msg, params...)
}

func (w *PineWriter) WarnExtra(extra, msg string, params ...interface{}) {
	w.parent.write(Warn, w.name, &extra, msg, params...)
}

func (w *PineExtraWriter) Warn(msg string, params ...interface{}) {
	w.parent.WarnExtra(w.extra, msg, params...)
}

func (w *PineWriter) Fatal(msg string, params ...interface{}) {
	w.Error(msg, params...)
	os.Exit(1)
}

func (w *PineWriter) FatalExtra(extra, msg string, params ...interface{}) {
	w.ErrorExtra(extra, msg, params...)
	os.Exit(1)
}

func (w *PineExtraWriter) Fatal(msg string, params ...interface{}) {
	w.parent.FatalExtra(w.extra, msg, params...)
}
