package limitedcache

// OpType is cache operation type enum.
type OpType uint8

const (
	// GetOp cache item operation
	GetOp = OpType(iota)
	// SetOp cache item operation
	SetOp
	// DeleteOp cache item operation
	DeleteOp
)

// CacheOp - operations on cache with key and cache file name
type CacheOp struct {
	E    OpType
	Key  string
	File string
	Err  error
}

// StringOp returns string name of cache operation
func (em *CacheOp) Operation() string {
	s := ""
	switch em.E {
	case GetOp:
		s = "get"
	case SetOp:
		s = "set"
	case DeleteOp:
		s = "delete"
	default:
		s = "unknown"
	}
	return s
}
