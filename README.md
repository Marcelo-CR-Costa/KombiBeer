# KombiBeer

![image](https://user-images.githubusercontent.com/103424049/178971486-3fede123-3f90-4f4e-9eeb-7aaa33207ab6.png)

* [Índice](#índice)
* 1- Descrição
* 2- Status
* 3- Funcionalidades
* 4- Acesso ao Projeto
* 5- Tecnologias Utilizada
* 6- Pessoas contribuidora
* 7- Conclusão
* [Descrição do Projeto](#descrição-do-Projeto)
*   O desafio foi construir uma API para o projeto KombiBeer, uma empresa que possui um catálogo de cervejas artesanais.Esta API será utilizada para compartilhar dados com seus parceiros e para seu sistema web.
*   * O sistema deverá criar, listar, excluir e alterar (parcial ou completamente) as cervejas via API. Deve ser utilizado uma ferramenta para fazer o versionamento do codigo, deverá rodar em conteiner e possuir um banco de dados e tambem deve conter uma forma de validar o funcionamento
* Para utilizar o programa deve-se baixar os arquivos do projeto,atarves do link descrito no acesso ao projeto, no vscode rodar o codigo main,(go run.main)
* [Status do Projeto](#status-do-Projeto)
* concluido parcialmente, foi orientado a fazer a Conteinerização de todo o projeto pelo Docker, (em andamento)
*   [Funcionalidades e Demonstração da Aplicação](#funcionalidades-e-demonstração-da-aplicação)
*  Para a utilização deste projeto é necessario ter instalado as seguintes ferramentas:
*  um editor de textao (VScode),
*  Postman
*  GoLang
*  Docker
*  Github
*  postgree
*  Utilização:
*  atraves do linK (https://github.com/Marcelo-CR-Costa/KombiBeer), baixar as pastas do projeto utilizando o Gitclone
*  abrir o VScode e carregar a pasta KombiBeer
*  no terminal do VsCode, rodar o codigo, (go run main.go)
*  ele irá iniciar a aplicação irá carregar o postman, tambem vai subir um servidor postgree
*  para carregar o postgree deve abrir uma aba no navegador e utilizar a porta ("54321:80")
*  será solicitado um usuario e senha: 
*  Usuario: "marcelo.costa@sensedia.com"
   Senha:   "123456"
*  a aplicação tembem roda no navegar de internet pelo http://localhost:8080/cervejas
*  Para testar a funcionalidade via postman ![image](https://user-images.githubusercontent.com/103424049/179528269-2b661ffe-68f1-4cf2-9ded-d9773c135af2.png)

*  atráves da coleção postman voce pode:
*  Buscar produtos por nome ou por ID
*  buscar todos os produtos cadastrados
*  criar produtos
*  cadastra produtos
*  editar produtos 
*  excluir produtos
*  ![image](https://user-images.githubusercontent.com/103424049/179529413-89883f14-758e-49de-b42e-ae7e83b590e3.png)
*  O projeto feito em Golang possui: 
* uma pasta controllers, que contem todas as funções,
* A pasta database que contem o codigo responsavel por fazer a conexção com banco de dados, 
* A pasta models que contem a struct( estrutura), 
* A pasta routes que contem todos os end points,
* O docker compose que contem o codigo responsavel, por subir um servidor no postgree
* Os arquivos go mod e go sun, para vincular o projeto ao github
e possivel utilizar o postman para validar as funcionalidades do projeto 
O projeto Cria Produto, Edita, Deleta e Busca, por nome e ID 
* [Acess/o ao Projeto](#acesso-ao-projeto)
* https://github.com/Marcelo-CR-Costa/KombiBeer
* [Tecnologias utilizadas](#tecnologias-utilizadas)
* Foi utilizado o editor VScode, a linguagem usada para escrever o codigo foi  Golang, Docker para criar os conteiners, Github, Postman, e postrgree como servidor, tambem foi usado como base para estudos o curso de api da Alura
* [Pessoas Contribuidoras](#pessoas-contribuidoras)
* WIlmar, Bruno, Wesley,
* [Pessoas Desenvolvedoras do Projeto](#pessoas-desenvolvedoras)
* O desafio do projeto foi criado pelo Bruno e Wilmar e desnvolvido por Marcelo
* [Conclusão](#conclusão)

Projeto parcialmente finalizado , as funcionalidaes solicitadas  se encontram funcionando perfeitamente, codigos sem falhas ou Bugs, os resultados aqui obtidosforam satisfatorios e auxiliou muito no meu aprendizado
o desafio foi mais dificil do que eu imaginava mais foi satisfatorio demais, gostaria de agradescer ao apoio e ajudas dos companheiros da Tribo Mesh
