package main

import (
	"reflect"
	"testing"
)

func TestRunCommand(t *testing.T) {
	input := "toggle 0,0 through 999,999"
	command := stringToCommand(input)

	var lights LightsType
	var lightBrightness LightsBrightnessType

	runCommand(command, &lights, &lightBrightness)

	var expectedLights LightsType
	var expectedLightBrightness LightsBrightnessType

	for i := 0; i < len(expectedLights); i++ {
		for j := 0; j < len(expectedLights[i]); j++ {
			expectedLights[i][j] = true
		}
	}

	for i := 0; i < len(expectedLightBrightness); i++ {
		for j := 0; j < len(expectedLightBrightness[i]); j++ {
			expectedLightBrightness[i][j] = 2
		}
	}

	if !reflect.DeepEqual(lights, expectedLights) {
		t.Errorf("expected %v, got %v", expectedLights, lights)
	}

	if !reflect.DeepEqual(lightBrightness, expectedLightBrightness) {
		t.Errorf("expected %v, got %v", expectedLightBrightness, lightBrightness)
	}
}
