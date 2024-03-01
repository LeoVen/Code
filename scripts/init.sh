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

# initial tools
sudo apt install build-essential procps curl file git gnupg jq sed net-tools vim -y

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
# golang
# remove previous version
rm -rf /usr/local/go
# download new version
wget https://go.dev/dl/go1.21.3.linux-amd64.tar.gz
# extract downloaded version
sudo tar -C /usr/local -xzf go1.21.3.linux-amd64.tar.gz
# add to env var
echo "\nexport PATH=\$PATH:/usr/local/go/bin:$HOME/go/bin" >> ~/.bashrc

# devops
# terraform
wget -O- https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
sudo apt update && sudo apt install terraform -y

# other tools
sudo apt install zip unzip ffmpeg graphviz openssl bc synaptic -y

# 1password cli
sudo apt install 1password-cli

# install homebrew
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
# add it to path
test -d ~/.linuxbrew && eval "$(~/.linuxbrew/bin/brew shellenv)"
test -d /home/linuxbrew/.linuxbrew && eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"
echo "eval \"\$($(brew --prefix)/bin/brew shellenv)\"" >> ~/.bashrc

# some nice aliases
alias hg='history | grep'
alias .1='cd ../'
alias .2='cd ../../'
alias .3='cd ../../../'
alias .4='cd ../../../../'
alias .5='cd ../../../../..'
