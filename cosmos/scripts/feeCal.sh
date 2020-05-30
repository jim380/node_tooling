#!/bin/bash
function fee {
        read -p "What denom would you like to pay in? (1:stake, 2:photinos): " ans
        case "$ans" in 
                "1"|"stake"|"1:stake") denom="stake"
                feeCal
                ;;
                "2"|"photinos"|"photino"|"2:photinos") denom="photinos"
                feeCal
                ;;
                *) echo "Invalid input. Exiting..."
                exit 0
                ;;
        esac
}

function feeCal {
        read -p "gasLimit (default: 200000): " gasLimit
        read -p "gasPrice (eg: 0.001$denom, enter amount in decimal ONLY): " gasPrice
        result=`python -c "from math import ceil; print ceil($gasLimit/(1/$gasPrice))"`
        echo $result
}

fee