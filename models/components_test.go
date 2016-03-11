package models

import (
	"log"
	"testing"
)

type componentTest struct {
	componentDir string
	expected     Component
}

var componentTests = []componentTest{
	{"./component_fixtures/EC2", Component{
		Name:          "Amazon Elastic Compute Cloud",
		Key:           "EC2",
		References:    []GeneralReference{GeneralReference{}},
		Verifications: []VerificationReference{VerificationReference{}, VerificationReference{}},
		Satisfies:     []Satisfies{Satisfies{}, Satisfies{}, Satisfies{}, Satisfies{}},
		SchemaVersion: 2.0,
	}},
	{"./component_fixtures/EC2WithKey", Component{
		Name:          "Amazon Elastic Compute Cloud",
		Key:           "EC2",
		References:    []GeneralReference{GeneralReference{}},
		Verifications: []VerificationReference{VerificationReference{}, VerificationReference{}},
		Satisfies:     []Satisfies{Satisfies{}, Satisfies{}, Satisfies{}, Satisfies{}},
		SchemaVersion: 2.0,
	}},
}

func TestLoadComponent(t *testing.T) {
	for _, example := range componentTests {

		openControl := &OpenControl{Components: make(map[string]*Component)}
		openControl.LoadComponent(example.componentDir)
		log.Println(openControl)
		actual := openControl.Components[example.expected.Key]

		if example.expected.Key != actual.Key {
			t.Errorf("Expected %s, Actual: %s", example.expected.Key, actual.Key)
		}
		if example.expected.Name != actual.Name {
			t.Errorf("Expected %s, Actual: %s", example.expected.Name, actual.Name)
		}

		if example.expected.SchemaVersion != actual.SchemaVersion {
			t.Errorf("Expected %s, Actual: %s", example.expected.SchemaVersion, actual.SchemaVersion)
		}

		if len(example.expected.References) != len(actual.References) {
			t.Errorf("Expected %d, Actual: %d", len(example.expected.References), len(actual.References))
		}

		if len(example.expected.Satisfies) != len(actual.Satisfies) {
			t.Errorf("Expected %d, Actual: %d", len(example.expected.Satisfies), len(actual.Satisfies))
		}

		if len(example.expected.Verifications) != len(actual.Verifications) {
			t.Errorf("Expected %d, Actual: %d", len(example.expected.Verifications), len(actual.Verifications))
		}

	}

}
