Sistema que você pode criar uma order informando o id, preço e taxa e por fim buscar a listagem de orders criadas acrescidas do valor final (preço + tax). 
Este cadastro e busca poderão ser realizados tanto em um WEB SERVER, quanto em um GRPC SERVER ou GRAPHQL SERVER.


**Como executar localmente**

Portas utilizadas:  BD : 3306<br/>
                    WEB_SERVER : 8000<br/>
                    GRPC_SERVER: 50051<br/>
                     GRAPHQL_SERVER : 8080<br/>
                W
**Docker**<br/>
   Executar o comando: **docker-compose up -d**<br/>
    
**Migrate**<br/>
    Caso tenha a lib: https://github.com/golang-migrate/migrate basta executar o comando:<br/> 
    <code>migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" up</code><br/><br/>
    Caso contrário poderá executar manualmente o create dentro do banco de dados que encontra-se no arquivo: <code>sql/00001_init.up.sql</code><br/>
    
**Execução**<br/>
    Executar o sistema com o comandado na pasta cmd/ordersystem: <code>go run main.go wire_gen.go</code><br/>
    A partir de agora você poderá executar tanto a criação de uma order quanto buscar todas as orders armazenadas.<br/>

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
