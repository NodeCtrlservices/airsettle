#!/bin/bash
truncate -s 0 ../air.log
airsettled tx airsettle add-execution-layer  "$(cat ./verificationkey.json)" "basic chain info"  --from alice --gas 3000000 -y 
sleep 3;
chainid=`cat ./chainid.test.air`
creator_address="cosmos1lturv7jcdzlacvuq3tw5xwa8j7w5sz7wtcum2y"

# Query Chains
sleep 1; airsettled query airsettle chain-list $creator_address
sleep 1; airsettled query airsettle chain-list-detailed $creator_address
sleep 1; airsettled query airsettle list-execution-layers
sleep 1; airsettled query airsettle verification-key $chainid
sleep 1; airsettled query airsettle show-execution-layer $chainid

# Create Batch
batchnumber=100 # wrong batch number error.
sleep 1; airsettled tx airsettle add-batch $chainid $batchnumber "0xMerkleRootHash" "0xPrevMerkleRootHash" "$(cat ./zkproof.json)" --from bob -y
batchnumber=1
sleep 1; airsettled tx airsettle add-batch $chainid $batchnumber "0xMerkleRootHash" "0xPrevMerkleRootHash" "$(cat ./zkproof.json)" --from bob -y
sleep 2; airsettled tx airsettle add-batch $chainid $batchnumber "0xMerkleRootHash" "0xPrevMerkleRootHash"  "$(cat ./zkproof.json)" --from $creator_address -y
sleep 2; airsettled query airsettle show-batch $chainid $batchnumber
batchnumber=2
sleep 2; airsettled tx airsettle add-batch $chainid $batchnumber "0xMerkleRootHash" "0xPrevMerkleRootHash" "$(cat ./zkproof.json)" --from $creator_address -y
sleep 2; airsettled query airsettle show-batch $chainid $batchnumber
batchnumber=3
sleep 2; airsettled tx airsettle add-batch $chainid $batchnumber "0xMerkleRootHash" "0xPrevMerkleRootHash" "$(cat ./zkproof.json)" --from $creator_address -y
sleep 2; airsettled query airsettle show-batch $chainid $batchnumber

airsettled query airsettle show-batch $chainid $batchnumber
# Verify Batch
sleep 1; airsettled query airsettle verify $chainid $batchnumber "$(cat ./inputs.json)"


sleep 3; airsettled tx airsettle add-validator "air1a4arux7tsy26dzec6j9effz679me9gqer8hycc" $chainid --from $creator_address -y
sleep 1; airsettled tx airsettle add-validator "air1varhzdweq72eqqpma8ymzjk2a0ka8rdr84zx60" $chainid --from $creator_address -y
sleep 3;

pollid=`cat ./pollid.test.air`

airsettled query airsettle list-add-validators-polls
airsettled query airsettle validator-poll-details $pollid
airsettled query airsettle show-execution-layer $chainid
airsettled tx airsettle submit-validator-vote $pollid true --from $creator_address -y
airsettled tx airsettle submit-validator-vote $pollid true --from bob -y

# airsettled tx airsettle submit-validator-vote "40305d4f-df50-4e39-b7cd-64b3325e9227" true --from alice


sleep 1; airsettled query airsettle verify $chainid $batchnumber "$(cat ./inputs.json)"
airsettled tx airsettle verify-msg $chainid $batchnumber "$(cat ./inputs.json)" --from alice -y