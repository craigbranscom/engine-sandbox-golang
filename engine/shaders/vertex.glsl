#version 410 core
layout (location = 0) in vec3 aPos;
// layout (location = 1) in vec3 aNormal;

uniform vec3 camera;
uniform vec3 rotation;

// uniform mat4 model;
// uniform mat4 view;
// uniform mat4 projection;

// out vec3 FragPos;
// out vec3 Normal;

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
    mat4 rotationMat = rotateX(rotation.x) * rotateY(rotation.y) * rotateZ(rotation.z);
    vec4 rotatedPos = rotationMat * vec4(aPos, 1.0);
    gl_Position = vec4(rotatedPos.x - camera.x, rotatedPos.y - camera.y, rotatedPos.z, 1.0);
}