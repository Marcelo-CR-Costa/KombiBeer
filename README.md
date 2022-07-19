# KombiBeer




* [Índice](#índice)

  1- Descrição
  
  2- Status
  
  3- Funcionalidades
  
  4- Utilização
  
  5- Tecnologias Utilizada
  
  6- Conclusão


![image](https://user-images.githubusercontent.com/103424049/178971486-3fede123-3f90-4f4e-9eeb-7aaa33207ab6.png)

   1- Descrição do Projeto


O desafio foi construir uma API para o projeto KombiBeer, uma empresa que possui um catálogo de cervejas artesanais. Esta API será utilizada para compartilhar dados com seus parceiros e para seu sistema web.

      
O sistema deverá criar, listar, excluir e alterar (parcial ou completamente) as cervejas via API. Deve ser utilizado uma ferramenta para fazer o         versionamento do codigo, deverá rodar em conteiner e possuir um banco de dados e tambem deve conter uma forma de validar o funcionamento. 
Para utilizar o programa deve-se baixar os arquivos do projeto,através do link descrito no acesso ao projeto, no vscode rodar o código main,(go            run.main)



2- [Status do Projeto](#status-do-Projeto)

Concluído parcialmente, foi orientado a fazer a Conteinerização de todo o projeto pelo Docker, (em andamento)

3- Funcionalidades 

Para a utilização deste projeto é necessário ter instalado as seguintes ferramentas:


  - Um editor de texto de sua preferência. (ex: VScode),

  - Postman

  - GoLang 1.18

  - Docker e Docker Compose

  - Git

  
  
4- Utilização:

- Baixar o projeto através do git: 

      git clone https://github.com/Marcelo-CR-Costa/KombiBeer.git


- Abrir o projeto no editor, (VSCode)


no terminal do VsCode, rodar os seguintes comandos: 

      
      Go mod tidy
      
      Docker-compose up (para subir o banco postgres)
      
      go run main.go


- Com a aplicação rodando, você pode utilizar o arquivo json que se encontra dentro da pasta collection_postman e importar para o seu postman. Com isto você poderá:


Buscar produtos por nome ou por ID.


Buscar todos os produtos cadastrados.


Criar produtos.


Editar produtos.


Excluir produtos.

 ![image](https://user-images.githubusercontent.com/103424049/179529413-89883f14-758e-49de-b42e-ae7e83b590e3.png)


5- Técnica Utilizada

 Foi utilizado o editor VScode, a linguagem usada para escrever o codigo foi  Golang, foi utilizado Docker para criar os conteiners, Github, Postman, e postrgree como servidor, tambem foi usado como base para estudos o curso de api da Alura.

O projeto feito em Golang possui: 
* Uma pasta controllers, que contem todas as funções,
* Uma pasta database que contem o codigo responsavel por fazer a conexção com banco de dados, 
* Uma pasta models que contem a struct( estrutura), 
* A pasta routes que contem todos os end points,
* O docker compose que contem o codigo responsavel, por subir um servidor no postgree.
* Os arquivos go mod e go sun, para vincular o projeto ao github.


6- [Conclusão](#conclusão)

O desafio do projeto foi criado pelo Bruno e Wilmar e desnvolvido por Marcelo
Projeto parcialmente finalizado , as funcionalidaes solicitadas  se encontram funcionando perfeitamente, codigos sem falhas ou Bugs, os resultados aqui obtidos foram satisfatorios e auxiliou muito no meu aprendizado.
O desafio foi mais dificil do que eu imaginava mais foi satisfatorio demais, gostaria de agradescer ao apoio e ajudas dos companheiros da Tribo Mesh.
