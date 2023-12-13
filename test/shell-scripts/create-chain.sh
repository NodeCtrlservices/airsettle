#!/bin/bash

:<<AIR_UPDATES
List of updates
- 1.0.0
    - 1.0.1
        - address prefix: air (from "cosmos")
        - added new message: delete_execution_layer
        - one account can create only one chain
        - added new query: show_execution_layer_by_address
        - added new query: show_execution_layer_by_id
        - added new message: verify
        - added new query: verify_batch
AIR_UPDATES
# actual shell script code starts here

#* create chain
ignite scaffold chain github.com/airchains-network/airsettle --address-prefix air --skip-git --path "airsettle"
cd airsettle


#* execution layer handlers
ignite scaffold type exelayer validator:array.string voting_power:array.uint latest_batch:uint latest_merkle_root_hash verification_key chain_info id creator -y
ignite scaffold message add_execution_layer verification_key chain_info --response success:bool,message -y
#! building...
ignite scaffold query show_execution_layer_by_address address --response exelayer:Exelayer -y
ignite scaffold query show_execution_layer_by_id id --response exelayer:Exelayer -y
ignite scaffold query list_all_execution_layers --response exelayer:Exelayer --paginated -y
#! to run & build...
# delete execution layer: if (latest_batch == 0)
ignite scaffold message delete_execution_layer --response -y
# verification key is too long, so its shown differently...
ignite scaffold type vkey id verification_key -y
ignite scaffold query verification_key id --response vkey -y

#* batch handlers 
ignite scaffold type batch batch_number:uint merkle_root_hash prev_merkle_root_hash zk_proof verified -y
ignite scaffold message add_batch id batch_number:uint merkle_root_hash prev_merkle_root_hash zk_proof -y
ignite scaffold query show_batch id batch_number:uint --response batch:Batch -y
ignite scaffold message verify id batch_number:uint witness -y
ignite scaffold query verify_batch id batch_number:uint witness --response result:bool,message -y

#* handle/add new validator for execution layer.
ignite scaffold type poll poll_id chain_id new_validator votes_done_by:array.string votes:array.string total_validators:uint is_complete:bool start_date poll_creator -y
ignite scaffold message add_validator new_validator_address chain_id --response voting_poll_id -y
ignite scaffold query list_validator_polls --response poll_ids:array.string --paginated -y
ignite scaffold query list_polls chainid --response poll:Poll --paginated
ignite scaffold query validator_poll_details poll_id --response poll:Poll -y
ignite scaffold message submit_validator_vote poll_id vote:bool --response success:bool,poll_result,message,description

#* handle wrong batch