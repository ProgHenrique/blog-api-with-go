# Api Blog
esse é um projeto desenvolvido em golang, uma api que faz um CRUD dos posts de um blog.

# Requisitos
1. ter o Go (Golang) instalado na sua máquina. Veja como instalar em <https://go.dev/doc/install>
2. ter o docker instalado na sua máquina. Veja como instalar em <https://docs.docker.com/get-docker/>
3. uma plataforma de api(insomnia, postman e hoppscotch), não é um requisito obrigatório.

# Início
1. clone esse repositório usando o comando
```go
  go get -u https://github.com/ProgHenrique/blog-api-with-go
```
  - espere até que o download esteja completo
2. acesse o repositório usando o comando
```bash
  cd blog-api-with-go
```
3. crie um servidor mysql usando docker da seguinte forma:
  - crie um container
    ```
      docker compose up -d
    ```
  - acesse o container
    ```
      docker compose exec mysql bash
    ```
  - em seguida, no terminal bash que se abriu execute
    ```
      mysql -uroot -p blog
    ```
  - isso exigirá uma senha, que caso nao tenha mudado no arquivo compose.yml, seguirá sendo `root`
  - crie uma nova tabela
    ```sql
      create table posts (id varchar(255), title varchar(255), content varchar(10000), author(80), created_at datetime default current_timestamp, updated_at datetime default current_timestamp on update current_timestamp);
    ```
4. execute a aplicação
```go
  go run main.go
```

## Rotas
1. POST "/posts/create" > Cria um novo post.
  - recebe um JSON no body da requisição. Exemplo
    ```json
      {
        "title": "Sample title", 
        "content": "Lorem Ipsum...", 
        "author": "Name of author"
      }
    ```
2. GET "/posts/get-all" > retorna um array com todos os posts.
3. GET "/posts/{_id}" > retorna o post especificado através do id.
4. DELETE "/posts/{_id}/delete" > deleta o post especificado através do id.
5. PUT "/posts/{_id}/update" > atualiza as informações de um blog especificado pelo id.
  - recebe um JSON no body da requisição. Exemplo
    ```json
      {
        "title": "Update title", 
        "content": "new content of post", 
        "author": "new author of this post"
      }
    ```