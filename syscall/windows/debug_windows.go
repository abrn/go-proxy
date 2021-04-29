// +build windows

package windows

import (
    "net"
    "proxy/log"
)

const (
    DNS_QUERY_STANDARD      uint32 = 0x00000000
    DNS_QUERY_BYPASS_CACHE  uint32 = 0x00000008
    DNS_QUERY_NO_HOSTS_FILE uint32 = 0x00000040
    DNS_QUERY_NO_NETBT      uint32 = 0x00000080
    DNS_QUERY_WIRE_ONLY     uint32 = 0x00000100
)

var ()

func PanicHandler() {

}

// ResolveDNSCheck - anti-sandbox technique as they will sometimes have a fake DNS resolver or none at all.
//      this also gives us an easy way to totally block access by IP address and will never affect
//      normal users since they have to resolve the domain to get here first.
func ResolveDNSCheck() bool {
    ip, err := net.LookupCNAME("api.rotmg.network")
    if err != nil {
        log.Logger.Trace("debug_windows: couldn't resolve CNAME record: %s\n", err)
        // we can't resolve the domain so fail at launching
        return false
    }
    log.Logger.Debug("Got CNAME data from DNS %s", ip)
    return true
}
