package lookup

import (
	"sourcecode.social/reiver/go-mstdn/api/v1"
	"sourcecode.social/reiver/go-mstdn/ent"
)

type LoaderFunc func(account *ent.Account, acct string) v1.Error
