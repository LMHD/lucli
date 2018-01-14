// Inspired by Jessie Frazelle

package lib

import (
	"errors"
	"fmt"
	"net"
	"os/exec"
	"runtime"

	log "github.com/Sirupsen/logrus"
)

// GetDisplay returns DISPLAY (for macOS), as would be needed for X11
func GetDisplay() (string, error) {

	display := ""

	if runtime.GOOS == "darwin" {
		log.Debugf("Running in macOS. Need to use host's IP in display")

		ip, err := externalIP()
		if err != nil {
			return "", err
		}

		display = fmt.Sprintf("%s:0", ip)
	} else {
		// TODO: Just use host's $DISPLAY
		// This probably works, but I've not got a Linux env handy to add this at the moment
	}

	log.Debugf("DISPLAY: %s", display)

	return display, nil
}

// https://stackoverflow.com/a/23558495
// externalIP returns the host's external IP address
func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

// StartXQuartz starts XQuartz and give access to it from host IP
func StartXQuartz() error {
	if runtime.GOOS == "darwin" {
		// DO STUFF

		xquartzCommand := exec.Command("open", "-a", "XQuartz")
		log.Debugf("Launching XQuartz")
		err := xquartzCommand.Run()
		if err != nil {
			return err
		}

		ip, err := externalIP()
		if err != nil {
			return err
		}

		xhostCommand := exec.Command("xhost", "+", ip)
		log.Debugf("xhost, to give host IP access to X11")
		err = xhostCommand.Run()
		if err != nil {
			return err
		}
	}
	return nil
}
