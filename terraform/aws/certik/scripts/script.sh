#!/bin/bash

# sleep until instance is ready
until [[ -f /var/lib/cloud/instance/boot-finished ]]; do
  sleep 1
done

apt-get update
apt-get upgrade -y
sed -E -i 's/persistent_peers = \".*\"/persistent_peers = \"f30e998a6357fa5b8ed535111efccde8453c74d5@18.212.214.119:26656,7a036a8e868eb24982387f3462353781ae7a8a34@34.239.45.231:26656,b6890b022101f88015d353d93999b613668dcbd2@34.200.219.207:26656,888d8dedad47edbec4c389f1c0900b40fd713168@ec2-35-171-159-53.compute-1.amazonaws.com:26656,b7f8e66f1bb9f37f9cb93d04f11c481da0e955a5@ec2-3-230-166-163.compute-1.amazonaws.com:26656,a55e76a621f631a7a7aa5552530beafa86c59c38@ec2-3-209-82-112.compute-1.amazonaws.com:26656,276ecee8bc5dbc5d59b0e803c82ab75b011bce06@ec2-18-205-176-26.compute-1.amazonaws.com:26656,d830ade96349c0769616f63668a803242c90bfa9@ec2-18-206-92-147.compute-1.amazonaws.com:26656,b04eb82655a25e9470fb800cd661685a1f4ebd89@ec2-35-175-112-147.compute-1.amazonaws.com:26656,75768fca7957c0b60cca23ea7cd8c7e8e87e9c44@ec2-3-236-9-204.compute-1.amazonaws.com:26656,1c5550c131c1d1ec747e8ecb5b932c7cf306115d@ec2-3-229-124-35.compute-1.amazonaws.com:26656,de4c2266a6a6255585f8783b043baa7344d15abb@ec2-3-234-212-182.compute-1.amazonaws.com:26656,9ccb2643e277ecd6586c6d0cfe09e205169018a1@ec2-100-24-209-16.compute-1.amazonaws.com:26656,5036cdb146af9959db3c6c5ba91fc56692a8644d@ec2-100-25-42-35.compute-1.amazonaws.com:26656,9e0f60199948b89f44f0fdb99de17f4c09284989@ec2-3-236-83-78.compute-1.amazonaws.com:26656,02d5de65fc250b8d05b32927de5af90b23c8f5c2@ec2-3-235-191-31.compute-1.amazonaws.com:26656,75cfdcd2673ddb3b33a8326f94e657a5a616d538@ec2-3-233-224-51.compute-1.amazonaws.com:26656,436ecfbf312d43f2852ab166c4466375898bac43@ec2-3-235-185-129.compute-1.amazonaws.com:26656\"/' $HOME/.${var.protocol}d/config
