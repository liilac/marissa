package vanity

// Manager is an object that manages address-finding worker threads
type Manager struct {
	// workers stores pointers to all worker threads
	workers []*Worker
}
