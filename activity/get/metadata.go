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
	History    bool                   `md:"history"`
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
	if h.KeysOnly, err = coerce.ToBool(values["keysOnly"]); err != nil {
		return err
	}
	if h.History, err = coerce.ToBool(values["history"]); err != nil {
		return err
	}
	if h.Query, err = coerce.ToObject(values["query"]); err != nil {
		return err
	}

	keys, err := coerce.ToObject(values["compositeKeys"])
	if err != nil {
		return err
	}
	if len(keys) == 0 {
		return nil
	}

	for k, v := range keys {
		var fields []string
		values, err := coerce.ToArray(v)
		if err != nil || len(values) == 0 {
			logger.Warnf("ignored composite key setting for key %s. error: %+v", k, err)
			continue
		}
		for _, n := range values {
			if f, ok := n.(string); ok && len(f) > 0 {
				path := f
				if !strings.HasPrefix(f, "$.") {
					// make it valid JsonPath expression
					path = "$." + f
				}
				fields = append(fields, path)
			}
		}
		if len(fields) > 0 {
			// pick only the first valid key definition
			// Note: we use map for only a single composite key to avoid Web UI issue on exporting array settings
			h.KeyName = k
			h.Attributes = fields
			logger.Debugf("configured composite key %s with fields %+v", k, fields)
			break
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
