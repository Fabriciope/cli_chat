# CLI CHAT
![Design_sem_nome-removebg](https://github.com/Fabriciope/TechNews/assets/79289410/aac614f4-0761-42b1-9764-14b2c08ca00e)
<br>

 Este projeto é um chat usando o protocolo TCP.


### Bibliotecas utilizadas
 - [Testify](https://github.com/stretchr/testify): pacote usado para fazer os asserts nos testes.
 - [Go terminal size](https://github.com/stretchr/testify): usado para capturar quando o tamanho do terminal é alterado para fazer a adaptação da interface.
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

 Para iniciar o chat do lado do cliente basta executar o comando abaixo trocando o `{n}` por algum número que você queira identificar aquela instância, 
faça isso a cada novo contêiner de cliente que você iniciar, os números não podem se repetir pois não pode haver dois contêineres com o mesmo nome.
 
 Execute o comando abaixo em uma nova sessão do terminal para cado novo usuário do chat que você queira entrar.
```bash
docker run -it --network cli_chat --name cli_chat-client{n} tcp_chat-client:prod
```
![docker-run-client](https://github.com/Fabriciope/TechNews/assets/79289410/56405d26-bf97-45e0-9f7e-31acf299d37a)
<br>

 Execute o comando abaixo em uma nova sessão do terminal pra ver todos os contêineres ativos no momento, tanto o do servidor quanto dos usuários que você iniciou.
```bash
docker container ls
```
![container-ls](https://github.com/Fabriciope/TechNews/assets/79289410/385f921d-6b43-4820-8fdf-87237e046e11)
<br>

 Caso tenha terminado a execução do servidor ou de algum usuário e queira subir o contêiner novamente, execute o comando abaixo substituindo o `{container name}` pelo nome do respectivo contêiner que você quer reiniciar, para visializar os nomes execute o comando enterior novamente.
```bash
docker start -i {container name}
```

<br><br>
**Contato:** fabricioalves.dev@gmail.com
