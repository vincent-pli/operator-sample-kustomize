# operator-sample-kustomize

This repo is sample for `Operator` build by [operator-sdk](https://github.com/operator-framework/operator-sdk), based on [kustomize](https://github.com/kubernetes-sigs/kustomize)


## Development Prerequisites
1. [`go`](https://golang.org/doc/install): The language Tektoncd-pipeline-operator is
   built in
2. [`git`](https://help.github.com/articles/set-up-git/): For source control
3. [`kubectl`](https://kubernetes.io/docs/tasks/tools/install-kubectl/): For
   interacting with your kube cluster
4. operator-sdk: https://github.com/operator-framework/operator-sdk
5. [kustomize](https://github.com/kubernetes-sigs/kustomize)
  
  
## Concepts
- **Template**:  
`Template` are `yaml` files by `Kustimize`, they are k8s resources which make up an application.  
The default location of `Template` is `./templates`.  
`Tempalte` could support multiple components.  

- **Extension**  
Something like plugin, could be remove/add easily.  
For now, there are 3 buildin `Extension`s:  
  - `nsinject`: if replace/add `namespace` for created sources.  

  - `ownerset`: if set the owner for created sources.  

  - `imagereplacement`: if replace the `url` of image in `Container` definition, use for private cluster case.

## Building the Operator Image
1. Enable go mod  

    `export GO111MODULE=on`
    
2. Build go and the container image  

    `operator-sdk build ${YOUR_REGISTORY}/operator-sample:${IMAGE_TAG}`
    
3. Push the container image  

    `docker push ${YOUR-REGISTORY}/operator-sample:${IMAGE-TAG}`
    
4. Edit the 'image' value in deploy/operator.yaml to match to your image 


## Take a try

The following steps will install sample application for your cluster.
1. Create namespace: `operator-sample`  

    `kubectl create namespace operator-sample`
    
2. Apply Operator CRD

    `kubectl apply -f deploy/crds/install_example.com_installs_crd.yaml`  
    
3. Deploy the Operator  

    `kubectl -n operator-sample apply -f deploy/`  

## The CRD
This is a sample of [crd](https://github.com/vincent-pli/operator-sample-kustomize/blob/master/deploy/crds/install.example.com_v1alpha1_install_cr.yaml)
```
apiVersion: install.example.com/v1alpha1
kind: Install
metadata:
  name: example-install
spec:
  # Add fields here
  targetNamespace: operator-sample
  setowner: true
  registry: 
    override:
      the-container: monopole/hello:1
```
It means:  
  1. Install application defined in `./templates`
  
  2. Put all the resource of the application to `namespace`: operator-sample
  
  3. Set the owner of all the resource to the CR
  
  4. Replace the image url to `monopole/hello:1` of container who's name is "the-container"
