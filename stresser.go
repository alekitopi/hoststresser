package main

import (
	"fmt"
	"bufio"
	"os"
	"github.com/sparrc/go-ping"
	"strings"
)

const (
	ColorPrefix = "\x1B[35m\x1B[1mStresser »\x07 "
	ColorErrorPrefix = "\x1B[41m\x1B[37m\x1B[1mERROR\x07 "
	Prefix = "Stresser » "
	ErrorPrefix = "ERROR "
)

func main() {
	for {
		Log("Introduce el host que deseas probar.")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		ip := scanner.Text()
		if len(ip) < 7 || strings.Contains(ip, "legacyhcf") {
			Error("El host introducido no es válido.")
		} else {
			Log("Utiliza ENTER para empezar/terminar la prueba.")
			scanner.Scan()
			running := true
			stop := false
			go func() {
				Log("Iniciando la prueba para el host " + ip + ".")
				for running == true {
					fmt.Println("[+] Enviando paquetes al host.")
					err := stresser(ip)
					if err != nil {
						Error("Ha ocurrido un error al intentar ejecutar la prueba.")
						os.Exit(1)
					}
				}
				stop = true
				Log("La prueba ha finalizado correctamente.")
			}()
			scanner.Scan()
			running = false
			for !stop {
				Log("La prueba ha sido finalizada.")
			}
		}
	}
}

func stresser(ip string) error {
	pinger, err := ping.NewPinger(ip)
	if err != nil {
		return err
	}
	pinger.Count = 65500
	pinger.Count = 1
	pinger.Timeout = 1
	pinger.Run()
	return nil
}

func Log(i string) {
	fmt.Println(Prefix + i)
}

func Error(i string) {
	Log(ErrorPrefix + i)
}