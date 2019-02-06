package fsutil

import "testing"

const (
	groovyProjectLocation    = "test_resources/gradle/groovy_project/"
	groovySubProjectLocation = groovyProjectLocation + "com.example.app/"

	kotlinProjectLocation    = "test_resources/gradle/kotlin_project/"
	kotlinSubProjectLocation = kotlinProjectLocation + "com.example.app/"

	defaultBuildFile       = "build.gradle"
	defaultBuildFileKotlin = "build.gradle.kts"
	defaultGradlewFile     = "gradlew"

	groovyProjectBuildFileLocation    = groovyProjectLocation + defaultBuildFile
	groovySubProjectBuildFileLocation = groovySubProjectLocation + defaultBuildFile

	kotlinProjectBuildFileKotlinLocation    = kotlinProjectLocation + defaultBuildFileKotlin
	kotlinSubProjectBuildFileKotlinLocation = kotlinSubProjectLocation + defaultBuildFileKotlin

	gradleLocation  = "test_resources/gradle/binary/"
	gradlewLocation = groovyProjectLocation + defaultGradlewFile

	javaSrcDir = groovySubProjectLocation + "src/main/java/"
)

type scenarios struct {
	file     string
	location string
	expected string
}

func Test_FindFile(t *testing.T) {
	tests := []scenarios{
		// find project default build file
		{
			file:     defaultBuildFile,
			location: groovyProjectLocation,
			expected: groovyProjectBuildFileLocation,
		},
		// find project default kotlin build file
		{
			file:     defaultBuildFileKotlin,
			location: kotlinProjectLocation,
			expected: kotlinProjectBuildFileKotlinLocation,
		},
		// find groovy sub project default build file
		{
			file:     defaultBuildFile,
			location: groovySubProjectLocation,
			expected: groovySubProjectBuildFileLocation,
		},
		// find kotlin sub project default build file
		{
			file:     defaultBuildFileKotlin,
			location: kotlinSubProjectLocation,
			expected: kotlinSubProjectBuildFileKotlinLocation,
		},
		// find project gradlew
		{
			file:     defaultGradlewFile,
			location: groovyProjectLocation,
			expected: gradlewLocation,
		},
		// find project gradlew from sub directory
		{
			file:     defaultGradlewFile,
			location: javaSrcDir,
			expected: gradlewLocation,
		},
	}

	for _, test := range tests {
		result := FindFile(test.file, test.location)
		if result != test.expected {
			t.Errorf("actual: %s, expected: %s", result, test.expected)
		}
	}
}

func Test_FindRootVolume(t *testing.T) {
	result := FindRootVolume("/tmp/something")
	if result != "/" {
		t.Error("Didnt find the correct root volume")
	}
}
