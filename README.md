# Envio de pacotes para o LoRa server


![Image01](https://forum.loraserver.io/uploads/default/optimized/1X/b412bbce85e9dbe872b4458fafcc569dab476712_1_566x500.png "LoRa architecture")

[Link imagem](https://imgur.com/a/lxuxnf0)

#### Requisitos
+ Instalação do LoRa Server, LoRa App Server e LoRa Gateway Bridge (https://www.loraserver.io/use/getting-started/)
+ Configuração através da Interface do Utilizador que se encontra através do link `https://localhost:8080`



#### Como funciona:

Este programa consiste em simular um nó a enviar informações para o servidor LoRa podendo ver o resultado dessas mensagens no terminal ou na interface web.

Foi testado numa máquina virtual que contém este programa, que irá comunicar com a máquina local. Aconselha-se a mudar no programa os links para poderem comunicar para a máquina que tem o loraserver instalado e muita atenção na instalação do loraserver porque se optar por frequências diferentes obriga-o a mudar alguns aspetos como a bandwidth e a frequency do programa.
Caso necessite mudar os ficheiros estão aí presentes, e na linha de comandos dentro da pasta que se encontra deverá apenas executar o `make`

#### Comandos:

Os comandos estão realizados de forma a que possamos enviar diferentes informações de cada dispositivo de sensores:

| Comandos        | output           | 
| ------------- |:-------------:| 
| ./gateway sensor    | Temperatura | 
| ./gateway sensor_jardim      | Humidade      |  
| ./gateway help| ver lista de comandos      | 


###### UBIWHERE
