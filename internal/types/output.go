package types

type TableRenderer interface {
	Headers() []string
	Rows() [][]string
	EmptyMessage() string
}

type TableRenderable interface {
	AsTableRenderer() TableRenderer
}
