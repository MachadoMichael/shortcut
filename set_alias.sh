#!/bin/bash

# Add an alias to your zsh configuration file
echo "alias sc='$(pwd)/bin/shortcut'" >> ~/.zshrc

# Reload the zshrc to apply the alias immediately
source ~/.zshrc

echo "Alias 'sc' created and added to your zsh environment."

