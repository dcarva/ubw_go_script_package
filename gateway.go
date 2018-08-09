package main

import (
	"os"

	"github.com/spf13/cobra"
	
	"fmt"
        "math/rand"
        "time"

        "github.com/brocaar/lorawan"

        MQTT "github.com/eclipse/paho.mqtt.golang"
        lds "github.com/iegomez/loraserver-device-sim"
)


//import "strconv"

const (
	ipFlag = "ip"
	portFlag="port"

)

func main() {
	cmd := &cobra.Command{
		Use:          "gateway",
		Short:        "Bem-vindo",
		SilenceUsage: true,
	}

	cmd.AddCommand(mainCmd())
	cmd.AddCommand(jardimCmd())

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func mainCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "sensor",
		RunE: func(cmd *cobra.Command, args []string) error {
			ip, err := cmd.Flags().GetString(ipFlag)
			
			port, err2 := cmd.Flags().GetString(portFlag)
			
			if err != nil || err2 != nil {
				return err
				return err2
			}

			/*cmd.Println("A porta:", port)
			
			cmd.Println("O IP da gateway:", ip)*/

			var end_ip string
			end_ip=ip+":"+port
			
			var end string
			end="tcp://"+end_ip
			
			cmd.Println(end)
			


	//Connect to the broker
	opts := MQTT.NewClientOptions()
	opts.AddBroker(end)
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

			for {
				
			

			
	s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	
	s2 := rand.NewSource(time.Now().UnixNano())
    r2 := rand.New(s2)
	
	
	s3 := rand.NewSource(time.Now().UnixNano())
    r3 := rand.New(s3)
	
	var temperatura int
	
	temperatura=r1.Intn(50)	
	temp_total := [1]byte{uint8(temperatura)}
	
	var pressao int
	
	pressao=r2.Intn(100)	
	pressao_total := [1]byte{uint8(pressao)}
	
	
	var humidade int
	
	humidade=r3.Intn(100)	
	humidade_total := [1]byte{uint8(humidade)}
	

	//Create the payload, data rate and rx info.

	payload := []byte{temp_total[0], pressao_total[0], humidade_total[0]}
	fmt.Println("temperatura",payload[0])
	fmt.Println("Pressão",payload[1])
	fmt.Println("Humidade",payload[2])
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
			
			
			
		time.Sleep(3 * time.Second)

	}
			
			
			
			
			
			
			
			return nil
		},
	}
	cmd.Flags().String(ipFlag, "127.0.0.1", "How to connect")
	cmd.Flags().String(portFlag, "1883", "How to connect")
	return cmd
}

func jardimCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "sensorjardim",
		RunE: func(cmd *cobra.Command, args []string) error {
			ip, err := cmd.Flags().GetString(ipFlag)
			
			port, err2 := cmd.Flags().GetString(portFlag)
			
			if err != nil || err2 != nil {
				return err
				return err2
			}

			/*cmd.Println("A porta:", port)
			
			cmd.Println("O IP da gateway:", ip)*/

			var end_ip string
			end_ip=ip+":"+port
			
			var end string
			end="tcp://"+end_ip
			
			cmd.Println(end)
			


	//Connect to the broker
	opts := MQTT.NewClientOptions()
	opts.AddBroker(end)
	opts.SetUsername("your-username")
	opts.SetPassword("your-password")


	client := MQTT.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("Connection error")
		fmt.Println(token.Error())
	}

	fmt.Println("Connection established.")

	//Build your node with known keys (ABP).
	nwsHexKey := "783f27d6a19e8701c4fcda7fdcba759e"
	appHexKey := "a5fe469807d274392020163c0764c647"
	devHexAddr := "018a1f94"
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
	appEUI := [8]byte{0, 0, 0, 0, 0, 0, 0, 3}
	devEUI := [8]byte{0, 0, 0, 0, 0, 0, 0, 3}

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

			for {    // Make a loop of values *-*-*-*-*-*-*-*-*-*-*-*
			
	
		s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
	
	s2 := rand.NewSource(time.Now().UnixNano())
    r2 := rand.New(s2)
	
	
	s3 := rand.NewSource(time.Now().UnixNano())
    r3 := rand.New(s3)
	
	var temperatura int
	
	temperatura=r1.Intn(50)	
	temp_total := [1]byte{uint8(temperatura)}
	
	var pressao int
	
	pressao=r2.Intn(100)	
	pressao_total := [1]byte{uint8(pressao)}
	
	
	var humidade int
	
	humidade=r3.Intn(100)	
	humidade_total := [1]byte{uint8(humidade)}
	

	//Create the payload, data rate and rx info.

	payload := []byte{temp_total[0], pressao_total[0], humidade_total[0]}
	fmt.Println("temperatura",payload[0])
	fmt.Println("Pressão",payload[1])
	fmt.Println("Humidade",payload[2])
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
			
		
		time.Sleep(3 * time.Second)

	}

	
			return nil
		},
	}
	cmd.Flags().String(ipFlag, "127.0.0.1", "How to connect")
	cmd.Flags().String(portFlag, "1883", "How to connect")
	return cmd
}



