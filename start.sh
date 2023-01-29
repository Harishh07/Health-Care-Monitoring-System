#!/bin/bash

fn=$1

session="work"

# set up tmux
tmux start-server

# create a new tmux session, starting vim from a saved session in the new window
tmux new-session -d -s $session "./sink/sink  -pda '10.6.50.85:64753' " #"vim -S ~/.vim/sessions/kittybusiness"
sleep 1

tmux split-window -h "./goreman -f Procfile3 start "
#tmux split-window -h 'python'


tmux -2 attach-session -d
