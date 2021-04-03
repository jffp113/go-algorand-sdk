package algod

import (
	"context"
	"fmt"
	"github.com/jffp113/go-algorand-sdk/client/v2/common"
	"github.com/jffp113/go-algorand-sdk/client/v2/common/models"
)

type AccountInformation struct {
	c       *Client
	account string
}

func (s *AccountInformation) Do(ctx context.Context, headers ...*common.Header) (result models.Account, err error) {
	err = s.c.get(ctx, &result, fmt.Sprintf("/v2/accounts/%s", s.account), nil, headers)
	return
}
