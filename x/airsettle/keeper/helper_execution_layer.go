package keeper

import (
	"fmt"
	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"strconv"
)

func (k Keeper) AddExecutionLayerHelper(ctx sdk.Context, exelayer types.Exelayer, creator string) error {

	// check if admin have a chain already.
	admin_store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainAdminKey))
	exeLayerId := []byte(exelayer.Id) // data to save
	byteChainId := admin_store.Get([]byte(creator))
	if byteChainId != nil {
		errormsg := "Admin already have a chain. Chain ID: " + string(byteChainId)
		return fmt.Errorf(errormsg)
	}

	// if not have chain, store execution layer id under ChainAdminKey
	admin_store.Set([]byte(creator), []byte(exeLayerId))

	// store verification key ( saved seperately to save memory & api call size )
	vk_store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.VerificationKey))
	vkey_formate := types.Vkey{
		Id:              exelayer.Id,
		VerificationKey: exelayer.VerificationKey,
	}
	vk_binary := k.cdc.MustMarshal(&vkey_formate)
	vk_store.Set([]byte(exelayer.Id), vk_binary)

	// Store Execution Layer data except vKey (verification key)
	exelayer.VerificationKey = "/verificationKey/" + exelayer.Id + "/" // store vKey path instade of whole verification key
	// exelayer.
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionLayerKey))
	b := k.cdc.MustMarshal(&exelayer)
	store.Set([]byte(exelayer.Id), b)

	// count execution layer ++ 
	count_store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.CounterStoreKey))
	count_byte := count_store.Get([]byte("exelayers"))
	if count_byte == nil {
		count_store.Set([]byte("exelayers"), []byte("1"))
	} else {
	
		count := string(count_byte)
		// string to uint64
		countUint64, err := strconv.ParseUint(count, 10, 64)
		if err != nil {
			return err
		}
		countUint64++
		// uint64 to string
		count = strconv.FormatUint(countUint64, 10)
		count_store.Set([]byte("exelayers"), []byte(count))
	}

	return nil // no error
}


func (k Keeper) GetExelayerById(ctx sdk.Context, id string) (val types.Exelayer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionLayerKey))

	b := store.Get([]byte(id))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}



func (k Keeper) GetExecutionlayers(ctx sdk.Context, id string) (val types.Exelayer, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionLayerKey))

	b := store.Get(types.ExelayerKey(
		id,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// get execution layer id by address
func (k Keeper) GetExelayerIdByAddress(ctx sdk.Context, address string) (val string, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainAdminKey))

	b := store.Get([]byte(address))
	if b == nil {
		return val, false
	}

	return string(b), true
}
