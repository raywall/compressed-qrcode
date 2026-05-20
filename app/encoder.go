package main

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"os"
	"strings"

	qrcode "github.com/skip2/go-qrcode"
)

// EncodeToQRCode lê um arquivo texto, minifica (quebras de linha → ';'),
// comprime com gzip e salva o QR Code como imagem PNG.
func EncodeToQRCode(inputTxtPath, outputQRPath string) error {
	data, err := os.ReadFile(inputTxtPath)
	if err != nil {
		return fmt.Errorf("erro ao ler arquivo de entrada: %w", err)
	}

	// Normaliza quebras de linha
	normalized := strings.ReplaceAll(string(data), "\r\n", "\n")
	normalized = strings.ReplaceAll(normalized, "\r", "\n")
	minified := strings.ReplaceAll(normalized, "\n", ";")

	// Compressão gzip
	var compressed bytes.Buffer
	gzWriter := gzip.NewWriter(&compressed)
	if _, err := gzWriter.Write([]byte(minified)); err != nil {
		gzWriter.Close()
		return fmt.Errorf("erro na escrita da compressão: %w", err)
	}
	if err := gzWriter.Close(); err != nil {
		return fmt.Errorf("erro ao fechar compressor: %w", err)
	}

	// Geração do QR Code (nível de correção Low para maximizar capacidade)
	if err := qrcode.WriteFile(compressed.String(), qrcode.Low, 512, outputQRPath); err != nil {
		return fmt.Errorf("erro ao gerar QR Code: %w", err)
	}
	return nil
}
