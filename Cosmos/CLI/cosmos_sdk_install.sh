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
echo "               System Update             "
echo "-----------------------------------------"
sudo apt-get update && sudo apt-get upgrade -y 
echo "-----------------------------------------"
echo "                Preparation              "
echo "-----------------------------------------"
sudo apt-get install make gcc g++ -y
sudo sudo apt autoremove
echo "-----------------------------------------"
echo "             Install Binaries            "
echo "-----------------------------------------"
if ! test -d $GOPATH/src/github.com/cosmos
then
    mkdir -p $GOPATH/src/github.com/cosmos
    cd $GOPATH/src/github.com/cosmos
    git clone https://github.com/cosmos/cosmos-sdk
fi
echo "-----------------------------------------"
echo "             Update Binary               "
echo "-----------------------------------------"
REPO=$GOPATH/src/github.com/cosmos/cosmos-sdk
if [ -d "$REPO" ]
then
    cd $REPO
    if [ -d "$REPO/.git" ]
    then
      echo "Updating $REPO at
`date`"
      git status
      echo "-----------------------------------------"
      echo "               Fetching                  "
      echo "-----------------------------------------"
      git fetch
      echo "-----------------------------------------"
      echo "                Pulling                  "
      echo "-----------------------------------------"
      git pull
    else
      echo "-----------------------------------------"
      echo "This is not a git folder."
      echo "-----------------------------------------"
    fi
    echo "Finished updating at
`date`"
    echo ""
fi
echo "-----------------------------------------"
echo "                 Checkout                "
echo "-----------------------------------------"
read -p "What version would you like to checkout?
Enter 'master' or specify a version number (e.g. 'v0.29.0')
" CHECKOUT_VERSION
echo "Installing $CHECKOUT_VERSION"
git checkout $CHECKOUT_VERSION
echo "-----------------------------------------"
echo "              Make & Install             "
echo "-----------------------------------------"
make get_tools && make get_vendor_deps && make install
echo "-----------------------------------------"
echo "            Fetch genesis.json           "
echo "-----------------------------------------"
read -p "Link to genesis.json in raw format:
" GENESIS
echo ""
if ! test -d $HOME/.gaiad/config
then
    mkdir -p $HOME/.gaiad/config
fi
curl $GENESIS > $HOME/.gaiad/config/genesis.json
echo "-----------------------------------------"
echo "               gaiad version             "
echo "-----------------------------------------"
gaiad version
echo "-----------------------------------------"
echo "              gaiacli version"
echo "-----------------------------------------"
gaiacli version
