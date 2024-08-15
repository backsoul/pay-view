package pkg

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func GetProxies(numberProxies int) ([]string, error) {
	var proxies []string

	// Hacer la solicitud HTTP
	res, err := http.Get("https://www.sslproxies.org/")
	if err != nil {
		return nil, fmt.Errorf("error al hacer la solicitud: %w", err)
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Cargar el documento HTML
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error al cargar el documento HTML: %w", err)
	}

	// Extraer los proxies de la tabla, limitando a 10
	doc.Find(".table tbody tr").EachWithBreak(func(index int, item *goquery.Selection) bool {
		if len(proxies) >= numberProxies {
			return false // Detener la iteración después de 10 proxies
		}

		ip := item.Find("td:nth-child(1)").Text()
		port := item.Find("td:nth-child(2)").Text()
		if ip != "" && port != "" {
			proxies = append(proxies, fmt.Sprintf("http://%s:%s", ip, port))
		}
		return true
	})

	return proxies, nil
}
