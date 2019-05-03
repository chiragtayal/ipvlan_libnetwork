package main

import(
	"log"
	"os"
	"os/signal"
	"syscall"

	"ipvlan_libnetwork/ipvlan"
	"github.com/opencontainers/runc/libcontainer/user"
	"github.com/docker/go-plugins-helpers/network"
	"github.com/docker/go-plugins-helpers/ipam"
)

func main() {

	u, err := user.LookupUser("root")
	if err != nil {
		panic(err)
	}

	nd := ipvlan.NewIpvlanDriver()
	nh := network.NewHandler(nd)
	if err := nh.ServeUnix("Net", u.Gid); err != nil {
		panic(err)
	}

	id := ipvlan.NewIpamDriver()
	ih := ipam.NewHandler(id)
	if err := ih.ServeUnix("Ipam", u.Gid); err != nil {
		panic(err)
	}
	log.Println("... Running ...")

	stopSignal := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(stopSignal, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-stopSignal
		log.Println(sig)
		done <- true
	}()
	<-done
	log.Println("... Stopping ...")
}
