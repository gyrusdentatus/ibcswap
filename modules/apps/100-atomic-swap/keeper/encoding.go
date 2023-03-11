package keeper

import "github.com/ibcswap/ibcswap/v6/modules/apps/100-atomic-swap/types"

// Unmarshal attempts to decode and return an LimitOrder object from
// raw encoded bytes.
func (k Keeper) Unmarshal(bz []byte) (types.AtomicSwapOrder, error) {
	var order types.AtomicSwapOrder
	if err := k.cdc.Unmarshal(bz, &order); err != nil {
		return types.AtomicSwapOrder{}, err
	}

	return order, nil
}

// MustUnmarshalOrder attempts to decode and return an LimitOrder object from
// raw encoded bytes. It panics on error.
func (k Keeper) MustUnmarshalOrder(bz []byte) types.AtomicSwapOrder {
	var order types.AtomicSwapOrder
	k.cdc.MustUnmarshal(bz, &order)
	return order
}

// MarshalOrder attempts to encode an LimitOrder object and returns the
// raw encoded bytes.
func (k Keeper) MarshalOrder(order types.AtomicSwapOrder) ([]byte, error) {
	return k.cdc.Marshal(&order)
}

// MustMarshalOrder attempts to encode an LimitOrder object and returns the
// raw encoded bytes. It panics on error.
func (k Keeper) MustMarshalOrder(order types.AtomicSwapOrder) []byte {
	return k.cdc.MustMarshal(&order)
}
