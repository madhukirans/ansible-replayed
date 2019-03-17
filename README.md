### Replayed

The replayd application accepts user data over HTTP and stores it to an in-memory buffer. It also accepts requests for that data over HTTP and responds with what has been buffered.

Please look into Makefile for necessary targets.
 
### Ansible:
  This application is tested/deployed using Ansible using tow different environments with two nodes each.
  Node operating system is redhat linux 7
  User need to build application using `make go-build` before deploying using ansible
  
  Inventories
      Prod 
      Dev  
 
 Command to run ansible on prod env
 ``` ansible-playbook replayed.yaml -i inventories/prod/hosts ```
 Alternatively user can also run using `make`
 ```make prod-playbook```

Note: - id_rsa file is encrypted using ansible-vault and need to decrypt id_rsa file using ansible-vault. User needs password to decrypt the file.