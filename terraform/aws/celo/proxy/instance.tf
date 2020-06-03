resource "aws_key_pair" "main_key" {
  key_name   = "terra"
  public_key = file(var.PATH_TO_PUBLIC_KEY)
}

resource "aws_instance" "celo_proxy" {
  ami           = var.AMIS[var.AWS_REGION]
  instance_type = "t2.micro"
  key_name      = aws_key_pair.main_key.key_name
  # security group
  #security_groups = ["aws_security_group.allow-ssh.id"]
  # user data
  user_data = data.template_cloudinit_config.cloudinit.rendered

  root_block_device {
    volume_size = 500
    volume_type = "gp2"
    delete_on_termination = true
  }

  provisioner "file" {
    source      = "scripts/setEnv.sh"
    destination = "/tmp/setEnv.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo mkdir -p ~/Downloads ~/Documents",
      "sudo mv /tmp/setEnv.sh ~/Downloads/setEnv.sh",
      "cd ~/Downloads && sudo chmod +x setEnv.sh",
      "./setEnv.sh",
    ]
  }

  provisioner "remote-exec" {
    inline = [
      # in progress !!!
      "sudo mkdir -p ~/Documents/celo-proxy-node",
      "cd ~/Documents/celo-proxy-node",
      "docker pull ${var.CELO_IMAGE}",
      "docker run -v $PWD:/root/.celo --rm -it ${var.CELO_IMAGE} init /celo/genesis.json",
      # export BOOTNODE_ENODES="$(docker run --rm --entrypoint cat $CELO_IMAGE /celo/bootnodes)"
      "sudo echo testpassword >> .password",
      # export PROXY_ADDRESS="$(docker run --name celo-proxy-password -it --rm  -v $PWD:/root/.celo $CELO_IMAGE account new --password /root/.celo/.password)"
      # docker run --name celo-proxy -dt --restart unless-stopped -p 30303:30303 -p 30303:30303/udp -p 30503:30503 -p 30503:30503/udp -v $PWD:/root/.celo $CELO_IMAGE --verbosity 3 --networkid $NETWORK_ID --nousb --syncmode full --proxy.proxy --proxy.proxiedvalidatoraddress $CELO_VALIDATOR_SIGNER_ADDRESS --proxy.internalendpoint :30503 --etherbase $PROXY_ADDRESS --unlock $PROXY_ADDRESS --password /root/.celo/.password --allow-insecure-unlock --bootnodes $BOOTNODE_ENODES --ethstats=test@baklava-celostats-server.celo-testnet.org"
    ]
  }

  connection {
    host        = coalesce(self.public_ip, self.private_ip)
    type        = "ssh"
    user        = var.INSTANCE_USERNAME
    private_key = file(var.PATH_TO_PRIVATE_KEY)
  }
}

output "ip" {
  value = aws_instance.celo_proxy.public_ip
}