package main

import (
    "bytes"
    "encoding/json"
    "log"
    "net"
)

type QueryRequest struct {
    Query string `json:"query"`
}

func main() {
    conn, err := net.Dial("tcp", "localhost:5432") //endereço onde está rodando o proxy
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    req := QueryRequest{Query: "SELECT * FROM uma_tabela where coluna = 'algum valor';"}
    jsonQuery, err := json.Marshal(req)
    if err != nil {
        log.Fatal(err)
    }

    // Enviar o payload JSON
    _, err = conn.Write(append(jsonQuery, '\n')) // Adiciona uma nova linha para delimitar
    if err != nil {
        log.Fatal(err)
    }

    // Ler a resposta (opcional)
    var response bytes.Buffer
    _, err = response.ReadFrom(conn)
    if err != nil {
        log.Fatal(err)
    }
    log.Println("Response:", response.String())
}
