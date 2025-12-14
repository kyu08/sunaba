#!/bin/bash
set -euo pipefail
IFS=$'\n\t'


ln -s ${PWD}/.tmux.conf ~/.tmux.conf
ln -s ${PWD}/.vimrc ~/.vimrc
