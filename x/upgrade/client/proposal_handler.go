package client

import (
	govclient "github.com/sisu-network/cosmos-sdk/x/gov/client"
	"github.com/sisu-network/cosmos-sdk/x/upgrade/client/cli"
	"github.com/sisu-network/cosmos-sdk/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
var CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal, rest.ProposalCancelRESTHandler)
