package pine

var typeEmoji = map[msgType]string{
	Success:   "✅",
	Warn:      "⚠️",
	Finish:    "🏁",
	Spawn:     "✨",
	Timing:    "⏱",
	Info:      "💬",
	Error:     "🚨",
	Terminate: "⛔️",
	Disk:      "💾",
	WTF:       "👻",
	Lock:      "🔒",
	Secure:    "🔑",
}

var typeMap = map[msgType]string{
	Error:     "Error",
	Terminate: "Terminate",
	Disk:      "Disk",
	WTF:       "WTF",
	Lock:      "Lock",
	Secure:    "Secure",
	Success:   "Success",
	Warn:      "Warn",
	Finish:    "Finish",
	Spawn:     "Spawn",
	Timing:    "Timing",
	Info:      "Info",
}
