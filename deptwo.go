package deptwo

import depone "github.com/sbenchmark/go-dependency-1"

func Name() string {
	return "io.github.sbenchmark:go-dependency-2"
}

func VersionNumber() string {
	return "1.0.3"
}

func Version() string {
	return Name() + "," + VersionNumber()
}

func Dependency() string {
	return depone.Version()
}

