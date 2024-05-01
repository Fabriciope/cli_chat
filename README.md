<h1 align="center">CLI CHAT :speech_balloon:</h1>

![banner-cli-chat](https://github.com/Fabriciope/TechNews/assets/79289410/9d12fb4f-5247-459e-b65d-d822f209df80)
<br>

[🌐 Documentation translated into English](#documentation-translated-into-english)

 Este projeto é um chat usando o protocolo TCP.


### Libraries used
 - [Testify](https://github.com/stretchr/testify): pacote usado para fazer os asserts nos testes.
 - [Go terminal size](https://github.com/stretchr/testify): usado para identificar quando o tamanho do terminal é alterado para fazer a adaptação da interface.
 - [Strip ANSI](https://github.com/acarl005/stripansi): utilizado para remover códigos de escape ANSI das strings usando regex.
<br>

## Instruções passo a passo para usar o chat localmente
> :warning: Para rodar o servidor tcp e o client será necessáro ter o docker instalado em sua  máquina para subir os contêineres com suas respectivas imagens, portanto se estiver usando windows basta usar o terminal linux com o wsl,
 caso já esteja em um ambiente linux, simplismente execute os comandos em uma nova sessão do terminal.

 Primeiramente faça o clone do projeto, antes verifique se tem o git instalado.
```bash
git clone https://github.com/Fabriciope/cli_chat.git cli-chat && cd cli-chat
```
<br>

 Agora execute os scrips abaixo para a criar a rede utilizada na comunicação entre os clientes e o servidor, e também criar as imagens que serão utilizadas para subir os contêineres do servidor e do cliente.
```bash
./scripts/create_network.sh && \
./scripts/build_images.sh
```
<br>

 Inicialize o servidor tcp executando o camando abaixo no seu terminal. Se você quiser subir o servidor em background para não precisar ver os logs adicione a flag `-d` ao comando para que o contêiner inicie em modo detached.
```bash
docker run -it --network cli_chat --name cli_chat-server tcp_chat-server:prod
```
![Captura de tela de 2024-04-22 23-18-10](https://github.com/Fabriciope/TechNews/assets/79289410/20c85a5e-0994-4676-8f32-660b5187726c)
<br>

 Para iniciar o chat do lado do cliente basta executar o comando abaixo trocando o `<n>` por algum número que você queira identificar aquela instância, 
faça isso a cada novo contêiner de cliente que você iniciar, os números não podem se repetir pois não pode haver dois contêineres com o mesmo nome.
 
 Execute o comando abaixo em uma nova sessão do terminal para cada novo usuário do chat que você queira entrar.
```bash
docker run -it --network cli_chat --name cli_chat-client<n> tcp_chat-client:prod
```
![docker-run-client](https://github.com/Fabriciope/TechNews/assets/79289410/56405d26-bf97-45e0-9f7e-31acf299d37a)
<br>

 Execute o comando abaixo em uma nova sessão do terminal pra ver todos os contêineres ativos no momento, tanto o do servidor quanto dos usuários que você iniciou.
```bash
docker container ls
```
![container-ls](https://github.com/Fabriciope/TechNews/assets/79289410/385f921d-6b43-4820-8fdf-87237e046e11)
<br>

 Caso tenha terminado a execução do servidor ou de algum usuário e queira subir o contêiner novamente, execute o comando abaixo substituindo o `<container name>` pelo nome do respectivo contêiner que você quer reiniciar, para visualizar os nomes execute o comando anterior novamente.
```bash
docker start -i <container name>
```

<br><br>
**Contato:** fabricioalves.dev@gmail.com

<br>

## Documentation translated into english

 This project is a chat using the TCP protocol.
 
 ### Bibliotecas utilizadas
 - [Testify](https://github.com/stretchr/testify): package used to make asserts in tests..
 - [Go terminal size](https://github.com/stretchr/testify): used to identify when the terminal size is changed to adapt the interface.
 - [Strip ANSI](https://github.com/acarl005/stripansi): used to remove ANSI escape codes from strings using regex.
<br>

## Step-by-step instructions for using chat locally
> :warning: To run the tcp server and the client you will need to have docker installed on your machine to upload the containers with their respective images, so if you are using windows just use the linux terminal with wsl,
 If you are already in a Linux environment, simply run the commands in a new terminal session.

 First clone the project, first check if you have git installed.
```bash
git clone https://github.com/Fabriciope/cli_chat.git cli-chat && cd cli-chat
```
<br>

 Now run the scripts below to create the network used in communication between clients and the server, and also create the images that will be used to upload the server and client containers.
```bash
./scripts/create_network.sh && \
./scripts/build_images.sh
```
<br>

 Initialize the tcp server by running the command below in your terminal. If you want to run the server in the background so you don't need to see the logs, add the `-d` flag to the command so that the container starts in detached mode.
```bash
docker run -it --network cli_chat --name cli_chat-server tcp_chat-server:prod
```
![Captura de tela de 2024-04-22 23-18-10](https://github.com/Fabriciope/TechNews/assets/79289410/20c85a5e-0994-4676-8f32-660b5187726c)
<br>

 To start the chat on the client side, simply execute the command below, replacing the `<n>` with any number that you want to identify that instance, 
 do this for every new client container you start, the numbers cannot be repeated as there cannot be two containers with the same name.
 
 Run the command below in a new terminal session for each new chat user you want to join.
```bash
docker run -it --network cli_chat --name cli_chat-client<n> tcp_chat-client:prod
```
![docker-run-client](https://github.com/Fabriciope/TechNews/assets/79289410/56405d26-bf97-45e0-9f7e-31acf299d37a)
<br>

 Run the command below in a new terminal session to see all currently active containers, both the server and the users you started.
```bash
docker container ls
```
![container-ls](https://github.com/Fabriciope/TechNews/assets/79289410/385f921d-6b43-4820-8fdf-87237e046e11)
<br>

 If you have finished running the server or a user and want to upload the container again, run the command below replacing `<container name>` with the name of the respective container you want to restart, to view the names run the previous command again.
```bash
docker start -i <container name>
```

<br><br>
**contact:** fabricioalves.dev@gmail.com
 
