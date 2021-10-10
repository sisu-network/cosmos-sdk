package client

import (
	govclient "github.com/sisu-network/cosmos-sdk/x/gov/client"
	"github.com/sisu-network/cosmos-sdk/x/params/client/cli"
	"github.com/sisu-network/cosmos-sdk/x/params/client/rest"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd, rest.ProposalRESTHandler)
