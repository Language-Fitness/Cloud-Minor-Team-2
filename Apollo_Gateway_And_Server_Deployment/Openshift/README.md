# OpenShift Deployment

Contains assets folder with all YAML files for deploying the resources needed to run gateway
and servers in openshift. Deployment is done using the file "kustomization.yml".

Run following command in corresponding folder to deploy:
```console
..\Openshift\kustomization: oc apply -k .
```