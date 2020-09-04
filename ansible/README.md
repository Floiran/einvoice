# Ansible scripts

Prepare environment:

```shell script
export GCP_AUTH_KIND=serviceaccount
export GCP_SERVICE_ACCOUNT_FILE=/home/filip/mfsr/webserver1-283520-386230d8738c.json
export GCP_SCOPES=https://www.googleapis.com/auth/compute,https://www.googleapis.com/auth/cloud-platform
```

Create development environment:
```shell script
ansible-playbook dev.yaml
```
Clean development environment:
```shell script
ansible-playbook dev-clean.yaml
```
Create production environment:
```shell script
ansible-playbook prod.yaml
```
Clean production environment:
```shell script
ansible-playbook prod-clean.yaml
```
