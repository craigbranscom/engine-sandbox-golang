package engine

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Profiler struct {
	frameCount      uint32
	lastFrameTime   float64
	framesPerSecond float64
}

func NewProfiler() *Profiler {
	return &Profiler{
		frameCount:      0,
		lastFrameTime:   0.0,
		framesPerSecond: 0.0,
	}
}

func (profiler *Profiler) UpdateProfiler() {
	currentTime := glfw.GetTime()

	profiler.frameCount++

	timeDelta := currentTime - profiler.lastFrameTime
	if timeDelta >= 1.0 {
		profiler.framesPerSecond = float64(profiler.frameCount) / timeDelta
		profiler.frameCount = 0
		profiler.lastFrameTime = currentTime
	}
}

func (profiler *Profiler) FramesPerSecond() float64 {
	return profiler.framesPerSecond
}
