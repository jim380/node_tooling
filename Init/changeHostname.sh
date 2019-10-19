#!/bin/bash
echo "-----------------------------------------"
echo "             Change Hostname             "
echo "-----------------------------------------"
#Assign existing hostname to $hostn
hostn=$(cat /etc/hostname)

#Display existing hostname
echo "Existing hostname is $hostn"

#Ask for new hostname $newhost
echo "Enter a new hostname: "
read newhost

#change hostname in /etc/hosts & /etc/hostname
sudo sed -i "s/$hostn/$newhost/g" /etc/hosts
sudo sed -i "s/$hostn/$newhost/g" /etc/hostname

#display new hostname
echo "Hostname is now set to $newhost"

#Press a key to reboot
read -s -n 1 -p "Press any key to reboot"
sudo reboot