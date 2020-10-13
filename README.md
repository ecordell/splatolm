# splatolm

`splatolm` takes as input a CSV file, and outputs the set of RBAC resources that OLM would generate for it on cluster.

## Install

```sh
$ git clone https://github.com/ecordell/splatolm.git && cd splatolm && go build && go install
```

## Usage

```sh
$ kubectl -n openshift-storage get csv ocs-operator.v4.5.0 -o yaml > ocs-csv.yaml
$ splatolm ocs-csv.yaml > rbac.yaml
$ kubectl -n openshift-storage apply -f rbac.yaml

# or
$ splatolm ocs-csv.yaml | kubectl -n openshift-storage apply -f
```
