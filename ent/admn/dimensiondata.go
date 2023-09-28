package admn

import (
	"encoding/json"

	"sourcecode.social/reiver/go-erorr"
	"sourcecode.social/reiver/go-opt"
)

var _ json.Marshaler = DimensionData{}

// DimensionData represents a Mastodon API "Admin::Dimension Data".
//
// See:
// https://web.archive.org/web/20230716101711/https://docs.joinmastodon.org/entities/Admin_Dimension/#data
type DimensionData struct {
	Key        opt.Optional[string] `json:"key"`
	HumanKey   opt.Optional[string] `json:"human_key"`
	Value      opt.Optional[string] `json:"value"`
	Unit       opt.Optional[string] `json:"unit"`
	HumanValue opt.Optional[string] `json:"human_value"`
}

func (receiver DimensionData) MarshalJSON() ([]byte, error) {

	var buffer []byte

	buffer = append(buffer, "{"...)

	{
		buffer = append(buffer, `"key":`...)

		marshaled, err := json.Marshal(receiver.Key)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.DimensionData.Key as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"human_key":`...)

		marshaled, err := json.Marshal(receiver.HumanKey)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.DimensionData.HumanKey as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"value":`...)

		marshaled, err := json.Marshal(receiver.Value)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.DimensionData.Value as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		if opt.Nothing[string]() != receiver.Unit {

			buffer = append(buffer, `,"unit":`...)

			marshaled, err := json.Marshal(receiver.Unit)
			if nil != err {
				return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.DimensionData.Unit as JSON: %w", err)
			}

			buffer = append(buffer, marshaled...)
		}
	}

	{
		if opt.Nothing[string]() != receiver.HumanValue {

			buffer = append(buffer, `,"human_value":`...)

			marshaled, err := json.Marshal(receiver.HumanValue)
			if nil != err {
				return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.DimensionData.HumanValue as JSON: %w", err)
			}

			buffer = append(buffer, marshaled...)
		}
	}

	buffer = append(buffer, "}"...)

	return buffer, nil
}
