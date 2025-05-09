sudo qemu-img create -f qcow2 /var/lib/libvirt/images/alpine-vmX.qcow2 KG
Criação e configuração de máquinas virtuais Alpine (Obs: K é meramente ilustrativo substitua pelo espaço desejado para reservar para a vm.)

Para criar uma máquina virtual utilizando o Alpine Linux, utilize o seguinte comando (substitua X e ? pelos valores apropriados):

sudo virt-install \
  --name alpine-vmX \
  --ram ? \
  --vcpus=1 \
  --os-variant=alpinelinux3.19 \
  --network bridge=br-lan,model=virtio \
  --network bridge=virbr0,model=virtio \
  --graphics none \
  --console pty,target_type=serial \
  --cdrom /var/lib/libvirt/images/alpine-standard-3.19.1-x86_64.iso \
  --disk path=/var/lib/libvirt/images/alpine-vmX.qcow2,format=qcow2

Após realizar esse processo para cada VM, é necessário configurar o sistema com o comando setup-alpine. Siga as instruções apresentadas na tela e, ao final, reinicie a máquina com o comando reboot.

Depois que o sistema reiniciar, entre novamente na máquina com as configurações criadas. Finalize a configuração e reinicie novamente.

Em seguida, execute o seguinte comando:

virsh dumpxml alpine-vmX > vmX.xml

Esse comando cria um arquivo XML com as configurações da máquina vmX, que será utilizado pelos scripts alpine-vmX.sh, previamente criados.

A partir deste ponto, você poderá reiniciar a VM sempre que necessário. Note que as configurações foram salvas, e você poderá configurar o ambiente normalmente, utilizando comandos apk para instalar o git e outras dependências necessárias.
