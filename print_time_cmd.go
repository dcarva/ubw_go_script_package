package main

import (
	"github.com/spf13/cobra"

	"fmt"
        "math/rand"
        "time"

        "github.com/brocaar/lorawan"

        MQTT "github.com/eclipse/paho.mqtt.golang"
        lds "github.com/iegomez/loraserver-device-sim"

)

func printTimeCmd() *cobra.Command {
	return &cobra.Command{
		Use: "tempoagora",
		RunE: func(cmd *cobra.Command, args []string) error {
			now := time.Now()
			prettyTime := now.Format(time.RubyDate)
			cmd.Println("Bem-vindo(a)! O tempo atual da máquina é", prettyTime)
			return nil
		},
	}
}

func printHelloCmd() *cobra.Command {
	return &cobra.Command{
	   Use: "ola",
                RunE: func(cmd *cobra.Command, args []string) error {
                        frase := "Olá espero que goste"
                        cmd.Println("Bem-vindo à máquina ",frase)
                        return nil
                },


	}


}

func mainCmd() *cobra.Command {

	return &cobra.Command{
	Use: "sensor",
		RunE: func(cmd *cobra.Command, args []string) error {




	//Connect to the broker
	opts := MQTT.NewClientOptions()
	opts.AddBroker("tcp://192.168.2.87:1883")
	opts.SetUsername("your-username")
	opts.SetPassword("your-password")

	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Connection error")
		fmt.Println(token.Error())
	}

	fmt.Println("Connection established.")

	//Build your node with known keys (ABP).
	nwsHexKey := "3bc0ddd455d320a6f36ff6f2a25057d0"
	appHexKey := "00de01b45b59a4df9cc2b3fa5eb0fe7c"
	devHexAddr := "07262b83"
	devAddr, err := lds.HexToDevAddress(devHexAddr)
	if err != nil {
		fmt.Printf("dev addr error: %s", err)
	}

	nwkSKey, err := lds.HexToKey(nwsHexKey)
	if err != nil {
		fmt.Printf("nwkskey error: %s", err)
	}

	appSKey, err := lds.HexToKey(appHexKey)
	if err != nil {
		fmt.Printf("appskey error: %s", err)
	}

	appKey := [16]byte{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
	appEUI := [8]byte{0, 0, 0, 0, 0, 0, 0, 2}
	devEUI := [8]byte{0, 0, 0, 0, 0, 0, 0, 2}

	device := &lds.Device{
		DevEUI:  devEUI,
		DevAddr: devAddr,
		NwkSKey: nwkSKey,
		AppSKey: appSKey,
		AppKey:  appKey,
		AppEUI:  appEUI,
		UlFcnt:  0,
		DlFcnt:  0,
	}

	/*
		*	Make up some random values.
		*
		*	These should be decoded at lora-app-server with a proper function.
		* 	For this example, the object should look like this:
			obj : {
				"temperature": {
					"value":((bytes[0]*256+bytes[1])/100),"unit":"°C"
				},
				"pressure": {
					"value":((bytes[2]*16*16*16*16+bytes[3]*256+bytes[4])/100),"unit":"hPa"
				},
				"humidity": {
					"value":((bytes[5]*256+bytes[6])/1024),"unit":"%"
				}
			}
		*
	*/
	rand.Seed(time.Now().UnixNano() / 10000)
	temp := [2]byte{uint8(rand.Intn(25)), uint8(rand.Intn(100))}
	pressure := [3]byte{uint8(rand.Intn(2)), uint8(rand.Intn(20)), uint8(rand.Intn(100))}
	humidity := [2]byte{uint8(rand.Intn(100)), uint8(rand.Intn(100))}

	//Create the payload, data rate and rx info.
	payload := []byte{temp[0], temp[1], pressure[0], pressure[1], pressure[2], humidity[0], humidity[1]}

	//Change to your gateway MAC to build RxInfo.
	gwMac := "b827ebfffeb13d1f"

	//Construct DataRate RxInfo with proper values according to your band (example is for US 915).

        dataRate := &lds.DataRate{
        //      Bandwidth:    500,
                Bandwidth: 125,
                Modulation:   "LORA",
                SpreadFactor: 8,
                BitRate:      0}

        rxInfo := &lds.RxInfo{
                Channel:   0,
                CodeRate:  "4/5",
                CrcStatus: 1,
                DataRate:  dataRate,
		Frequency: 869000000,
                LoRaSNR:   7,
                Mac:       gwMac,
                RfChain:   1,
                Rssi:      -57,
                Size:      23,
                Time:      time.Now().Format(time.RFC3339),
                Timestamp: int32(time.Now().UnixNano() / 1000000000),
        }

	    //Now send an uplink
        err = device.Uplink(client, lorawan.UnconfirmedDataUp, 1, rxInfo, payload)
        if err != nil {
                fmt.Printf("couldn't send uplink: %s\n", err)
        }
return nil
},
}


}

