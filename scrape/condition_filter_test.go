package scrape

import (
	"net/url"
	"testing"
)

func TestCondition_Set(t *testing.T) {
	tests := []struct {
		value    string
		expected Condition
	}{
		{"new", ConditionNew},
		{"used", ConditionUsed},
		{"unknown", ConditionUnknown},
		{"any", ConditionAny},
	}

	for _, test := range tests {
		var c Condition
		err := c.Set(test.value)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if c != test.expected {
			t.Errorf("unexpected condition value: got %v, want %v", c, test.expected)
		}
	}

	invalidValue := "invalid"
	var c Condition
	err := c.Set(invalidValue)
	if err == nil {
		t.Errorf("expected error for invalid condition value: %s", invalidValue)
	}
}

func TestSetQueryParam(t *testing.T) {
	// Create a new Condition instance for each test case
	tests := []struct {
		condition Condition
		expected  string
	}{
		{ConditionNew, "3"},
		{ConditionUsed, "4"},
		{ConditionUnknown, "10"},
		{Condition(100), ""},
	}

	// Iterate over the test cases
	for _, test := range tests {
		// Create a new url.Values instance
		q := url.Values{}

		// Call the SetQueryParam method with the test condition
		test.condition.SetQueryParam(q)

		// Check if the query parameter was set correctly
		if q.Get("LH_ItemCondition") != test.expected {
			t.Errorf(
				"SetQueryParam(%v) = %s; want %s",
				test.condition,
				q.Get("LH_ItemCondition"),
				test.expected,
			)
		}
	}
}

func TestCondition_String(t *testing.T) {
	tests := []struct {
		name      string
		condition Condition
		want      string
	}{
		{
			name:      "ConditionNew",
			condition: ConditionNew,
			want:      "new",
		},
		{
			name:      "ConditionUsed",
			condition: ConditionUsed,
			want:      "used",
		},
		{
			name:      "ConditionUnknown",
			condition: ConditionUnknown,
			want:      "unknown",
		},
		{
			name:      "Default",
			condition: Condition(10),
			want:      "any",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.condition.String(); got != tt.want {
				t.Errorf("Condition.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
