package internal

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	pkg "github.com/backsoul/viewer/pkg/utils"
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
		return myIP, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return myIP, err
	}

	err = json.Unmarshal(body, &myIP)
	if err != nil {
		return myIP, err
	}

	return myIP, nil
}

func RunBrowser(proxy string, url string) error {
	opts := []chromedp.ExecAllocatorOption{
		chromedp.ProxyServer(proxy),
		chromedp.Flag("headless", true),
		chromedp.WindowSize(1920, 1080),
		chromedp.Flag("disable-sync", true),
		chromedp.Flag("ignore-certificate-errors", false),
		chromedp.Flag("disable-gpu", false), // A veces puede ayudar deshabilitar el GPU en modo headless
		chromedp.UserAgent(pkg.GetRandomUserAgent()),
		chromedp.Flag("disable-extensions", true),
		chromedp.Flag("disable-plugins", true),
		chromedp.Flag("disable-infobars", true),
		chromedp.Flag("disable-translate", true),
		chromedp.Flag("no-sandbox", false),
		chromedp.Flag("disable-dev-shm-usage", true),
		chromedp.Flag("enable-automation", false),
		chromedp.Flag("disable-blink-features", "AutomationControlled"),
		chromedp.Flag("start-maximized", true),
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
		chromedp.Sleep(1000*time.Second),
	)

	if err != nil {
		return nil
	}
	color.Green("finished view: %s,location: %s, proxy: %s", url, ip.Country, proxy)
	return nil
}
