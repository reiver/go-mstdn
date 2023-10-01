package account_id

import (
	"sourcecode.social/reiver/go-mstdn/api/v1"
	"sourcecode.social/reiver/go-mstdn/ent"
)

type LoaderFunc func(account *ent.Account, accountid string) v1.Error
