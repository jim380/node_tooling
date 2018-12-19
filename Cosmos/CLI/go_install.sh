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
echo "            Necessary Packages           "
echo "-----------------------------------------"
sudo apt-get install htop git curl bash-completion jq -y
sudo apt-get install libgmp3-dev -y
echo "-----------------------------------------"
echo "              Download Golang            "
echo "-----------------------------------------"
if ! test -d $HOME/Downloads
then
    mkdir -p $HOME/Downloads
fi
cd $HOME/Downloads
read -p "Link to package to be downloaded
" GO_LINK
echo ""
wget $GO_LINK
echo "You can find the downloaded package here:
$HOME/Downloads"
echo "-----------------------------------------"
echo "              Extract Package            "
echo "-----------------------------------------"
read -p "Paste in the name of the package file downloaded:
" GO_FILE
echo ""
sudo tar -C /usr/local -xvf $GO_FILE
echo "-----------------------------------------"
echo "          Set Folder Permissions         "
echo "-----------------------------------------"
sudo chown root:root /usr/local/go
sudo chmod 755 /usr/local/go
cd ~
echo "-----------------------------------------"
echo "             Set Go Workspace            "
echo "-----------------------------------------"
read -p "Where would you like your Go Workspace folder to be?
Path: " GO_WS_PATH
cd $GO_WS_PATH
read -p "Give the folder a name: " GO_WS_NAME
GO_PATH=$PWD/$GO_WS_NAME
echo "Your Go Workspace folder has been set to $GO_PATH"
mkdir -p $GO_WS_NAME{,/bin,/pkg,/src}
echo "Success"
cd ~
echo "-----------------------------------------"
echo "             Set Env Variables           "
echo "-----------------------------------------"
sudo sh -c "echo 'export PATH=$PATH:/usr/local/go/bin' >> /etc/profile"
echo "export GOPATH=$GO_PATH" >> ~/.profile
echo "export PATH=$GOPATH/bin:$PATH" >> ~/.profile
echo "export GOPATH=$GO_PATH" >> ~/.bash_profile
echo "export GOBIN=$GO_PATH/bin" >> ~/.bash_profile
echo "export PATH=$PATH:$GOBIN" >> ~/.bash_profile
echo "Success"
echo "-----------------------------------------"
echo "                  Reboot                 "
echo "-----------------------------------------"
echo "Go has been successfully installed.
Reboot the system to take effects."