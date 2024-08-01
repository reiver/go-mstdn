package account_id

import (
	"github.com/reiver/go-mstdn/api/v1"
	"github.com/reiver/go-mstdn/ent"
)

type LoaderFunc func(account *ent.Account, accountid string) v1.Error
