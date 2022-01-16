package prelaunch_data_generator

import (
	"github.com/kurtosis-tech/eth2-merge-kurtosis-module/kurtosis-module/impl/prelaunch_data_generator/cl"
	"github.com/kurtosis-tech/eth2-merge-kurtosis-module/kurtosis-module/impl/prelaunch_data_generator/el_genesis"
	"github.com/kurtosis-tech/kurtosis-core-api-lib/api/golang/lib/services"
	"github.com/kurtosis-tech/stacktrace"
	"text/template"
	"time"
)

type PrelaunchDataGeneratorContext struct {
	serviceCtx *services.ServiceContext
	networkId string
	depositContractAddress string
	totalTerminalDifficulty uint64
}

func newPrelaunchDataGeneratorContext(serviceCtx *services.ServiceContext, networkId string, depositContractAddress string, totalTerminalDifficulty uint64) *PrelaunchDataGeneratorContext {
	return &PrelaunchDataGeneratorContext{serviceCtx: serviceCtx, networkId: networkId, depositContractAddress: depositContractAddress, totalTerminalDifficulty: totalTerminalDifficulty}
}

func (ctx *PrelaunchDataGeneratorContext) GenerateELGenesisData(
	chainspecAndGethGenesisGenerationConfigTemplate *template.Template,
	nethermindGenesisConfigJsonTemplate *template.Template,
) (*el_genesis.ELGenesisData, error) {
	genesisUnixTimestamp := uint64(time.Now().Unix())
	result, err := el_genesis.GenerateELGenesisData(
		ctx.serviceCtx,
		chainspecAndGethGenesisGenerationConfigTemplate,
		nethermindGenesisConfigJsonTemplate,
		genesisUnixTimestamp,
		ctx.networkId,
		ctx.depositContractAddress,
		ctx.totalTerminalDifficulty,
	)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred generating the EL genesis data")
	}
	return result, nil
}

func (ctx *PrelaunchDataGeneratorContext) GenerateCLValidatorData() {
	
}

func (ctx *PrelaunchDataGeneratorContext) GenerateCLGenesisData(
	genesisGenerationConfigYmlTemplate *template.Template,
	genesisGenerationMnemonicsYmlTemplate *template.Template,
	secondsPerSlot uint32,
	altairForkEpoch uint64,
	mergeForkEpoch uint64,
	preregisteredValidatorKeysMnemonic string,
	numValidatorNodes uint32,
	numValidatorsPerNode uint32,
) (*cl.CLPrelaunchData, error) {
	genesisUnixTimestamp := uint64(time.Now().Unix())
	result, err := cl.GenerateCLPrelaunchData(
		genesisGenerationConfigYmlTemplate,
		genesisGenerationMnemonicsYmlTemplate,
		ctx.serviceCtx,
		genesisUnixTimestamp,
		ctx.networkId,
		ctx.depositContractAddress,
		ctx.totalTerminalDifficulty,
		secondsPerSlot,
		altairForkEpoch,
		mergeForkEpoch,
		preregisteredValidatorKeysMnemonic,
		numValidatorNodes,
		numValidatorsPerNode,
	)
	if err != nil {
		return nil, stacktrace.Propagate(err, "An error occurred generating the CL prelaunch data")
	}
	return result, nil
}
