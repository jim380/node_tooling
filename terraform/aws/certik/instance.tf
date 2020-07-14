resource "aws_key_pair" "main_key" {
  key_name   = "ctk"
  public_key = file(var.PATH_TO_PUBLIC_KEY)
}

resource "aws_instance" "dev_ctk" {
  ami           = var.AMIS[var.AWS_REGION]
  instance_type = var.NODE_INSTANCE_MODEL
  key_name      = aws_key_pair.main_key.key_name
  # security group
  vpc_security_group_ids = [aws_security_group.dev-ctk.id]

  # user data
  user_data = data.template_cloudinit_config.cloudinit.rendered

  root_block_device {
    volume_size = var.NODE_INSTANCE_VOLUME
    volume_type = "gp2"
    delete_on_termination = true
  }

  # get binaries ready
  # daemon
  provisioner "file" {
    source      = "files/certikd"
    destination = "/tmp/certikd"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo mkdir -p ~/Downloads ~/Documents",
      "sudo chmod +x /tmp/certikd",
      "sudo mv /tmp/certikd /usr/local/bin",
    ]
  }

  # cli
  provisioner "file" {
    source      = "files/certikcli"
    destination = "/tmp/certikcli"
  }

  provisioner "remote-exec" {
    inline = [
      "sudo chmod +x /tmp/certikcli",
      "sudo mv /tmp/certikcli /usr/local/bin",
    ]
  }
  
# create system service
  provisioner "file" {
    content = <<-EOF
      [Unit]
      Description=${var.protocol}d
      After=network-online.target

      [Service]
      User=ubuntu
      TimeoutStartSec=0
      LimitNOFILE=65535
      CPUWeight=90
      IOWeight=90
      Restart=always
      RestartSec=3
      ExecStart=/usr/local/bin/${var.protocol}d start
      KillSignal=SIGTERM
      StandardOutput=file:/var/log/${var.protocol}d.log
      StandardError=file:/var/log/${var.protocol}d.log

      [Install]
      WantedBy=multi-user.target
    EOF

    destination = "/tmp/${var.protocol}d.service"
  }

  provisioner "file" {
    content = <<-EOF
      [Unit]
      Description=LCD Server
      After=network.target ${var.protocol}d.service

      [Service]
      User=ubuntu
      TimeoutStartSec=0
      LimitNOFILE=65535
      CPUWeight=90
      IOWeight=90
      Restart=always
      RestartSec=3
      ExecStart=/usr/local/bin/${var.protocol}cli rest-server --laddr tcp://0.0.0.0:1317 --trust-node
      StandardOutput=file:/var/log/${var.protocol}cli.log
      StandardError=file:/var/log/${var.protocol}cli.log

      [Install]
      WantedBy=multi-user.target
    EOF

    destination = "/tmp/${var.protocol}cli.service"
  }

  # get genesis ready
  provisioner "file" {
    source      = "files/genesis.json"
    destination = "/tmp/genesis.json"
  }
  
  provisioner "file" {
    source      = "scripts/script.sh"
    destination = "/tmp/script.sh"
  }

  provisioner "remote-exec" {
    inline = [
      "chmod +x /tmp/script.sh",
      "sudo mv /tmp/script.sh ~/Downloads",
      "sudo sed -i -e 's/\r$//' ~/Downloads/script.sh",  # Remove the spurious CR characters.
    ]
  }

  provisioner "remote-exec" {
    inline = [
      "${var.protocol}d init ${var.moniker}",
      "sudo mv /tmp/genesis.json ~/.${var.protocol}d/config",
      "sudo ~/Downloads/script.sh", # add persistent peers
      "sudo mv /tmp/${var.protocol}d.service /etc/systemd/system",
      "sudo mv /tmp/${var.protocol}cli.service /etc/systemd/system",
      "sudo systemctl enable ${var.protocol}d",
      "sudo systemctl daemon-reload",
      "sudo systemctl start ${var.protocol}d",
      "sudo systemctl enable ${var.protocol}cli",
      "sudo systemctl daemon-reload",
      "sudo systemctl start ${var.protocol}cli",
    ]
  }

  connection {
    host        = coalesce(self.public_ip, self.private_ip)
    type        = "ssh"
    user        = var.INSTANCE_USERNAME
    private_key = file(var.PATH_TO_PRIVATE_KEY)
  }
}

