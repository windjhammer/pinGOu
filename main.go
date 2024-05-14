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

func pingHost(host string, count int) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("ping", "-n", strconv.Itoa(count), host)
	case "linux":
		cmd = exec.Command("ping", "-c", strconv.Itoa(count), host)
	default:
		fmt.Println("SO não suportado")
		return
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	err = appendToFile("ping_log.txt", string(output))
	if err != nil {
		fmt.Println("Erro ao escrever para o arquivo:", err)
		return
	}
	fmt.Println("Resultado salvo no arquivo ping_log.txt")
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

	fmt.Print("Insira o número de vezes que deseja que o endereço seja pingado: ")
	countStr, _ := reader.ReadString('\n')
	countStr = strings.TrimSpace(countStr)
	count, err := strconv.Atoi(countStr)
	if err != nil {
		fmt.Println("Número de tentativas inválido.")
		return
	}

	pingHost(host, count)
}
