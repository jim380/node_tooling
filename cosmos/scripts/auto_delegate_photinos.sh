#!/bin/bash
function pause {
 read -n1 -rsp $'Press any key to continue or Ctrl+C to exit...\n'
}

while true
        do
                CHAINID=game_of_stakes_6
                echo "----------------------------------"
                echo "|           User Input           |"  
                echo "----------------------------------"
                read -s -p "Passphrase: " passphrase
                read -p "
Fee: " FEE
                echo ""
                echo "➤ Fee has been set to $FEE photinos."
                ASSET="`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node --output=json | jq ".value.BaseVestingAccount.BaseAccount.coins" | jq ".[1].amount"| bc`"
               	if [  $ASSET == "0" ]
                then
                        echo ""
                        echo "➤ No STAKE available in balance to be delegated yet."
                        echo ""
                        pause   
                else
                STAKE="`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node --output=json | jq ".value.BaseVestingAccount.BaseAccount.coins" | jq ".[1].amount" | bc`"
                while [[ $STAKE -ne 0 ]]
                        do
                                echo "➤ Stake available pre-withdrawl: $STAKE "
                                echo ""
                                echo "✓✓✓"
                                echo ""
                                #SEQ="`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.sequence" | bc`"
                                #SEQUENCE=$(($SEQ + 1))
                                echo "----------------------------------"
                                echo "|            Withdraw            |"
                                echo "----------------------------------"
                                #echo "➤ Prev Seq: $SEQ "
                                #echo "➤ Next Seq: $SEQUENCE "
                                echo $passphrase|gaiacli tx distr withdraw-rewards --chain-id=$CHAINID cosmosvaloper1pjmngrwcsatsuyy8m3qrunaun67sr9x7z5r2qs --from="CypherCore" --commission --async --fees="$FEE""photinos"
                                sleep 30s
                                echo ""
                                echo "➤ Stake available post-withdrawl: $STAKE "
                                echo ""
                                echo "✓✓✓"
                                echo ""
                                echo "----------------------------------"
                                echo "|            Delegate            |"
                                echo "----------------------------------"
                                #STAKE="`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node --output=json | jq ".value.BaseVestingAccount.BaseAccount.coins" | jq ".[1].amount" | bc`"
                                #SEQ1="`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.sequence" | bc`"
                                #SEQUENCE1=$(($SEQ1 + 1))
                                echo ""
                                #echo "➤ Prev Seq: $SEQ1 "
                                #echo "➤ Next Seq: $SEQUENCE1 "
                                echo ""
                                echo $passphrase|gaiacli tx staking delegate --from="CypherCore" cosmosvaloper1pjmngrwcsatsuyy8m3qrunaun67sr9x7z5r2qs --chain-id=$CHAINID "$STAKE""stake" --async --fees="$FEE""photinos" #--sequence=$SEQUENCE1
                                sleep 10s
                                echo ""
                                VOTINGPOWER="`gaiacli status | jq ".validator_info.voting_power" | bc`"
                                echo "➤ Voting Power: $VOTINGPOWER "
                                echo ""
                                #pause
                                sleep 600s
                done
                echo ""
                pause
                #sleep 12000
                fi
        done
