package ldap

import (
	//"encoding/hex"
	"fmt"
	//"github.com/mavricknz/asn1-ber"
	"testing"
	//"bytes"
)

var addDNs []string = []string{"cn=Jon Boy,ou=People,dc=example,dc=com"}
var addAttrs []EntryAttribute = []EntryAttribute{
	EntryAttribute{
		Name: "objectclass",
		Values: []string{
			"person", "inetOrgPerson", "organizationalPerson", "top",
		},
	},
	EntryAttribute{
		Name: "cn",
		Values: []string{
			"Jon Boy",
		},
	},
	EntryAttribute{
		Name: "givenName",
		Values: []string{
			"Jon",
		},
	},
	EntryAttribute{
		Name: "sn",
		Values: []string{
			"Boy",
		},
	},
}

// Just testing structure and basic dump.
func TestAdd(t *testing.T) {
	fmt.Println("TestAdd starting...\n")
	for _, dn := range addDNs {
		addReq := NewAddRequest(dn)
		for _, attr := range addAttrs {
			addReq.AddAttribute(&attr)
		}
		fmt.Printf("%s", addReq.DumpAddRequest())
	}
	fmt.Println("TestAdd finsished.")
}
