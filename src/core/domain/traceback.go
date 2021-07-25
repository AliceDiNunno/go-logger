package domain

type TracebackEntry struct {
	File   string
	Line   int
	Method string
}

type Traceback struct {
	Message   string
	Traceback []TracebackEntry
}
