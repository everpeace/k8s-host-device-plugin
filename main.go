package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"syscall"

	"github.com/fsnotify/fsnotify"
	pluginapi "k8s.io/kubelet/pkg/apis/deviceplugin/v1beta1"
)

const (
	ConfigFilePath = "/k8s-host-device-plugin/config.json"
)

func main() {
	log.Println("Starging K8s HostDevice Plugin.")

	log.Println("Starting FS watcher.")
	watcher, err := newFSWatcher(pluginapi.DevicePluginPath)
	if err != nil {
		log.Printf("Failed to created FS watcher: %v", err)
		os.Exit(1)
	}
	defer watcher.Close()

	log.Println("Starting OS watcher.")
	sigs := newOSWatcher(syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	log.Println("Reading /k8s-host-device-plugin/config.json")
	raw, err := ioutil.ReadFile(ConfigFilePath)
	if err != nil {
		log.Printf("Fatal error reading config file: %v", err)
		os.Exit(1)
	}

	var config HostDevicePluginConfig
	json.Unmarshal(raw, &config)

	s, _ := json.Marshal(config)
	log.Println("loaded config: ", string(s))

	restart := true
	var devicePlugin *HostDevicePlugin

L:
	for {
		if restart {
			if devicePlugin != nil {
				devicePlugin.Stop()
			}

			devicePlugin, err = NewHostDevicePlugin(config)
			if err != nil {
				log.Printf("Fatal error configuring host device: %v", err)
				os.Exit(1)
			}
			expandedHostDevicesStr := []string{}
			for _, hd := range devicePlugin.hostDevices {
				expandedHostDevicesStr = append(expandedHostDevicesStr, fmt.Sprintf("%+v", hd))
			}
			log.Printf("expanded host devices: %s\n", strings.Join(expandedHostDevicesStr, ","))

			if err := devicePlugin.Serve(); err != nil {
				log.Println("Could not contact Kubelet, retrying. Did you enable the device plugin feature gate?")
			} else {
				restart = false
			}
		}

		select {
		case event := <-watcher.Events:
			if event.Name == pluginapi.KubeletSocket && event.Op&fsnotify.Create == fsnotify.Create {
				log.Printf("inotify: %s created, restarting.", pluginapi.KubeletSocket)
				restart = true
			}

		case err := <-watcher.Errors:
			log.Printf("inotify: %s", err)

		case s := <-sigs:
			switch s {
			case syscall.SIGHUP:
				log.Println("Received SIGHUP, restarting.")
				restart = true
			default:
				log.Printf("Received signal \"%v\", shutting down.", s)
				devicePlugin.Stop()
				break L
			}
		}
	}
}
