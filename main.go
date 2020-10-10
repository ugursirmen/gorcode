package main

import (
	"fmt"
	"log"

	"github.com/asticode/go-astikit"
	"github.com/asticode/go-astilectron"
)

func main() {
	// Set logger
	l := log.New(log.Writer(), log.Prefix(), log.Flags())

	a, err := astilectron.New(l, astilectron.Options{
		AppName:           "Gorcode",
		BaseDirectoryPath: "web",
	})
	if err != nil {
		l.Fatal(fmt.Errorf("main: creating astilectron failed: %w", err))
	}
	defer a.Close()

	a.HandleSignals()

	if err = a.Start(); err != nil {
		l.Fatal(fmt.Errorf("main: starting astilectron failed: %w", err))
	}

	var w *astilectron.Window
	if w, err = a.NewWindow("web/index.html", &astilectron.WindowOptions{
		Center: astikit.BoolPtr(true),
		Height: astikit.IntPtr(850),
		Width:  astikit.IntPtr(1200),
	}); err != nil {
		l.Fatal(fmt.Errorf("main: new window failed: %w", err))
	}

	if err = w.Create(); err != nil {
		l.Fatal(fmt.Errorf("main: creating window failed: %w", err))
	}

	// w.OpenDevTools()

	a.Wait()
}
