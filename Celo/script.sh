#!/bin/bash
#                                                                                                         
#                                                  jim380 <admin@cyphercore.io>
#  ============================================================================
#  
#  Copyright (C) 2020 jim380
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
# export CELO_IMAGE=us.gcr.io/celo-testnet/celo-node:baklava
# export NETWORK_ID=121119
# export VALIDATOR_NAME=CypherCore
# export CELO_VALIDATOR_SIGNER_ADDRESS=
# export PROXY_INTERNAL_IP=
# export PROXY_EXTERNAL_IP=
# export PROXY_ENODE=
# export CELO_ATTESTATION_SIGNER_ADDRESS=

URL="https://docs.celo.org/getting-started/baklava-testnet/running-a-validator#running-the-attestation-service"

function main {
    echo -e "Which machine are you on:\n\n1) Local\n2) Validator\n3) Proxy\n4) Attestation\n\nEnter down below (e.g. \"1\" or \"Local\"):"
    read input
    case $input in
      "1" | "Local" | "local")
        input="local"
        echo -e "\n$input machine selected\n"
        execute
      ;;
      "2" | "Validator" | "validator")
        input="validator"
        echo -e "\nValidator machine selected\n"
        execute
      ;;
      "3" | "Proxy" | "proxy")
        input="proxy"
        echo -e "\n$input machine selected\n"
        execute
      ;;
      "4" | "Attestation" | "attestation")
        input="attestation" 
        echo -e "\n$input machine selected\n"
        execute
      ;;
    
      *) 
        echo -e "Invalid input\n" 
      ;;
    esac
}

function execute {
    case $input in
      "local")
        preCheck
        echo -e "\nCleaning up $input machine\n"
        # Stop running node instance and remove it
        docker stop celo-accounts && docker rm celo-accounts
        # Keys check
        keyChecking
        # Delete chain data
        chainDataDel
        # Restart the node
        nodeRestart
      ;;
      "validator")
        preCheck
        echo -e "\nCleaning up $input machine\n"
        # Stop running node instance and remove it
        docker stop celo-validator && docker rm celo-validator
        # Keys check
        keyChecking
        # Delete chain data
        chainDataDel
        # Restart the node
        nodeRestart
      ;;
      "proxy")
        preCheck
        echo -e "\nCleaning up $input machine\n"
        # Stop running node instance and remove it
        docker stop celo-proxy && docker rm celo-proxy
        # Keys check
        keyChecking
        # Delete chain data
        chainDataDel
        # Restart the node
        nodeRestart
      ;;
      "attestation")
        preCheck
        echo -e "\nCleaning up $input machine\n"
        # Stop running node instance and remove it
        docker stop celo-attestations && docker rm celo-attestations
        docker stop celo-attestation-service && docker rm celo-attestation-service
        # Keys check
        keyChecking
        # Delete chain data
        chainDataDel
        # Restart the node
        nodeRestart
      ;;
    
      *) 
        echo -e "Invalid input\n" 
      ;;
    esac
}

function preCheck {
    envVarCheck $CELO_IMAGE
    envVarCheck $NETWORK_ID
    case $input in
      "local")
        echo -e "\nChecking env variables on $input machine\n"
        echo "Skipped."
      ;;
      "validator")
        echo -e "\nChecking env variables on $input machine:\n"
        envVarCheck $CELO_VALIDATOR_ADDRESS
        envVarCheck $CELO_VALIDATOR_GROUP_ADDRESS
        envVarCheck $CELO_VALIDATOR_SIGNER_ADDRESS
        envVarCheck $CELO_VALIDATOR_SIGNER_PUBLIC_KEY
        envVarCheck $CELO_VALIDATOR_SIGNER_SIGNATURE
        envVarCheck $CELO_VALIDATOR_SIGNER_BLS_PUBLIC_KEY
        envVarCheck $CELO_VALIDATOR_SIGNER_BLS_SIGNATURE
        envVarCheck $PROXY_ENODE
        envVarCheck $PROXY_EXTERNAL_IP
        envVarCheck $PROXY_INTERNAL_IP
        envVarCheck $VALIDATOR_NAME
      ;;
      "proxy")
        echo -e "\nChecking env variables on $input machine\n"
        envVarCheck $CELO_VALIDATOR_SIGNER_ADDRESS
        envVarCheck $BOOTNODE_ENODES
        envVarCheck $VALIDATOR_NAME
      ;;
      "attestation")
        echo -e "\nChecking env variables on $input machine\n"
        envVarCheck $CELO_VALIDATOR_SIGNER_ADDRESS
        envVarCheck $CELO_ATTESTATION_SIGNER_ADDRESS
        envVarCheck $BOOTNODE_ENODES
        envVarCheck $VALIDATOR_NAME
        envVarCheck $CELO_ATTESTATION_SIGNER_SIGNATURE
        envVarCheck $CELO_ATTESTATION_SIGNER_ADDRESS
      ;;
    
      *) 
        echo -e "Invalid input\n" 
      ;;
    esac
}

function envVarCheck {
    [ -z "$1" ] && echo '$GOPATH is empty' || echo '$GOPATH is set to:'
    echo $1
    # envVarCheck $BLAH
    # envVarCheck $BLAH1
    # envVarCheck $BLAH2
}

function keyChecking {
    case $input in
      "local")
        echo -e "\nChecking keys on $input machine\n"
        ls ~/Documents/celo-accounts-node/keystore
        echo ""
      ;;
      "validator")
        echo -e "\nChecking keys on $input machine\n"
        ls ~/Documents/celo-validator-node/keystore
        echo ""
      ;;
      "proxy")
        echo -e "\nNo keys are stored on $input machine\n"
      ;;
      "attestation")
        echo -e "\nChecking keys on $input machine\n"
        ls ~/Documents/celo-attestations-node/keystore
        echo ""
      ;;
    
      *) 
        echo -e "Invalid input\n" 
      ;;
    esac
}

function chainDataDel {
    case $input in
        "local")
          echo -e "\nDeleting chain data on $input machine\n"
          cd ~/Documents/celo-accounts-node
          sudo rm -rf geth* && sudo rm static-nodes.json
          echo ""
        ;;
        "validator")
          echo -e "\nDeleting chain data on $input machine\n"
          cd ~/Documents/celo-validator-node
          sudo rm -rf geth* && sudo rm static-nodes.json
          echo ""
        ;;
        "proxy")
          echo -e "\nDeleting chain data on $input machine\n"
          cd ~/Documents/celo-proxy-node
          mv geth/nodekey nodekey
          sudo rm -rf geth* && sudo rm static-nodes.json
          mkdir geth
          mv nodekey geth/nodekey
        ;;
        "attestation")
          echo -e "\nDeleting chain data on $input machine\n"
          cd ~/Documents/celo-attestations-node
          cd celo-accounts-node
          sudo rm -rf geth* && sudo rm static-nodes.json
          echo ""
        ;;
    
        *) 
          echo -e "Invalid input\n" 
        ;;
    esac
}

function nodeRestart {
    docker pull $CELO_IMAGE
    case $input in
      "local")
        echo -e "\nRestarting node on $input machine\n"
        docker run -v $PWD:/root/.celo --rm -it $CELO_IMAGE init /celo/genesis.json
        docker run -v $PWD:/root/.celo --rm -it --entrypoint cp $CELO_IMAGE /celo/static-nodes.json /root/.celo/
        docker run --name celo-accounts -dt --restart always -p 127.0.0.1:8545:8545 -v $PWD:/root/.celo $CELO_IMAGE --verbosity 3 --networkid $NETWORK_ID --syncmode full --rpc --rpcaddr 0.0.0.0 --rpcapi eth,net,web3,debug,admin,personal
        echo ""
      ;;
      "validator")
        echo -e "\nRestarting node on $input machine\n"
        cd ~/Documents/celo-validator-node
        docker run -v $PWD:/root/.celo --rm -it $CELO_IMAGE init /celo/genesis.json
        docker run --name celo-validator -dt --restart always -p 30303:30303 -p 30303:30303/udp -v $PWD:/root/.celo $CELO_IMAGE --verbosity 3 --networkid $NETWORK_ID --syncmode full --mine --istanbul.blockperiod=5 --istanbul.requesttimeout=3000 --etherbase $CELO_VALIDATOR_SIGNER_ADDRESS --nodiscover --proxy.proxied --proxy.proxyenodeurlpair=enode://$PROXY_ENODE@$PROXY_INTERNAL_IP:30503\;enode://$PROXY_ENODE@$PROXY_EXTERNAL_IP:30303  --unlock=$CELO_VALIDATOR_SIGNER_ADDRESS --password /root/.celo/.password --ethstats=$VALIDATOR_NAME@baklava-ethstats.celo-testnet.org
        echo ""
      ;;
      "proxy")
        echo -e "\nRestarting node on $input machine\n"
        cd ~/Documents/celo-proxy-node
        docker run -v $PWD:/root/.celo --rm -it $CELO_IMAGE init /celo/genesis.json
        export BOOTNODE_ENODES=`docker run --rm --entrypoint cat $CELO_IMAGE /celo/bootnodes`
        docker run --name celo-proxy -dt --restart always -p 30303:30303 -p 30303:30303/udp -p 30503:30503 -p 30503:30503/udp -v $PWD:/root/.celo $CELO_IMAGE --verbosity 3 --networkid $NETWORK_ID --syncmode full --proxy.proxy --proxy.proxiedvalidatoraddress $CELO_VALIDATOR_SIGNER_ADDRESS --proxy.internalendpoint :30503 --etherbase $CELO_VALIDATOR_SIGNER_ADDRESS --bootnodes $BOOTNODE_ENODES --ethstats=$VALIDATOR_NAME-proxy@baklava-ethstats.celo-testnet.org
      ;;
      "attestation")
        echo -e "\nPlease check here ($URL) for instructions on setting up an $input machine\n"
      ;;
      *) 
        echo -e "Invalid input\n" 
      ;;
    esac
}
# main

## TODO: https://docs.celo.org/getting-started/baklava-testnet/running-a-validator#register-the-accounts