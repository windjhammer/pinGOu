package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func pingHost(host string, count int) error {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("ping", "-n", strconv.Itoa(count), host)
	case "linux":
		cmd = exec.Command("ping", "-c", strconv.Itoa(count), host)
	default:
		return fmt.Errorf("sistema operacional não suportado")
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("erro ao executar ping: %v", err)
	}

	if err := appendToFile("ping_log.txt", string(output)); err != nil {
		return fmt.Errorf("erro ao escrever para o arquivo: %v", err)
	}
	return nil
}

func appendToFile(filename string, data string) error {
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if _, err := file.WriteString(data); err != nil {
		return err
	}
	return nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Insira o endereço que quer pingar: ")
	host, _ := reader.ReadString('\n')
	host = strings.TrimSpace(host)

	if host == "" {
		fmt.Println("Endereço inválido. Por favor, insira um endereço válido.")
		return
	}

	fmt.Print("Insira o número de vezes que deseja que o endereço seja pingado: ")
	countStr, _ := reader.ReadString('\n')
	countStr = strings.TrimSpace(countStr)
	count, err := strconv.Atoi(countStr)
	if err != nil {
		fmt.Println("Número de tentativas inválido. Por favor, insira um número inteiro válido.")
		return
	}

	if count <= 0 {
		fmt.Println("O número de tentativas deve ser maior que zero. Por favor, insira um número positivo.")
		return
	}

	fmt.Printf("ping %s %d vezes, não feche o terminal...\n", host, count)
	if err := pingHost(host, count); err != nil {
		fmt.Println("Erro:", err)
		return
	}

	fmt.Println("Ping concluído. Resultado salvo no arquivo ping_log.txt")
}

