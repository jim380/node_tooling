#!/bin/bash
# Give non-root User with Sudo Access
echo "-----------------------------------------"
echo "               Add New User              "
echo "-----------------------------------------"
if [ $(id -u) -eq 0 ]; then
	read -p "Enter username : " username
	#read -s -p "Enter password : " password
	egrep "^$username" /etc/passwd >/dev/null
	if [ $? -eq 0 ]; then
		echo "$username already exists. Please pick a new one."
		exit 1
	else
		# pass=$(perl -e 'print crypt($ARGV[0], "password")' $password)
		# echo $pass

        useradd -m $username
		
		echo "-----------------------------------------"
		echo "                Add Passwd               "
		echo "-----------------------------------------"
		sudo passwd $username

		echo "-----------------------------------------"
		echo "    Give /home/<username> Permissions    "
		echo "-----------------------------------------"
        chown $username /home/$username -R

        echo "-----------------------------------------"
		echo "        Add Passwd to /etc/passwd        "
		echo "-----------------------------------------"
		etcPasswd="$username:x:1000:1000::/home/$username:/bin/bash"
		sed -i '$a'$etcPasswd'' /etc/passwd

		echo "-----------------------------------------"
		echo "              Add Sudo Entry             "
		echo "-----------------------------------------"
		echo "$username    ALL=(ALL:ALL) ALL" | sudo EDITOR='tee -a' visudo

		[ $? -eq 0 ] && echo "User has been added to system!" || echo "Failed to add a user!"
	fi
else
	echo "Only users with root privileges may add a user to the system"
	exit 2
fi

echo "-----------------------------------------"
echo "                  Reboot                 "
echo "-----------------------------------------"
read -s -n 1 -p "Press any key to reboot"
sudo reboot
