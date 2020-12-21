package get

import (
	"strings"

	"github.com/project-flogo/core/data/coerce"
)

// Settings of the activity
type Settings struct {
	KeyName    string                 `md:"keyName"`
	Attributes []string               `md:"attributes"`
	Query      map[string]interface{} `md:"query"`
	KeysOnly   bool                   `md:"keysOnly"`
}

// Input of the activity
type Input struct {
	Data              interface{} `md:"data,required"`
	PrivateCollection string      `md:"privateCollection"`
	PageSize          int32       `md:"pageSize"`
	Bookmark          string      `md:"bookmark"`
}

// Output of the activity
type Output struct {
	Code     int           `md:"code"`
	Message  string        `md:"message"`
	Bookmark string        `md:"bookmark"`
	Result   []interface{} `md:"result"`
}

// FromMap sets settings from a map
// construct composite key definition of format {"index": ["field1, "field2"]}
func (h *Settings) FromMap(values map[string]interface{}) error {
	var err error
	if h.KeyName, err = coerce.ToString(values["keyName"]); err != nil {
		return err
	}
	if h.KeysOnly, err = coerce.ToBool(values["keysOnly"]); err != nil {
		return err
	}
	if h.Query, err = coerce.ToObject(values["query"]); err != nil {
		return err
	}

	if attrs, err := coerce.ToArray(values["attributes"]); err == nil && len(attrs) > 0 {
		for _, v := range attrs {
			if f, err := coerce.ToString(v); err == nil {
				if !strings.HasPrefix(f, "$.") {
					// make it valid JsonPath expression
					f = "$." + f
				}
				h.Attributes = append(h.Attributes, f)
			}
		}
	}
	if len(h.Attributes) == 0 {
		if len(h.KeyName) > 0 {
			logger.Warnf("ignore key definition for %s: No attribute is specified\n", h.KeyName)
			h.KeyName = ""
		}
	}
	return nil
}

// ToMap converts activity input to a map
func (i *Input) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"data":              i.Data,
		"privateCollection": i.PrivateCollection,
		"pageSize":          i.PageSize,
		"bookmark":          i.Bookmark,
	}
}

// FromMap sets activity input values from a map
func (i *Input) FromMap(values map[string]interface{}) error {

	var err error
	if i.Data, err = coerce.ToAny(values["data"]); err != nil {
		return err
	}
	if i.PrivateCollection, err = coerce.ToString(values["privateCollection"]); err != nil {
		return err
	}
	if i.PageSize, err = coerce.ToInt32(values["pageSize"]); err != nil {
		return err
	}
	if i.Bookmark, err = coerce.ToString(values["bookmark"]); err != nil {
		return err
	}

	return nil
}

// ToMap converts activity output to a map
func (o *Output) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"code":     o.Code,
		"message":  o.Message,
		"bookmark": o.Bookmark,
		"result":   o.Result,
	}
}

// FromMap sets activity output values from a map
func (o *Output) FromMap(values map[string]interface{}) error {

	var err error
	if o.Code, err = coerce.ToInt(values["code"]); err != nil {
		return err
	}
	if o.Message, err = coerce.ToString(values["message"]); err != nil {
		o.Message = ""
	}
	if o.Bookmark, err = coerce.ToString(values["bookmark"]); err != nil {
		return err
	}
	if o.Result, err = coerce.ToArray(values["result"]); err != nil {
		return err
	}

	return nil
}
