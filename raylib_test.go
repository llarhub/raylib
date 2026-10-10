//go:build llgo

package raylib_test

import (
	"testing"

	raylib "github.com/llarhub/raylib"
)

func TestImageGenerationAndLayout(t *testing.T) {
	image := raylib.GenImageColor(2, 3, raylib.Color{R: 10, G: 20, B: 30, A: 255})
	if !image.IsImageValid() {
		t.Fatal("GenImageColor returned an invalid image")
	}
	defer image.Unload()
	if image.Width != 2 || image.Height != 3 {
		t.Fatalf("image dimensions = %dx%d, want 2x3", image.Width, image.Height)
	}
	if image.Format != 7 {
		t.Fatalf("image format = %d, want RGBA8 format 7", image.Format)
	}
}

func TestConstantsAndValueTypes(t *testing.T) {
	if raylib.FLAG_WINDOW_RESIZABLE == 0 || raylib.KEY_ESCAPE == raylib.KEY_NULL {
		t.Fatal("raylib constants were not generated")
	}
	vector := raylib.Vector2{X: 1.5, Y: -2.5}
	if vector.X != 1.5 || vector.Y != -2.5 {
		t.Fatalf("Vector2 fields = (%v, %v)", vector.X, vector.Y)
	}
}
