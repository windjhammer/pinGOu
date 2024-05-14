# PinGOu

Ferramenta em Go para enviar pings para um host específico N vezes e salvar num txt.

## Requisitos

- Go [baixe aqui](https://golang.org/)

## Uso

1. Clone o repositório ou faça o download do arquivo `main.go`.
2. Navegue até o diretório onde o arquivo `main.go` está localizado.
3. Execute o seguinte comando para compilar o programa:

    ```sh
    go build main.go
    ```

4. Isso criará um executável chamado `main` no mesmo diretório.
5. Execute o `main`:

    ```sh
    ./main
    ```

6. Siga as instruções no terminal para inserir o endereço que deseja pingar e o número de vezes que deseja que o ping seja enviado.

## Funcionalidades

- Funciona com Windows e Linux.
- Salva o output do ping em um arquivo (`ping_log.txt`).
- Trata entradas inválidas do usuário, como endereços vazios ou números de tentativas negativos.

## Contribuição

Contribuições são bem-vindas! Se você encontrar problemas ou tiver sugestões de melhorias, sinta-se à vontade para abrir uma issue ou enviar um pull request.
