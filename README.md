# Envio de pacotes para o LoRa server


![Image01](https://forum.loraserver.io/uploads/default/optimized/1X/b412bbce85e9dbe872b4458fafcc569dab476712_1_566x500.png "LoRa architecture")

[Link imagem](https://imgur.com/a/lxuxnf0)

#### Requisitos
+ Instalação do LoRa Server, LoRa App Server e LoRa Gateway Bridge (https://www.loraserver.io/use/getting-started/)
+ Configuração através da Interface do Utilizador que se encontra através do link `https://localhost:8080`
+ go get github.com/spf13/cobra
+ go get github.com/iegomez/loraserver-device-sim
+ go get github.com/eclipse/paho.mqtt.golang
+ go get github.com/brocaar/lorawan



#### Como funciona:

Este programa consiste em simular um nó a enviar informações para o servidor LoRa podendo ver o resultado dessas mensagens no terminal ou na interface web.

Foi testado numa máquina virtual que contém este programa, que irá comunicar com a máquina local. Aconselha-se a mudar no programa os links para poderem comunicar para a máquina que tem o loraserver instalado e muita atenção na instalação do loraserver porque se optar por frequências diferentes obriga-o a mudar alguns aspetos como a bandwidth e a frequency do programa.
Caso necessite mudar os ficheiros estão aí presentes, e na linha de comandos dentro da pasta que se encontra deverá apenas executar o `make`

#### Comandos:

Os comandos estão realizados de forma a que possamos enviar diferentes informações de cada dispositivo de sensores:

| Comandos        | Descrição |
| ------------- |-------------|
| ./gateway -h    | ajuda |
| ./gateway nome_sensor --port=1883 | Mudar a porta do MQTT Broker caso seja diferente      | 
| ./gateway nome_sensor --ip=localhost |Mudar para o IP que está instalado o seu LoRa server instalado | 



#### Se pretender usar o node-red poderá utilizar o flow que está criado (https://gist.github.com/dcarva/2d320d2a3a649ec238b4a0ccbe510b78) ou (https://flows.nodered.org/flow/2d320d2a3a649ec238b4a0ccbe510b78) deverá instalar algumas bibliotecas como base64 e dashboard

###### UBIWHERE
