#!/bin/bash
#                                                                                                         
#                                                  jim380 <admin@cyphercore.io>
#  ============================================================================
#  
#  Copyright (C) 2019 jim380
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
GR_VERSION="6.4.2" #old 6.4.1
PR_VERSION="2.13.0" #old 2.12.0
NE_VERSION="0.18.1" #old 0.17.0
OS=`uname -s`
ARCH=`uname -m`

function main {
        echo -e "What would you like to install:\n\n1) Grafana\n2) Prometheus\n3) Node Exporter\n4) All\n\nEnter down below (e.g. \"1\" or \"Grafana\"):"
        read input

        case $input in
          "1" | "Grafana")
            echo -e "\nWhat version of Grafana would you like to install? Example: v6.4.2\nFind releases here: https://github.com/grafana/grafana/tags\n"
            read -p 'v' GR_VERSION
            installGrafana
          ;;

          "2" | "Prometheus")
            echo -e "\nWhat version of Prometheus would you like to install? Example: v2.13.0\nFind releases here: https://github.com/prometheus/prometheus/tags\n"
            read -p 'v' PR_VERSION
            installPrometheus
          ;;

          "3" | "Node Exporter")
            echo -e "\nWhat version of Node Exporter would you like to install? Example: v0.18.1\nFind releases here: https://github.com/prometheus/node_exporter/tags\n"
            read -p 'v' NE_VERSION
            installNodeExporter
          ;;

          "4" | "All")
            echo -e "\nWhat versions would you like to install? Find releasese here:\n"
            echo -e "Grafana: https://github.com/grafana/grafana/tags\nPrometheus: https://github.com/prometheus/prometheus/tags\nNode Exporter: https://github.com/prometheus/node_exporter/tags\n"
            read -p 'Grafana: v' GR_VERSION
            read -p 'Prometheus: v' PR_VERSION
            read -p 'Node Exporter: v' NE_VERSION
            installGrafana
            installPrometheus
            installNodeExporter
          ;;
        
          *) 
            echo -e "Invalid input - $input\n" 
          ;;
        esac
}

function installGrafana {
    echo "-----------------------------------------"
    echo "              Install Grafana            "
    echo "-----------------------------------------"
    if ! test -d $HOME/Downloads
    then
        mkdir -p $HOME/Downloads
    fi

    # 64-bit Linux
    if [ "$OS" == "Linux" ] && [ "$ARCH" == "x86_64" ]
    then
        PACKAGE_GR=grafana_$GR_VERSION\_amd64.deb
        pushd ~/Downloads #> /dev/null
            echo "-----------------------------------------"
            echo "                Downloading              "
            echo "-----------------------------------------"
            wget https://dl.grafana.com/oss/release/$PACKAGE_GR
            if [ $? -ne 0 ]; then 
                echo "Failed to download !!!"
                exit 1
            fi
            sudo apt-get install -y adduser libfontconfig1
            echo "-----------------------------------------"
            echo "              Install Package            "
            echo "-----------------------------------------"
            sudo dpkg -i $PACKAGE_GR
            echo "-----------------------------------------"
            echo "               Remove Package            "
            echo "-----------------------------------------"
            rm -rf $PACKAGE_GR
            echo "-----------------------------------------"
            echo "                Start Server             "
            echo "-----------------------------------------"
            sudo systemctl daemon-reload
            sudo systemctl start grafana-server
            sudo systemctl enable grafana-server.service
        popd
        cd ~
        success
        exit 0
    fi
}

function installPrometheus {
    echo "-----------------------------------------"
    echo "            Install Prometheus           "
    echo "-----------------------------------------"
    if ! test -d $HOME/Downloads
    then
        mkdir -p $HOME/Downloads
    fi

    # 64-bit Linux
    if [ "$OS" == "Linux" ] && [ "$ARCH" == "x86_64" ]
    then
        PACKAGE_PR=prometheus-$PR_VERSION.linux-amd64.tar.gz
        pushd ~/Downloads #> /dev/null
            echo "-----------------------------------------"
            echo "                Downloading              "
            echo "-----------------------------------------"
            wget https://github.com/prometheus/prometheus/releases/download/v$PR_VERSION/$PACKAGE_PR
            if [ $? -ne 0 ]; then 
                echo "Failed to download !!!"
                exit 1
            fi
            echo "-----------------------------------------"
            echo "              Extract Package            "
            echo "-----------------------------------------"
            tar xvfz $PACKAGE_PR
            echo "-----------------------------------------"
            echo "               Remove Package            "
            echo "-----------------------------------------"
            rm -rf $PACKAGE_PR
        popd
        success
    fi
}

function installNodeExporter {
    echo "-----------------------------------------"
    echo "          Install Node Exporter          "
    echo "-----------------------------------------"
    if ! test -d $HOME/Downloads
    then
        mkdir -p $HOME/Downloads
    fi

    # 64-bit Linux
    if [ "$OS" == "Linux" ] && [ "$ARCH" == "x86_64" ]
    then
        PACKAGE_NE=node_exporter-$NE_VERSION.linux-amd64.tar.gz
        pushd ~/Downloads
            echo "-----------------------------------------"
            echo "                Downloading              "
            echo "-----------------------------------------"
            wget https://github.com/prometheus/node_exporter/releases/download/v$NE_VERSION/$PACKAGE_NE
            if [ $? -ne 0 ]; then 
                echo "Failed to download !!!"
                exit 1
            fi
            echo "-----------------------------------------"
            echo "              Extract Package            "
            echo "-----------------------------------------"
            tar xvfz $PACKAGE_NE
            echo "-----------------------------------------"
            echo "               Remove Package            "
            echo "-----------------------------------------"
            rm -rf $PACKAGE_NE
        popd
        success
    fi
}

function success {
    echo "-----------------------------------------"
    echo "                 Success                 "
    echo "-----------------------------------------"
    echo "Task has successfully been completed."
}

main
