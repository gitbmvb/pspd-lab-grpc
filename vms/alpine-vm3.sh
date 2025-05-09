#!/bin/bash

VM_NAME="alpine-vm3"

# Start the VM
echo "Starting $VM_NAME..."
virsh start "$VM_NAME"

# Wait a moment for it to boot (optional)
sleep 2

# Attach to the console
echo "Connecting to $VM_NAME console (press Ctrl+] to exit)..."
virsh console "$VM_NAME"

