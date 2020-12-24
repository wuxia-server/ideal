package module

type Option func(*Options)

type Options struct {
	Debug    bool
	Parse    bool
	ConfPath string
	WorkDir  string
	LogDir   string
}

func Debug(v bool) Option {
	return func(o *Options) {
		o.Debug = v
	}
}

func Parse(v bool) Option {
	return func(o *Options) {
		o.Parse = v
	}
}

func ConfPath(v string) Option {
	return func(o *Options) {
		o.ConfPath = v
	}
}

func LogDir(v string) Option {
	return func(o *Options) {
		o.LogDir = v
	}
}
