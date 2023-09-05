package engine

import (
	"github.com/go-gl/glfw/v3.3/glfw"
)

type Profiler struct {
	frameCount            uint32
	lastFrameTime         float64
	framesPerSecond       float64
	targetFrameTime       float64
	targetFramesPerSecond float64
}

func NewProfiler(targetFPS float64) *Profiler {
	targetFrameTime := 1.0 / targetFPS
	return &Profiler{
		frameCount:            0,
		lastFrameTime:         0.0,
		framesPerSecond:       0.0,
		targetFrameTime:       targetFrameTime,
		targetFramesPerSecond: targetFPS,
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

func (profiler *Profiler) TargetFrameTime() float64 {
	return profiler.targetFrameTime
}
