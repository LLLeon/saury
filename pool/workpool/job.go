package workpool

// Job represents work to be done.
type Job interface {
	Do() error
}