package pathenv

import "os"

func Path() string { return os.Getenv("PATH") }
