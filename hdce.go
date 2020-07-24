package main

import (
  "fmt"
  "encoding/base64"
)

func main() {
  var usuarioHarbor string
  var senhaHarbor string

  fmt.Println("Encoding docker config by username and password")
  fmt.Println("                     by EricoGR 20200723 v0.0.1\n")

  fmt.Printf("Digite o usuário harbor: ")
  fmt.Scanln(&usuarioHarbor)

  fmt.Printf("Digite a senha harbor: ")
  fmt.Scanln(&senhaHarbor)

  fmt.Println("\n" + encodingDockerConfig(usuarioHarbor, senhaHarbor))
}

func encodingDockerConfig(usuarioHarbor string, senhaHarbor string) (encodedDockerConfig string) {
  return encoderBase64(templateDockerConfig(encoderBase64(concatenaUsuarioSenha(usuarioHarbor, senhaHarbor))))
}

func templateDockerConfig(auth string) (dockerConfig string){
  dockerConfig = "{\"auths\":{\"https://docker.totvs.io/v1/\":{\"auth\":\"" + auth + "\"}}}"
  return
}

func concatenaUsuarioSenha(usuario string, senha string) (concatenado string) {
  concatenado = usuario + ":" + senha
  return
}

func encoderBase64(texto string) (codificado string) {
  data := []byte(texto)
  codificado = base64.StdEncoding.EncodeToString(data)
  return
}
