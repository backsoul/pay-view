package internal

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/fatih/color"
	"github.com/go-rod/stealth"
)

type MyIPResponse struct {
	IP      string `json:"ip"`
	Country string `json:"country"`
	CC      string `json:"cc"`
}

func GetInformationIP(proxy string) (ip MyIPResponse, err error) {
	urlMyIP := "https://api.myip.com"
	var myIP MyIPResponse
	proxyURL, err := url.Parse(proxy)
	if err != nil {
		fmt.Println("Error parsing proxy URL:", err)
		return myIP, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,             // Permitir conexiones inseguras (no recomendado para producción)
			MinVersion:         tls.VersionTLS10, // Forzar al menos TLS 1.2
			CipherSuites: []uint16{
				tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
				tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
				tls.TLS_RSA_WITH_AES_128_GCM_SHA256,
				tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			},
		},
	}

	client := &http.Client{
		Transport: transport,
	}

	resp, err := client.Get(urlMyIP)
	if err != nil {
		fmt.Println("Error making request:", err)
		return myIP, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return myIP, err
	}

	err = json.Unmarshal(body, &myIP)
	if err != nil {
		fmt.Printf("Error unmarshalling response: %v\nResponse Body: %s\n", err, string(body))
		return myIP, err
	}

	return myIP, nil
}

func RunBrowser(proxy string, url string) error {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.ProxyServer(proxy),
		chromedp.Flag("headless", false),
		chromedp.WindowSize(1920, 1080),
		chromedp.Flag("disable-sync", true),
		chromedp.Flag("ignore-certificate-errors", false),
		chromedp.Flag("disable-gpu", false), // A veces puede ayudar deshabilitar el GPU en modo headless
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36"),
	}

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	ip, err := GetInformationIP(proxy)
	if err != nil {
		return nil
	}

	color.Green("viewed: %s,location: %s, proxy: %s", url, ip.Country, proxy)
	err = chromedp.Run(ctx,
		chromedp.Evaluate(stealth.JS, nil),
		chromedp.Navigate(url),
		chromedp.Sleep(10000*time.Second),
	)

	if err != nil {
		color.Red("proxy failed entry platform: %s", proxy)
		return nil
	}
	color.Green("finished view: %s,location: %s, proxy: %s", url, ip.Country, proxy)
	return nil
}

func RunBrowserOndetah(url string) ([]string, error) {
	var statuses []string

	opts := []chromedp.ExecAllocatorOption{
		chromedp.Flag("headless", true),
		chromedp.WindowSize(1920, 1080),
		chromedp.Flag("disable-sync", true),
		chromedp.Flag("ignore-certificate-errors", false),
		chromedp.Flag("disable-gpu", false),
		chromedp.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.149 Safari/537.36"),
	}

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx)
	defer cancel()

	var htmlContent string
	err := chromedp.Run(ctx,
		chromedp.Evaluate(stealth.JS, nil),
		chromedp.Navigate(url),
		chromedp.WaitVisible("body", chromedp.ByQuery),
		chromedp.Sleep(2*time.Second),
		chromedp.Evaluate(`document.body.innerHTML`, &htmlContent),
	)

	if err != nil {
		return statuses, fmt.Errorf("failed to run chromedp: %w", err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		return statuses, err
	}

	doc.Find(".m-timeline-2__item").Each(func(i int, s *goquery.Selection) {
		iconClass, _ := s.Find(".m-timeline-2__item-cricle i").Attr("class")

		var prefix string
		switch {
		case strings.Contains(iconClass, "fa-clock") && strings.Contains(iconClass, "m--font-warning"):
			prefix = "created -"
		case strings.Contains(iconClass, "far fa-check-circle") && strings.Contains(iconClass, "m--font-success"):
			prefix = "finished -"
		default:
			prefix = "process -"
		}

		text := strings.TrimSpace(s.Find(".m-timeline-2__item-text--bold").Text())
		statuses = append(statuses, text)
		fmt.Println(prefix + text)
	})

	return statuses, nil
}
