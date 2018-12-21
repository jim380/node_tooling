#!/bin/bash
pause(){
 read -n1 -rsp $'Press any key to continue or Ctrl+C to exit...\n'
}
while true
        do
                CHAINID=game_of_stakes_1
                read -s -p "Passphrase: " passphrase
                # echo $passphrase|gaiacli tx dist withdraw-rewards --chain-id "genki-2000" --from "main" --is-validator
                read -p "
Fee: " FEE
                echo "Fee has been set to $FEE STAKE."
                # STEAK=`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.coins" | jq ".[0].amount" | bc`
                ASSET="`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.coins" | jq ".[0].denom"| bc`"
                if [  $ASSET == "photinos" ]
                then
                        echo "No STAKE yet"  
                        pause   
                else
                STEAK=`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.coins" | jq ".[0].amount" | bc`
                # echo "$STEAK"
                NETSTAKE=$(($STEAK - $FEE))
                # echo "Net stake: $NETSTAKE"
                while [[ $STEAK -ne 0 ]] && [[ $NETSTAKE -gt 0 ]]
                        do
                                echo "----------------------------------"
                                echo " Blance: ""$((STEAK ))"" "
                                echo "----------------------------------"
                                echo "----------------------------------"
                                echo " Withdraw "
                                echo "----------------------------------"
                                echo $passphrase|gaiacli tx dist withdraw-rewards --chain-id=$CHAINID --from="CypherCore" --is-validator
                                sleep 10s
                                echo "----------------------------------"
                                echo " Delegate "
                                echo "----------------------------------"
                                echo $passphrase|gaiacli tx stake delegate --from="CypherCore" --validator "cosmosvaloper1pjmngrwcsatsuyy8m3qrunaun67sr9x7z5r2qs" --chain-id "game_of_stakes" --amount "$NETSTAKE""STAKE" --sequence 0
                                sleep 10s
                                echo "----------------------------------"
                                echo " Balance "
                                echo "----------------------------------"
                                gaiacli status | jq ".[].voting_power"
                                sleep 600s
                done
                echo "----------------------------------"
                echo "You only have $STEAK STAKE"
                echo "----------------------------------"
                sleep 12000
                fi
        done