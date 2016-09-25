# netscaler-nitro-go

Lightweight go library for netscaler nitro API.

examples:
```go
	// Create a nitro client
	nClient := client.GetNitroClient("http", "[IP ADDRESS]", "config",
		"[ID]", "[PASSWORD]", true)
		
	// Enable SSL
    nsfeatureReq := datatypes.NsfeatureReq{
        Nsfeature: &datatypes.Nsfeature{
       			Feature: []string{"SSL"},
   		},
   	}
	err := nClient.Enable(&nsfeatureReq, true)
	if err != nil {
		fmt.Printf("[ERROR]" + err.Error())
	}

	// Create a virtual server
	lbvserverReq := datatypes.LbvserverReq{
		Lbvserver: &datatypes.Lbvserver{
			Name:        op.String("[VIP Name]"),
			Ipv46:       op.String("[VIP IP]"),
			Port:        op.Int(80),
			ServiceType: op.String("HTTP"),
			Lbmethod:    op.String("ROUNDROBIN"),
		},
	}

	err = nClient.Add(&lbvserverReq)
	if err != nil {
		fmt.Printf("[ERROR]" + err.Error())
	}
```