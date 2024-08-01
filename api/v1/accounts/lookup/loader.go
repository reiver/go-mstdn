package lookup

import (
	"github.com/reiver/go-mstdn/api/v1"
	"github.com/reiver/go-mstdn/ent"
)

type LoaderFunc func(account *ent.Account, acct string) v1.Error
