package admn

import (
	"encoding/json"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-opt"
	"github.com/reiver/go-nul"

	"github.com/reiver/go-mstdn/ent"
)

var _ json.Marshaler = Account{}

// Account represents a Mastodon API "Admin::Account".
//
// Note that this is NOT an admin-account.
// But instead this is an administrated-account.
// I.e., the administrator's view of an account.
//
// See:
// https://docs.joinmastodon.org/entities/Admin_Account/
type Account struct {
	ID                     opt.Optional[string] `json:"id"`
	UserName               opt.Optional[string] `json:"username"`
	Domain                 nul.Nullable[string] `json:"domain"`
	CreatedAt              opt.Optional[string] `json:"created_at"`
	EMail                  opt.Optional[string] `json:"email"`
	IP                     nul.Nullable[string] `json:"ip"`
	IPs                  []IP                   `json:"ips"`
	Locale                 opt.Optional[string] `json:"locale"`
	InviteRequest          nul.Nullable[string] `json:"invite_request"`
	Role                   ent.Role             `json:"role"`
	Confirmed              opt.Optional[bool]   `json:"confirmed"`
	Approved               opt.Optional[bool]   `json:"approved"`
	Disabled               opt.Optional[bool]   `json:"disabled"`
	Silenced               opt.Optional[bool]   `json:"silenced"`
	Suspended              opt.Optional[bool]   `json:"suspended"`
	Account                ent.Account          `json:"account"`
	CreatedByApplicationID opt.Optional[string] `json:"created_by_application_id"`
	InvitedByAccountID     opt.Optional[string] `json:"invited_by_account_id"`
}

func (receiver Account) MarshalJSON() ([]byte, error) {

	var buffer []byte

	buffer = append(buffer, "{"...)

	{
		buffer = append(buffer, `"id":`...)

		marshaled, err := json.Marshal(receiver.ID)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.ID as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"username":`...)

		marshaled, err := json.Marshal(receiver.UserName)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.UserName as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"domain":`...)

		marshaled, err := json.Marshal(receiver.Domain)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.Domain as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"created_at":`...)

		marshaled, err := json.Marshal(receiver.CreatedAt)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.CreatedAt as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"email":`...)

		marshaled, err := json.Marshal(receiver.EMail)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.EMail as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"ip":`...)

		marshaled, err := json.Marshal(receiver.IP)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.IP as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"ips":`...)

		var src interface{} = receiver.IPs
		if nil == receiver.IPs {
			src = []IP{}
		}

		marshaled, err := json.Marshal(src)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.IPs as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"locale":`...)

		marshaled, err := json.Marshal(receiver.Locale)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.Locale as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"invite_request":`...)

		marshaled, err := json.Marshal(receiver.InviteRequest)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.InviteRequest as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"role":`...)

		marshaled, err := json.Marshal(receiver.Role)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.Role as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"confirmed":`...)

		marshaled, err := json.Marshal(receiver.Confirmed)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.Confirmed as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"approved":`...)

		marshaled, err := json.Marshal(receiver.Approved)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.Approved as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"disabled":`...)

		marshaled, err := json.Marshal(receiver.Disabled)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.Disabled as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"silenced":`...)

		marshaled, err := json.Marshal(receiver.Silenced)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.Silenced as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"suspended":`...)

		marshaled, err := json.Marshal(receiver.Suspended)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.Suspended as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		buffer = append(buffer, `,"account":`...)

		marshaled, err := json.Marshal(receiver.Account)
		if nil != err {
			return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.Account as JSON: %w", err)
		}

		buffer = append(buffer, marshaled...)
	}

	{
		if opt.Nothing[string]() != receiver.CreatedByApplicationID {

			buffer = append(buffer, `,"created_by_application_id":`...)

			marshaled, err := json.Marshal(receiver.CreatedByApplicationID)
			if nil != err {
				return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.CreatedByApplicationID as JSON: %w", err)
			}

			buffer = append(buffer, marshaled...)
		}
	}

	{
		if opt.Nothing[string]() != receiver.InvitedByAccountID {

			buffer = append(buffer, `,"invited_by_account_id":`...)

			marshaled, err := json.Marshal(receiver.InvitedByAccountID)
			if nil != err {
				return nil, erorr.Errorf("mstdn/ent/admn: could not marshal admn.Account.InvitedByAccountID as JSON: %w", err)
			}

			buffer = append(buffer, marshaled...)
		}
	}

	buffer = append(buffer, "}"...)

	return buffer, nil
}
