package main

import (
	"compress/gzip"
	"fmt"
	"image"
	"io"
	"os"
	"strings"

	"github.com/makiuchi-d/gozxing"
	gozxingqr "github.com/makiuchi-d/gozxing/qrcode"
)

// DecodeFromQRCode lê uma imagem PNG contendo um QR Code, decodifica os bytes
// comprimidos, descomprime e reconstrói o texto original (substituindo ';' por '\n').
func DecodeFromQRCode(qrImagePath, outputTxtPath string) error {
	file, err := os.Open(qrImagePath)
	if err != nil {
		return fmt.Errorf("erro ao abrir imagem QR: %w", err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("erro ao decodificar imagem: %w", err)
	}

	bmp, err := gozxing.NewBinaryBitmapFromImage(img)
	if err != nil {
		return fmt.Errorf("erro ao criar bitmap: %w", err)
	}
	qrReader := gozxingqr.NewQRCodeReader()
	result, err := qrReader.Decode(bmp, nil)
	if err != nil {
		return fmt.Errorf("erro ao ler QR Code: %w", err)
	}
	compressedData := result.GetText()

	gzReader, err := gzip.NewReader(strings.NewReader(compressedData))
	if err != nil {
		return fmt.Errorf("erro ao criar leitor gzip: %w", err)
	}
	defer gzReader.Close()

	decompressed, err := io.ReadAll(gzReader)
	if err != nil {
		return fmt.Errorf("erro na descompressão: %w", err)
	}
	minified := string(decompressed)

	original := strings.ReplaceAll(minified, ";", "\n")

	if err := os.WriteFile(outputTxtPath, []byte(original), 0644); err != nil {
		return fmt.Errorf("erro ao salvar arquivo de saída: %w", err)
	}
	return nil
}
