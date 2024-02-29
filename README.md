**Descrição do desafio**

    Olá devs!
    Agora é a hora de botar a mão na massa. Para este desafio, você precisará criar o usecase de listagem das orders.
    Esta listagem precisa ser feita com:
    - Endpoint REST (GET /order)
    - Service ListOrders com GRPC
    - Query ListOrders GraphQL
    Não esqueça de criar as migrações necessárias e o arquivo api.http com a request para criar e listar as orders.
    
    Para a criação do banco de dados, utilize o Docker (Dockerfile / docker-compose.yaml), com isso ao rodar o comando docker compose up tudo deverá subir, preparando o banco de dados.
    Inclua um README.md com os passos a serem executados no desafio e a porta em que a aplicação deverá responder em cada serviço.

**Como executar localemente**

Portas utilizadas: BD : 3306
                   WEB_SERVER : 8000
                   GRPC_SERVER: 50051
                   GRAPHQL_SERVER : 8080
                

1 - Após baixar o repositorio execute o <code>docker-compose up -d</code> para criação do banco de dados e do rabbitmq </br>
2 - Neste momento você poderá executar a migration do banco de dados.</br>
 &nbsp; &nbsp; &nbsp;1 - Caso tenha a lib: https://github.com/golang-migrate/migrate basta executar o comando: <code>migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" up</code>.
        Caso contrário poderá executar manualmente o create que encontra-se no arquivo: sql/00001_init.up.sql

3 - Neste momento você pode executar o sistema com o comandado na pasta cmd/ordersystem: <code>go run main.go wire_gen.go</code></br>
4 - A partir de agora você poderá executar tanto a criação de uma order quanto buscar todas as orders armazenadas. Essas duas solicitações poderão ser realizadas tanto em um Web server quanto GRPC e GRAPHQL</br>

    WEB_SERVER: Encontra-se na pasta api exemplo de ambos os casos. 
    
    GRPC_SERVER: Utilizando o client Evans(https://github.com/ktr0731/evans)     
                    1 - abrir o client com o comando: evans -r repl
                    2 - fazer a chamada com o item deseado: 
                          1 - call CreateOrder
                          2 - ListOrders   
    
    GRAPHQL_SERVER: Utilizar o GraphQL playgroud: **localhost:8080** para executar a função desejada. 
                  Ex:  mutation createOrder {
                          createOrder(input: {id:"b2", Price: 10, Tax:10}){
                            id
                            Price
                            Tax
                            FinalPrice
                          }
                        }
      
                       query orders {
                         orders {
                           id
                             Price
                             Tax
                             FinalPrice
                            }
                          } 
