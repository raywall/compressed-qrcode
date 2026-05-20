package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	// Definição das flags
	mode := flag.String("mode", "", "Modo de operação: 'encode' ou 'decode'")
	input := flag.String("input", "", "Arquivo de entrada (encode: arquivo .txt; decode: imagem QR Code)")
	output := flag.String("output", "", "Arquivo de saída (encode: imagem .png; decode: arquivo .txt)")
	flag.Parse()

	if *mode == "" || *input == "" || *output == "" {
		fmt.Println("Uso:")
		fmt.Println("  Codificar: go run main.go --mode encode --input texto.txt --output qrcode.png")
		fmt.Println("  Decodificar: go run main.go --mode decode --input qrcode.png --output restaurado.txt")
		os.Exit(1)
	}

	switch *mode {
	case "encode":
		if err := EncodeToQRCode(*input, *output); err != nil {
			fmt.Fprintf(os.Stderr, "Erro na codificação: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("QR Code gerado com sucesso em %s\n", *output)

	case "decode":
		if err := DecodeFromQRCode(*input, *output); err != nil {
			fmt.Fprintf(os.Stderr, "Erro na decodificação: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Texto reconstruído salvo em %s\n", *output)

	default:
		fmt.Fprintf(os.Stderr, "Modo inválido: %s. Use 'encode' ou 'decode'.\n", *mode)
		os.Exit(1)
	}
}
