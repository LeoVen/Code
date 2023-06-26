# everything to get started in ubuntu
sudo apt update -y
sudo apt upgrade -y

# terminal
sudo apt install zsh -y
# setup zsh
# install oh-my-zsh plugins
sh -c "$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)"
# install zsh theme power10k
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ~/powerlevel10k
echo 'source ~/powerlevel10k/powerlevel10k.zsh-theme' >> ~/.zshrc

# programming languages
# C, C++
sudo apt install gcc -y
# rust
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
# setup rust
echo "\nexport CARGO_BUILD_TARGET_DIR=~/.target" >> ~/.bashrc
source "$HOME/.cargo/env"
# install nvm
curl https://raw.githubusercontent.com/creationix/nvm/master/install.sh | bash
source ~/.bashrc
# install node
nvm install node
# python
sudo apt install python3 python3-pip -y

# devops
# terraform
wget -O- https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
sudo apt update && sudo apt install terraform -y

# other tools
sudo apt install zip unzip vim -y
