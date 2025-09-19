package api


type Volume struct {
	ID string
	SizeBytes int64
	ExtentSz int64
	Replicas []string
}

type Node struct {
	ID string
	Addr string
}