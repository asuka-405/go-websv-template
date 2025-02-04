
function install_air() {
  go install github.com/air-verse/air@latest
  // air init
}

function install_docker_archlinux() {
  sudo pacman -S docker
  sudo systemctl enable docker
  sudo systemctl start docker
  sudo usermod -aG docker $USER
  sudo echo "net.ipv4.ip_forward=1" >> /etc/sysctl.conf
  sudo sysctl --system
}