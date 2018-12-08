package main

import (
	"fmt"
	"github.com/devans10/go-purestorage/flasharray"
	"os"
)

func main() {

	//enter username and password OR api_token
	//enter array for target
	username := ""
	password := ""
	api_token := ""
	rest_version := ""
	target := "<array>"
	verify_https := false
	ssl_cert := false
	user_agent := ""
	suffix := "test2"
	c, err := flasharray.NewClient(target, username, password, api_token, rest_version, verify_https, ssl_cert, user_agent, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	v, verr := c.Volumes.CreateVolume("vol-"+suffix, "1G", nil)
	if verr != nil {
		fmt.Println(verr)
		os.Exit(1)
	}
	fmt.Println("Volume Create")
	fmt.Println("-------------")
	fmt.Println("Name: ", v.Name)
	fmt.Println("Serial: ", v.Serial)
	fmt.Println("Size: ", v.Size)

	h, herr := c.Hosts.CreateHost("host-"+suffix, nil)
	if err != nil {
		fmt.Println(herr)
		os.Exit(1)
	}
	fmt.Println("Host Create")
	fmt.Println("-------------")
	fmt.Println("Name: ", h.Name)

	hg, hgerr := c.Hostgroups.CreateHostgroup("hgroup-"+suffix, nil)
	if err != nil {
		fmt.Println(hgerr)
		os.Exit(1)
	}
	fmt.Println("Hostgroup Create")
	fmt.Println("-------------")
	fmt.Println("Name: ", hg.Name)

	pg, pgerr := c.Protectiongroups.CreateProtectiongroup("pgroup-"+suffix, nil)
	if pgerr != nil {
		fmt.Println(pgerr)
		os.Exit(1)
	}
	fmt.Println("Protectiongroup Create")
	fmt.Println("-------------")
	fmt.Println("Name: ", pg.Name)

	_, pg2err := c.Protectiongroups.DestroyProtectiongroup("pgroup-"+suffix, nil)
	if pg2err != nil {
		fmt.Println(pg2err)
		os.Exit(1)
	}
	_, hg2err := c.Hostgroups.DeleteHostgroup("hgroup-"+suffix, nil)
	if hg2err != nil {
		fmt.Println(hg2err)
		os.Exit(1)
	}
	_, h2err := c.Hosts.DeleteHost("host-"+suffix, nil)
	if h2err != nil {
		fmt.Println(h2err)
		os.Exit(1)
	}
	_, v2err := c.Volumes.DeleteVolume("vol-"+suffix, nil)
	if v2err != nil {
		fmt.Println(v2err)
		os.Exit(1)
	}

}
