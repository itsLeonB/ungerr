package ungerr

type AppError interface {
	error
	Details() any
	HttpStatus() int
	GrpcStatus() uint32
	ToLogAttrs() []LogAttr
}
