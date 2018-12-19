#!/bin/bash
#                                                                                                         
#                                                  jim380 <admin@cyphercore.io>
#  ============================================================================
#  
#  Copyright (C) 2018 jim380
#  
#  Permission is hereby granted, free of charge, to any person obtaining
#  a copy of this software and associated documentation files (the
#  "Software"), to deal in the Software without restriction, including
#  without limitation the rights to use, copy, modify, merge, publish,
#  distribute, sublicense, and/or sell copies of the Software, and to
#  permit persons to whom the Software is furnished to do so, subject to
#  the following conditions:
#  
#  The above copyright notice and this permission notice shall be
#  included in all copies or substantial portions of the Software.
#  
#  THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
#  EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
#  MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
#  IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
#  CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
#  TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
#  SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
#  
#  ============================================================================
echo "-----------------------------------------"
echo "*                Welcome                *"
echo "-----------------------------------------"
retries=0
max_retries=10
CHAIN_ID=genki-3000
# ==============================================
# Main Menu
# ==============================================
while [ $retries -lt $max_retries ]
do 
    read -p "What would you like to do?
1) Initialize light client
2) Query remote node for status
3) Create or query a Gaia CLI configuration file
4) Querying subcommands
5) Transactions subcommands
6) Start LCD (light-client daemon), a local REST server
7) Add or view local private keys
8) Print the app version
9) Help about any command

Option #: " USR_INPUT
echo "--------------------------"
    case $USR_INPUT in
# ==============================================
# Main - 1
# ==============================================
        "1")
            #clear
            echo Initialize light client
            ;;
# ==============================================
# Main - 2
# ==============================================
        "2")
            #clear
            echo "Query remote node for status"
            ;;
# ==============================================
# Main - 3
# ==============================================
        "3")
            #clear
            echo Create or query a Gaia CLI configuration file
            ;;
# ==============================================
# Main - 4
# ==============================================
        "4")
            #clear
            read -p "What would you like to query?
1) tendermint-validator-set
2) block
3) txs
4) tx
5) account
6) gov
7) stake
8) slashing

Option #: " USR_INPUT_2
echo "--------------------------"
            case $USR_INPUT_2 in
        # ==============================================
        # Main - 4 --> Sub - 1 
        # ==============================================
                "1")
                    echo "Command to execute:

gaiacli query tendermint-validator-set --chain-id=$CHAIN_ID --trust-node=true
"
                    echo ""
                    read -p "Looks good? (y or n): " USR_INPUT_3
                    case $USR_INPUT_3 in
                        "y")
                            gaiacli query tendermint-validator-set --chain-id=$CHAIN_ID --trust-node=true
                            ;;
                        "n")
                            echo "back to menu"
                            ;;
                          *)
                            echo Invalid input.
                            ;;
                    esac
                    ;;
        # ==============================================
        # Main - 4 --> Sub - 2 
        # ==============================================
                "2")
                    read -p "Block #: " BLOCKHEIGHT
                    echo "Command to execute:

gaiacli query block $BLOCKHEIGHT --trust-node=true
"
                    read -p "Looks good? (y or n): " USR_INPUT_3
                    case $USR_INPUT_3 in
                        "y")
                            gaiacli query block $BLOCKHEIGHT --trust-node=true
                            ;;
                        "n")
                            echo "back to menu"
                            ;;
                          *)
                            echo Invalid input.
                            ;;
                    esac
                    ;;
        # ==============================================
        # Main - 4 --> Sub - 3 
        # ==============================================
                "3")
                    
                    echo Create or query a Gaia CLI configuration file
                    ;;
        # ==============================================
        # Main - 4 --> Sub - * 
        # ==============================================
                  *)
                    echo Invalid input.
                    ;;
            esac
        # ==============================================
            sleep 60
            ;;
# ==============================================
# Main - 5
# ==============================================
        "5")
            echo Transactions subcommands
            ;;
# ==============================================
# Main - 6
# ==============================================
        "6")
            echo "Start LCD (light-client daemon), a local REST server"
            ;;
# ==============================================
# Main - 7
# ==============================================
        "7")
            echo "Add or view local private keys"
            ;;
# ==============================================
# Main - 8
# ==============================================
        "8")
            echo "Print the app version"
            ;;
# ==============================================
# Main - 9
# ==============================================
        "9")
            echo "Help about any command"
            ;;
# ==============================================
# Main - *
# ==============================================
        *)
            echo Invalid input.
            ;;
    esac
# ==============================================
done