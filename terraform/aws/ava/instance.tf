resource "aws_key_pair" "main_key" {
  key_name   = "terra"
  public_key = file(var.PATH_TO_PUBLIC_KEY)
}

resource "aws_instance" "ava_node" {
  ami           = var.AMIS[var.AWS_REGION]
  instance_type = var.NODE_INSTANCE_MODEL
  key_name      = aws_key_pair.main_key.key_name
  # security group
  vpc_security_group_ids = [aws_security_group.dev-default.id]
  # user data
  user_data = data.template_cloudinit_config.cloudinit.rendered

  root_block_device {
    volume_size = var.NODE_INSTANCE_VOLUME
    volume_type = "gp2"
    delete_on_termination = true
  }

  # get binary ready
  provisioner "file" {
    source      = "binary/ava"
    destination = "/tmp/ava"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo mkdir -p ~/Downloads ~/Documents",
      # "sudo mv ~/go_install.sh ~/Downloads"
      # "sudo chmod +x ~/Downloads/go_install.sh",
      # "sudo ~/go_install.sh -v 1.14.3",
      "sudo mv /tmp/ava ~/Downloads/ava",
      "sudo chmod +x ~/Downloads/ava",
    ]
  }

  # compile gecko
  # provisioner "remote-exec" {
  #   inline = [
  #     "go get -v -d github.com/ava-labs/gecko/...",
  #     "cd $GOPATH/src/github.com/ava-labs/gecko && ./scripts/build.sh",
  #     "sudo mv ./build/ava $GOPATH/bin",
  #   ]
  # }
  
  # create system service
  provisioner "file" {
    content = <<-EOF
      [Unit]
      Description=AVA Gecko
      After=network-online.target

      [Service]
      User=ubuntu
      TimeoutStartSec=0
      LimitNOFILE=65535
      CPUWeight=90
      IOWeight=90
      Restart=always
      ExecStart=/home/ubuntu/Downloads/ava
      KillSignal=SIGTERM
      StandardOutput=file:/var/log/ava.log
      StandardError=file:/var/log/ava.log

      [Install]
      WantedBy=multi-user.target
    EOF

    destination = "/tmp/ava.service"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo mv /tmp/ava.service /etc/systemd/system",
      "sudo systemctl enable ava",
      "sudo systemctl daemon-reload",
      "sudo systemctl start ava",
    ]
  }

  connection {
    host        = coalesce(self.public_ip, self.private_ip)
    type        = "ssh"
    user        = var.INSTANCE_USERNAME
    private_key = file(var.PATH_TO_PRIVATE_KEY)
  }
}

