package types

// DefaultIndex is the default capability global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default Capability genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		// this line is used by starport scaffolding # ibc/genesistype/default
		// this line is used by starport scaffolding # genesis/types/default
		Params:      DefaultParams(),
		Pools:       nil,
		Delegations: nil,
	}
}

func NewGenesis(params Params, pools []*Pool, delegations []*Delegation) *GenesisState {
	return &GenesisState{
		Params:      params,
		Pools:       pools,
		Delegations: delegations,
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// this line is used by starport scaffolding # ibc/genesistype/validate

	// this line is used by starport scaffolding # genesis/types/validate

	return nil
}
