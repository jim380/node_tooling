#!/bin/bash
# install python if you haven't. sudo apt-get install python
function pause {
 read -n1 -rsp $'Press any key to continue or Ctrl+C to exit...\n'
}

                echo "----------------------------------"
                echo "|           User Input           |"  
                echo "----------------------------------"
                read -p "Chain-id: " CHAINID
                echo "➤ Chain-id has been set to $CHAINID"
                echo ""
                read -s -p "Passphrase: " passphrase
                echo ""
                read -p "
Fee: " FEE
                echo "➤ Fee has been set to $FEE iris."
                echo ""
                read -p "Delegate ratio (e.g. 0.5): " delegateRatio
                echo "➤ Delegate ratio has been set to $delegateRatio."
                echo ""
                delegateAddr="iaa1kfhee2nqrg64krqa97q3ufw9d0phzp3j83mhg4"
                validatorAddr="iva1kfhee2nqrg64krqa97q3ufw9d0phzp3jjq3c4j"
                KEYNAME="Cypher Core-5CCA4F526B9F85DA"
                echo "----------------------------------"
                echo "|        Staring in 3 sec         |"  
                echo "----------------------------------"
                sleep 3s
while true
        do
                currentBalance="`iriscli bank account --chain-id=$CHAINID $delegateAddr | jq ".coins[0]" | bc | sed -e 's/\(.*\)iris/\1/'`"
                isEnough=`python -c "print $currentBalance > $FEE"`
               	if [  $isEnough == "False" ]
                then
                        echo ""
                        echo "➤ Not enough stake to pay for fees."
                        echo ""
                else
                        STAKE_PRE="`iriscli bank account --chain-id=$CHAINID $delegateAddr | jq ".coins[0]" | bc | sed -e 's/\(.*\)iris/\1/'`"
                        isEnoughWithdraw=`python -c "print $STAKE_PRE > $FEE"`
                        if [ $isEnoughWithdraw == "True" ]
                        then
                                echo ""
                                echo "➤ Current balance: $STAKE_PRE "
                                echo ""
                                echo "✓✓✓"
                                echo ""
                                echo "----------------------------------"
                                echo "|            Withdraw            |"
                                echo "----------------------------------"
                                echo $passphrase|iriscli distribution withdraw-rewards --chain-id=$CHAINID --from="$KEYNAME" --fee=$FEE"iris"
                                echo ""
                                echo "➤ Withdrawal completed. Holding for 15 sec to update balance."
                                sleep 15s
                                echo ""
                                STAKE_POST="`iriscli bank account --chain-id=$CHAINID $delegateAddr | jq ".coins[0]" | bc | sed -e 's/\(.*\)iris/\1/'`"
                                echo "➤ Balance post-withdrawl: $STAKE_POST "
                                echo ""
                                echo "✓✓✓"
                                echo ""
                                echo "----------------------------------"
                                echo "|            Delegate            |"
                                echo "----------------------------------"
                                STAKE="`iriscli bank account --chain-id=$CHAINID $delegateAddr | jq ".coins[0]" | bc | sed -e 's/\(.*\)iris/\1/'`"
                                delegateStake=`python -c "print $STAKE*$delegateRatio"`
                                echo ""
                                echo "➤ Stake: $STAKE; Fee: $FEE; Delegat amount: $delegateStake"
                                echo ""
                                isEnoughDelegate=`python -c "print $delegateStake > $FEE"`
                                if [ $isEnoughDelegate == "True" ]
                                then
                                        echo $passphrase|iriscli stake delegate --from="$KEYNAME" --address-validator=$validatorAddr --chain-id=$CHAINID --amount="$delegateStake""iris" --fee=$FEE"iris" #--sequence=$SEQUENCE1
                                        echo ""
                                        echo "➤ Delegate completed. Holding for 15 sec to update voting power."
                                        sleep 15s
                                        VOTINGPOWER="`iriscli status | jq ".validator_info.voting_power" | bc`"
                                        echo ""
                                        echo "➤ Voting Power: $VOTINGPOWER"

                                else 
                                        echo "Delegate amount is enough to pay for fees."
                                fi
                        fi
                fi
                echo ""
                echo "Sleeping for 1 hr"
                sleep 600s
        done
