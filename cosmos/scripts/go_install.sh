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
GO_VERSION="1.11.5"
OS=`uname -s`
# HOME_DIR=$HOME
# GO_HOME=$HOME_DIR/go
GO_ROOT=/usr/local/go
ARCH=`uname -m`

function usage {
    printf "./go_install.sh -v <version> \n"
    printf "Example: ./go_install.sh -v 1.11.5 \n"
    exit 1
}

while getopts ":v:" opt; do
  case $opt in
    v) GO_VERSION="$OPTARG"
    ;;
    \?) echo "Invalid option -$OPTARG" >&2
    echo 
    usage
    ;;
  esac
done 

function preparation {
    echo "-----------------------------------------"
    echo "               System Update             "
    echo "-----------------------------------------"
    sudo apt-get update && sudo apt-get upgrade -y 
    echo "-----------------------------------------"
    echo "            Necessary Packages           "
    echo "-----------------------------------------"
    sudo apt-get install htop git curl bash-completion jq -y
    sudo apt-get install libgmp3-dev -y
}

function scan {
echo "-----------------------------------------"
echo "   Scan if an old version was installed   "
echo "-----------------------------------------"
if [ -d /usr/local/go ]
then 
    echo "An older version was installed."
    read -p "Would you like to remove it? [y/n]: " ans
    case "$ans" in 
        "y"|"yes"|"Y"|"Yes"|"YES") sudo rm -rf /usr/local/go
        echo "Old version successfully removed."
        ;;
        *) echo "Exiting..."
           exit 0
        ;;
    esac
fi
echo "No old version was found. Proceed."
}

function install {
    echo "-----------------------------------------"
    echo "              Install Golang             "
    echo "-----------------------------------------"
    if ! test -d $HOME/Downloads
    then
        mkdir -p $HOME/Downloads
    fi
    #cd $HOME/Downloads

    # 64-bit Linux
    if [ "$OS" == "Linux" ] && [ "$ARCH" == "x86_64" ]
    then
        PACKAGE=go$GO_VERSION.linux-amd64.tar.gz
        pushd ~/Downloads #> /dev/null
            echo "-----------------------------------------"
            echo "                Downloading              "
            echo "-----------------------------------------"
            wget https://dl.google.com/go/$PACKAGE
            if [ $? -ne 0 ]; then 
                echo "Failed to download !!!"
                exit 1
            fi
            echo "-----------------------------------------"
            echo "              Extract Package            "
            echo "-----------------------------------------"
            sudo tar -C /usr/local -xzf $PACKAGE
            rm -rf $PACKAGE
        popd #> /dev/null
        cd ~
        permission
        setup
        success
        exit 0
    fi

    # 64-bit MacOS
    if [ "$OS" == "Darwin" ] && [ "$ARCH" == "x86_64" ]
    then 
        PACKAGE=go$GO_VERSION.darwin-amd64.pkg
        pushd ~/Downloads #> /dev/null
            echo "-----------------------------------------"
            echo "                Downloading              "
            echo "-----------------------------------------"
            wget wget https://dl.google.com/go/$PACKAGE
            if [ $? -ne 0 ]; then 
                echo "Failed to download !!!"
                exit 1
            fi
            echo "-----------------------------------------"
            echo "              Extract Package            "
            echo "-----------------------------------------"
            sudo /usr/sbin/installer -pkg $PACKAGE -target /
            rm -rf $PACKAGE
        popd #> /dev/null
        cd ~
        permission
        setup
        success
        exit 0
    fi

    # ARM
    if [ "$OS" == "Linux" ] && [ "$ARCH" == "armv7l" ]
    then
        PACKAGE=go$GO_VERSION.linux-armv6l.tar.gz
        pushd ~/Downloads #> /dev/null
            echo "-----------------------------------------"
            echo "                Downloading              "
            echo "-----------------------------------------"
            wget https://dl.google.com/go/$PACKAGE
            if [ $? -ne 0 ]; then 
                echo "Failed to download !!!"
                exit 1
            fi
            echo "-----------------------------------------"
            echo "              Extract Package            "
            echo "-----------------------------------------"
            sudo tar -C /usr/local -xzf $PACKAGE
            rm -rf $PACKAGE
            echo "Extracting done !"
        popd #> /dev/null
        cd ~
        permission
        setup
        success
        exit 0
    fi
        
    errorout
}

function permission {
    echo "-----------------------------------------"
    echo "              Set Permissions            "
    echo "-----------------------------------------"
    sudo chown root:root /usr/local/go
    sudo chmod 755 /usr/local/go
    echo "Permissions set !"
}

function setup {
    echo "-----------------------------------------"
    echo "             Set Go Workspace            "
    echo "-----------------------------------------"
    echo "Where would you like your Go Workspace folder to be? (example: /home)"
    read -p "Path: " GO_WS_PATH
    cd $GO_WS_PATH
    read -p "Give the folder a name: " GO_WS_NAME
    GO_PATH=$PWD/$GO_WS_NAME
    echo "Your Go Workspace folder has been set to $GO_PATH"
    mkdir -p $GO_WS_NAME{,/bin,/pkg,/src}
    cd ~
    echo "-----------------------------------------"
    echo "             Set Env Variables           "
    echo "-----------------------------------------"
    sudo sh -c "echo 'export PATH=\$PATH:/usr/local/go/bin' >> /etc/profile"
    echo "export GOPATH=$GO_PATH" >> ~/.profile
    echo "export PATH=$GO_PATH/bin:\$PATH" >> ~/.profile
    echo "Setup done !"
}

function success {
    echo "-----------------------------------------"
    echo "                 Success                 "
    echo "-----------------------------------------"
    echo "Go has been successfully installed.
    Reboot the system to take effects."
}
function errorout {
    echo "OS/ARCH not currently supported"
    exit 1
}

preparation
scan
install
permission
setup