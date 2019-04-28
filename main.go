package main

import(
	"log"
	"os"
	"os/signal"
	"syscall"

	"ipvlan_libnetwork/ipvlan"
	"github.com/opencontainers/runc/libcontainer/user"
	"github.com/docker/go-plugins-helpers/network"
)

func main() {

	d := ipvlan.NewDriver()
	u, err := user.LookupUser("root")
	if err != nil {
		panic(err)
	}

	h := network.NewHandler(d)
	if err2 := h.ServeUnix("Net", u.Gid); err2 != nil {
		panic(err2)
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
