#!/bin/bash
pause(){
 read -n1 -rsp $'Press any key to continue or Ctrl+C to exit...\n'
}
while true
        do
                CHAINID=game_of_stakes_3
                echo "----------------------------------"
                echo "|           User Input           |"  
                echo "----------------------------------"
                read -s -p "Passphrase: " passphrase
                # echo $passphrase|gaiacli tx dist withdraw-rewards --chain-id "genki-2000" --from "main" --is-validator
                read -p "
Fee: " FEE
                echo ""
                echo "➤ Fee has been set to $FEE photinos."
                # STEAK=`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.coins" | jq ".[0].amount" | bc`
                ASSET="`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.coins" | jq ".[0].denom"| bc`"
                if [  $ASSET == "photinos" ]
                then
                        echo ""
                        echo "➤ No STAKE available in balance to be delegated yet. "
                        echo ""
                        pause   
                else
                STEAK="`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.coins" | jq ".[0].amount" | bc`"
                while [[ $STEAK -ne 0 ]]
                        do
                                echo "➤ Stake available pre-withdrawl: $STEAK "
                                echo ""
                                echo "✓✓✓"
                                echo ""
                                SEQ="`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.sequence" | bc`"
                                SEQUENCE=$(($SEQ + 1))
                                echo "----------------------------------"
                                echo "|            Withdraw            |"
                                echo "----------------------------------"
                                echo "➤ Prev Seq: $SEQ "
                                echo "➤ Next Seq: $SEQUENCE "
                                echo $passphrase|gaiacli tx dist withdraw-rewards --chain-id=$CHAINID --from="CypherCore" --is-validator --fee="$FEE""photinos"
                                sleep 30s
                                echo ""
                                echo "➤ Stake available post-withdrawl: $STEAK "
                                echo ""
                                echo "✓✓✓"
                                echo ""
                                echo "----------------------------------"
                                echo "|            Delegate            |"
                                echo "----------------------------------"
                                STAKE="`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.coins" | jq ".[0].amount" | bc`"
                                SEQ1="`gaiacli query account --chain-id=$CHAINID cosmos1pjmngrwcsatsuyy8m3qrunaun67sr9x78qhlvr --trust-node | jq ".value.sequence" | bc`"
                                SEQUENCE1=$(($SEQ1 + 1))
                                echo ""
                                echo "➤ Prev Seq: $SEQ1 "
                                echo "➤ Next Seq: $SEQUENCE1 "
                                echo ""
                                echo $passphrase|gaiacli tx stake delegate --from="CypherCore" --validator="cosmosvaloper1pjmngrwcsatsuyy8m3qrunaun67sr9x7z5r2qs" --chain-id=$CHAINID --amount="$STAKE""STAKE" --fee="$FEE""photinos" --sequence=$SEQUENCE1
                                sleep 10s
                                echo ""
                                VOTINGPOWER="`gaiacli status | jq ".validator_info.voting_power" | bc`"
                                echo "➤ Voting Power: $VOTINGPOWER "
                                echo ""
                                #pause
                                sleep 600s
                done
                echo ""
                echo "➤ You only have $STEAK STAKE "
                echo ""
                pause
                #sleep 12000
                fi
        done