package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/skip2/go-qrcode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite o texto para gerar o QR Code: ")
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	fmt.Print("Digite o nome do arquivo para salvar o QR Code (ex: qrcode.png): ")
	fileName, _ := reader.ReadString('\n')
	fileName = strings.TrimSpace(fileName)

	if !strings.HasSuffix(fileName, ".png") {
		fileName += ".png"
	}

	err := qrcode.WriteFile(text, qrcode.Medium, 256, fileName)
	if err != nil {
		fmt.Printf("Erro ao gerar o QR Code: %v\n", err)
		return
	}

	fmt.Printf("QR Code gerado e salvo como %s\n", fileName)
}
