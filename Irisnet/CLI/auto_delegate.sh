#!/bin/bash
# install python if you haven't. sudo apt-get install python
function pause {
 read -n1 -rsp $'Press any key to continue or Ctrl+C to exit...\n'
}

while true
        do
                echo "----------------------------------"
                echo "|           User Input           |"  
                echo "----------------------------------"
                read -p "Chain-id: " CHAINID
                echo "➤ Chain-id has been set to $CHAINID"
                read -s -p "Passphrase: " passphrase
                read -p "
Fee: " FEE
                echo ""
                echo "➤ Fee has been set to $FEE iris."
                delegateAddr="iaa1kfhee2nqrg64krqa97q3ufw9d0phzp3j83mhg4"
                validatorAddr="iva1kfhee2nqrg64krqa97q3ufw9d0phzp3jjq3c4j"
                KEYNAME="Cypher Core-5CCA4F526B9F85DA"
                ASSET="`iriscli bank account --chain-id=$CHAINID $delegateAddr | jq ".coins[0]" | bc | sed -e 's/\(.*\)iris/\1/'`" 
               	if [  $ASSET == "0" ]
                then
                        echo ""
                        echo "➤ No STAKE available in balance to be delegated yet."
                        echo ""
                        pause   
                else
                STAKE="`iriscli bank account --chain-id=$CHAINID $delegateAddr | jq ".coins[0]" | bc | sed -e 's/\(.*\)iris/\1/'`"
                while [ $STAKE != "0" ]
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
                                echo $passphrase|iriscli distribution withdraw-rewards --chain-id=$CHAINID --from="$KEYNAME" --fee=$FEE"iris"
                                sleep 30s
                                echo ""
                                echo "➤ Stake available post-withdrawl: $STAKE "
                                echo ""
                                echo "✓✓✓"
                                echo ""
                                echo "----------------------------------"
                                echo "|            Delegate            |"
                                echo "----------------------------------"
                                #STAKE="`iriscli bank account --chain-id=$CHAINID faa1kfhee2nqrg64krqa97q3ufw9d0phzp3jl7a0gg | jq ".coins[0]" | bc | sed -e 's/\(.*\)iris/\1/'`"
                                delegateStake=`python -c "print $STAKE - $FEE"`
                                echo "Stake: $STAKE; Fee: $FEE; Delegated Stake: $delegateStake"
                                #SEQ1="`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.sequence" | bc`"
                                #SEQUENCE1=$(($SEQ1 + 1))
                                echo ""
                                #echo "➤ Prev Seq: $SEQ1 "
                                #echo "➤ Next Seq: $SEQUENCE1 "
                                echo ""
                                echo $passphrase|iriscli stake delegate --from="Cypher Core-5CCA4F526B9F85DA" --address-validator=$validatorAddr --chain-id=$CHAINID --amount="$delegateStake""iris" --fee=$FEE"iris" #--sequence=$SEQUENCE1
                                sleep 10s
                                echo ""
                                VOTINGPOWER="`iriscli status | jq ".validator_info.voting_power" | bc`"
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
