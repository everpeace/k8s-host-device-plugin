# k8s-host-device-plugin

This is a very thin device plugin which just exposing device files in host to containers.

Those device files are assumed to be non limited resources, things like "capabilitites" of a host (i.e. /dev/kvm).

However, due to [kubernetes/kubernetes#59380](https://github.com/kubernetes/kubernetes/issues/59380),  current device plugin(`v1beta1`) doesn't support unlimited extended resources.  Currently a hack would be to set the number of resources advertised by the device plugin to a very high mumber.

## Versioning(Tag) Convention
`$KUBEERNETES_VERSION`-`K8S_HOST_DEVICE_PLUGIN_VERSION`

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
#   - number of devices:  100  # please make this value sufficiently large,
#                              # in this case, capacity is `github.com/everpeace-random: 100`
kubectl create -f host-devices.yaml

# deploy device plugin
kubectl create -f host-device-plugin.yaml

# create test pod
# this requests 10 everpeace-random resources. (10 is meaningless.)
kubectl create -f test-pod.yaml
```

## Practical example

If you wanted to expose infiniband devices,  please look into [ib-example](ib-example) directory


## Release Notes
### `1.10.1-0.0.1`
first release

### `1.11.0-0.0.1`
just bump k8s version for `0.0.1`

### `1.11.1-0.0.1`
just bump k8s version for `0.0.1`

### `1.11.1-0.0.2`
support checking device existence [#3](https://github.com/everpeace/k8s-host-device-plugin/pull/3)

### `1.21.3-0.1.0`

- support light-weight glob support [#8](https://github.com/everpeace/k8s-host-device-plugin/pull/8)
- upgrade kubernetes version to 1.21.3 [#7](https://github.com/everpeace/k8s-host-device-plugin/pull/7)
- make it go module [#6](https://github.com/everpeace/k8s-host-device-plugin/pull/6)
- Send only changes of health [#4](https://github.com/everpeace/k8s-host-device-plugin/pull/4)

### `1.22.4-0.1.0`

- upgrade kubernetes version to 1.22.4 [#11](https://github.com/everpeace/k8s-host-device-plugin/pull/11)

### `1.23.5-0.1.0`

- upgrade kubernets 1.23.5 and golang 1.17 [#13](https://github.com/everpeace/k8s-host-device-plugin/pull/13)

### `1.24.4-0.1.0`

- upgrade kubernetes 1.24.4 and golang 1.19 [#16](https://github.com/everpeace/k8s-host-device-plugin/pull/16)

### `1.25.4-0.1.0`

- upgrade kubernetes 1.25.4 [#17](https://github.com/everpeace/k8s-host-device-plugin/pull/17)
- fix binary path in build container in Dockerfile [#18](https://github.com/everpeace/k8s-host-device-plugin/pull/18)

### `1.26.4-0.1.0`

- Upgrade Kubernetes dependencies to v1.26.4 [#22](https://github.com/everpeace/k8s-host-device-plugin/pull/22)

### `1.28.3-0.1.0`

- Upgrade Kubernetes dependencies to 1.28.3 [#23](https://github.com/everpeace/k8s-host-device-plugin/pull/24)

Note: Further release notes will be published in Github [Releases](https://github.com/everpeace/k8s-host-device-plugin/releases) page.
