package rest

// Structs

type RestOptions struct {
	Actions []string
}

// Static functions

func NewRestOptions() *RestOptions {
	return &RestOptions{}
}
