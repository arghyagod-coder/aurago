echo "Downloading aura..."
wget "https://github.com/arghyagod-coder/GoLang/releases/download/1.0.0/aura-linux-amd64" 

echo "Adding aura into PATH..."

mkdir -p ~/.aura

mv ./aura-linux-amd64 ~/.aura/aura
chmod +x ~/.aura/aura
set -U fish_user_paths ~/.aura/ $fish_user_paths 
echo "aura installation is completed!"
echo "You need to restart the shell to use aura." 