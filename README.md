# k8s-host-device-plugin [![Docker Automated build](https://img.shields.io/docker/automated/everpeace/k8s-host-device-plugin.svg)](https://hub.docker.com/r/everpeace/k8s-host-device-plugin)

This is a very thin device plugin which just exposing device files in host to containers.

Those device files are assumed to be non limited resources, things like "capabilitites" of a host (i.e. /dev/kvm).

However, due to [kubernetes/kubernetes#59380](https://github.com/kubernetes/kubernetes/issues/59380),  current device plugin(`v1beta1`) doesn't support unlimited extended resources.  Currently a hack would be to set the number of resources advertised by the device plugin to a very high mumber.

## How to

see [example](example) folder.

```
cd example

# create configmap with key=`config.json`, value=config json
# this config
#   - resource name: github.com/everpeace-random
#   - host device file: /dev/random
#   - container path to be mounted: /dev/everpeace-mount
#   - permission: "r"
#   - number of devices:  100  (github.com/everpeace-random: 100)
kubectl create -f host-devices.yaml

# deploy device plugin
kubectl create -f host-device-plugin.yaml

# create test pod
# this requests 10 everpeace-random resources. (10 is meaningless.)
kubectl create -f test-pod.yaml
```

## Practical example

If you wanted to expose infiniband devices,  please look into [ib-example](ib-example) directory
