package scrape

import (
	"fmt"
	"net/url"
)

// Condition is the condition of the product. It is used to perform product's condition filtering
type Condition int

const (
	ConditionAny Condition = iota
	ConditionNew
	ConditionUsed
	ConditionUnknown
)

// String returns the string representation of the condition
func (c *Condition) String() string {
	switch *c {
	case ConditionNew:
		return "new"
	case ConditionUsed:
		return "used"
	case ConditionUnknown:
		return "unknown"
	default:
		return "any"
	}
}

// Set sets the condition from a string
func (c *Condition) Set(value string) error {
	switch value {
	case "new":
		*c = ConditionNew
	case "used":
		*c = ConditionUsed
	case "unknown":
		*c = ConditionUnknown
	case "any":
		*c = ConditionAny
	default:
		return fmt.Errorf("invalid condition: %s", value)
	}
	return nil
}

// SetQueryParam sets the query parameter for the condition
func (c Condition) SetQueryParam(q url.Values) {
	var queryParam string
	switch c {
	case ConditionNew:
		queryParam = "3"
	case ConditionUsed:
		queryParam = "4"
	case ConditionUnknown:
		queryParam = "10"
	default:
		return
	}
	q.Set("LH_ItemCondition", queryParam)
}
