package keeper

import (
	"github.com/airchains-network/airsettle/x/airsettle/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"fmt"
)

func (k Keeper) AddExecutionLayerHelper(ctx sdk.Context, exelayer types.Exelayer, creator string) error {

	// check if admin have a chain already.
	admin_store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ChainAdminKey))
	exeLayerId := []byte(exelayer.Id) // data to save
	byteChainId := admin_store.Get([]byte(creator))
	if byteChainId != nil {
		errormsg := "Admin already have a chain. Chain ID: " + string(byteChainId)
		return  fmt.Errorf(errormsg)
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ExecutionLayerKey))
	b := k.cdc.MustMarshal(&exelayer)
	store.Set([]byte(exelayer.Id), b)

	return nil // no error
}
