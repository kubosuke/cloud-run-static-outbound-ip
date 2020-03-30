#!/usr/bin/env bash

set -ex

ssh -i ./ssh_key "tunnel@${GCE_IP?:GCE_IP environment variable not set}" \
    -N -D localhost:5000 \
    -o StrictHostKeyChecking=no &
    # -N: non-execute any commands at target host
    # -D: create Tunnel
    # -o: define Option(like sshd.conf)
./app

# wait -n helps us exit immediately if one of the processes above exit.
# this way, Cloud Run can restart the container to be healed.
wait -n
