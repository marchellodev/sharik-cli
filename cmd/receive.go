package cmd

import (
	"fmt"
	"github.com/marchellodev/sharic/lib"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(receiveCmd)
}

var receiveCmd = &cobra.Command{
	Use:   "receive",
	Short: "Looks for shariks in your local network",
	Long:  `Looks for other devices sharik files using sharik in your local network`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("getting your ip...")
		ip, _ := lib.GetLocalIp()
		//
		//if ping {
		//	fmt.Println("ip was fetched the ping way")
		//} else {
		//	fmt.Println("ip was fetched the dumb, probably not-working way, please file an issue")
		//}

		ips := strings.Split(ip.String(), ".")
		fmt.Println("the search mask is:", ips[0]+"."+ips[1]+"."+ips[2]+".*")

		// todo as a single function
		lib.RunDiscoveryDaemon(2*time.Second, func(peer lib.Peer, status int) {
			// todo interactive list
			if status == lib.PeerAdd {
				fmt.Println("Discovered sharik: http://" + peer.String())
			}
			if status == lib.PeerRemove {
				fmt.Println("Sharik was closed: http://" + peer.String())
			}
		})

	},
}
