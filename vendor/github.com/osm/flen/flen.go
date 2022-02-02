package flen

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// flagSet contains the CommandLine flag set.
var flagSet *flag.FlagSet = flag.CommandLine

// envPrefix holds a prefix string that will be prepended to all environment
// variables.
var envPrefix string

// SetEnvPrefix sets the envPrefix variable.
func SetEnvPrefix(p string) {
	envPrefix = p
}

// Parse parses the flags, command line arguments will be preferred over
// environment variables.
func Parse() error {
	// Iterate over all the flags and add the environment variable name to
	// the help text.
	flagSet.VisitAll(func(f *flag.Flag) {
		f.Usage = fmt.Sprintf("%s %s", fmt.Sprintf("[%s]", flagToEnv(f.Name)), f.Usage)
	})

	// Iterate over the environment variables and set the flag value to
	// the value that exists in the environment.
	for _, line := range os.Environ() {
		// Split the environment variable by the first '=' character,
		// convert the environment variable name to the equivalent
		// flag name.
		split := strings.SplitN(line, "=", 2)
		key := envToFlag(split[0])
		value := split[1]

		// Check if there's a flag for the given key name, if there's
		// a flag and the value is set we'll use it.
		f := flagSet.Lookup(key)
		if f != nil && value != "" {
			flagSet.Set(key, value)
		}
	}

	// Parse the flags.
	return flagSet.Parse(os.Args[1:])
}

// envToFlag converts the environment variable name to a kebab cased string,
// if the envPrefix is set we'll strip that part from the string as well.
func envToFlag(e string) string {
	flag := strings.ReplaceAll(strings.ToLower(e), "_", "-")

	if strings.HasPrefix(flag, strings.ToLower(envPrefix)+"-") {
		return flag[len(envPrefix)+1:]
	}

	return flag
}

// flagToEnv converts the kebab cased string to upper cased snake casing, if
// envPrefix is set we'll prepend the prefix to the name as well.
func flagToEnv(f string) string {
	if envPrefix != "" {
		return strings.ToUpper(envPrefix) + "_" + strings.ReplaceAll(strings.ToUpper(f), "-", "_")
	}

	return strings.Replace(strings.ToUpper(f), "-", "_", -1)
}
