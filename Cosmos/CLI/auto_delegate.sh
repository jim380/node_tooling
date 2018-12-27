#!/bin/bash
pause(){
 read -n1 -rsp $'Press any key to continue or Ctrl+C to exit...\n'
}
while true
        do
                CHAINID=genki-4000
                echo "----------------------------------"
                echo " User Input "  
                echo "----------------------------------"
                read -s -p "Passphrase: " passphrase
                # echo $passphrase|gaiacli tx dist withdraw-rewards --chain-id "genki-2000" --from "main" --is-validator
                read -p "
Fee: " FEE
                echo ""
                echo "➤ Fee has been set to $FEE STAKE."
                # STEAK=`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.coins" | jq ".[0].amount" | bc`
                ASSET="`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.coins" | jq ".[0].denom"| bc`"
                if [  $ASSET == "photinos" ]
                then
                        echo "----------------------------------"
                        echo " No STAKE yet "  
                        echo "----------------------------------"
                        echo ""
                        pause   
                else
                STEAK=`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.coins" | jq ".[0].amount" | bc`
                NETSTAKE=$(($STEAK - $FEE))
                # echo "Net stake: $NETSTAKE"
                while [[ $STEAK -ne 0 ]] && [[ $NETSTAKE -gt 0 ]]
                        do
                                echo "----------------------------------"
                                echo " Blance Pre-withdrawl: ""$((STEAK))"" "
                                echo "----------------------------------"
                                echo ""
                                echo "✓✓✓"
                                echo ""
                                echo "----------------------------------"
                                echo " Withdraw "
                                echo "----------------------------------"
                                echo $passphrase|gaiacli tx dist withdraw-rewards --chain-id=$CHAINID --from="CypherCore" --is-validator --fee="$FEE""STAKE"
                                sleep 5s
                                STEAK=`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.coins" | jq ".[0].amount" | bc`
                                NETSTAKE=$(($STEAK - $FEE))
                                SEQ=`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.sequence" | bc`
                                SEQUENCE=$(($SEQ + 1))
                                echo "----------------------------------"
                                echo " Blance Post-withdrawl: ""$((STEAK))"" "
                                echo "----------------------------------"
                                echo ""
                                echo "✓✓✓"
                                echo ""
                                echo "----------------------------------"
                                echo " Delegate "
                                echo "----------------------------------"
                                echo $passphrase|gaiacli tx stake delegate --from="CypherCore" --validator="cosmosvaloper1pjmngrwcsatsuyy8m3qrunaun67sr9x7z5r2qs" --chain-id=$CHAINID --amount="$NETSTAKE""STAKE" --fee="$FEE""STAKE" --sequence=$SEQUENCE
                                sleep 10s
                                VOTINGPOWER=`gaiacli status | jq ".validator_info.voting_power" | bc`
                                echo "----------------------------------"
                                echo " Voting Power: $VOTINGPOWER "
                                echo "----------------------------------"
                                sleep 600s
                done
                echo "----------------------------------"
                echo " You only have $STEAK STAKE "
                echo "----------------------------------"
                sleep 12000
                fi
        done