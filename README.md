# Golang Engine Sandbox

# Engine Subsystems

* Importer - Imports an existing scene YAML file and builds the node tree.

* Pipeline - Manages the render pipeline.

* Node Tree - Handles tree-wide functions.

* Nodes - Prebuilt nodes composed of Components that are placed in the node tree to describe the scene.

* Components - Composable data packages that are used to construct Nodes.

## Setup

Install Go and clone the repo. Download all project dependencies. Run with `go run main.go`.

## Usage

* Pan the camera by using the `WASD` keys.

* Rotate the object with the `UI, JK, NM` keys. They will affect the `X, Y, Z` vertex positions respectively.

## Example Shaders
fragment.glsl
```
#version 410 core
out vec4 FragColor;
void main()
{
	FragColor = vec4(1.0, 0.5, 0.2, 1.0);
}
```

vertex.glsl
```
#version 410 core
layout (location = 0) in vec3 aPos;
uniform vec2 camera;
uniform float rotationX;
uniform float rotationY;
uniform float rotationZ;
mat4 rotateX(float angle) {
    float c = cos(angle);
    float s = sin(angle);
    return mat4(
        1.0, 0.0, 0.0, 0.0,
        0.0, c, -s, 0.0,
        0.0, s, c, 0.0,
        0.0, 0.0, 0.0, 1.0
    );
}
mat4 rotateY(float angle) {
    float c = cos(angle);
    float s = sin(angle);
    return mat4(
        c, 0.0, s, 0.0,
        0.0, 1.0, 0.0, 0.0,
        -s, 0.0, c, 0.0,
        0.0, 0.0, 0.0, 1.0
    );
}
mat4 rotateZ(float angle) {
    float c = cos(angle);
    float s = sin(angle);
    return mat4(
        c, -s, 0.0, 0.0,
        s, c, 0.0, 0.0,
        0.0, 0.0, 1.0, 0.0,
        0.0, 0.0, 0.0, 1.0
    );
}
void main()
{
    mat4 rotationMat = rotateX(rotationX) * rotateY(rotationY) * rotateZ(rotationZ);
    vec4 rotatedPos = rotationMat * vec4(aPos, 1.0);
    gl_Position = vec4(rotatedPos.x - camera.x, rotatedPos.y - camera.y, rotatedPos.z, 1.0);
}
```