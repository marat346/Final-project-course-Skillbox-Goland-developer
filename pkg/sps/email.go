package sps

import (
	"bufio"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

type EmailData struct {
	Country      string `json:"country"`
	Provider     string `json:"provider"`
	DeliveryTime int    `json:"delivery_time"`
}

// Индексы данных электронной почты
const (
	COUNTRY_EMAIL = iota
	PROVIDER_EMAIL
	DELIVERY_TIME_EMAIL
)

// GetStatusEmail - получает список данных электронной почты из csv-файла
func GetStatusEmail(csvPath string) ([]EmailData, error) {
	file, err := os.Open(csvPath)
	if err != nil {
		return []EmailData{}, err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return []EmailData{}, err
	}

	err = file.Close()
	if err != nil {
		return []EmailData{}, err
	}

	reader := strings.NewReader(string(content))
	scanner := bufio.NewScanner(reader)
	EmailList := make([]EmailData, 0)

	for scanner.Scan() {
		line := scanner.Text()
		email, ok := parseEmail(line)

		if ok {
			EmailList = append(EmailList, email)
		}
	}

	return EmailList, nil
}

// parseEmail - анализирует строку из csv-файла. Проверяет правильность данных
func parseEmail(line string) (EmailData, bool) {
	email := strings.Split(line, ";")

	switch {
	case len(email) != 3:
		fallthrough
	case !isValidCountry(email[COUNTRY_EMAIL]):
		fallthrough
	case !isValidEmailProvider(email[PROVIDER_EMAIL]):
		fallthrough
	case !isValidDeliveryTime(email[DELIVERY_TIME_EMAIL]):
		return EmailData{}, false
	}

	deliveryTime, _ := strconv.Atoi(email[DELIVERY_TIME_EMAIL])

	return EmailData{
		Country:      email[COUNTRY_EMAIL],
		Provider:     email[PROVIDER_EMAIL],
		DeliveryTime: deliveryTime,
	}, true
}

// GetSlowFastEmailProvider - сортирует всех поставщиков услуг в стране в соответствии со средним временем доставки писем.
// Он возвращает два слайса.
// Первый содержит трех самых быстрых провайдеров, второй - трех самых медленных.
func GetSlowFastEmailProvider(data []EmailData, code string) (slow []EmailData, fast []EmailData) {

	emailsByCountry := make([]EmailData, 0)
	for _, email := range data {
		if email.Country == code {
			emailsByCountry = append(emailsByCountry, email)
		}
	}

	sort.SliceStable(emailsByCountry, func(i, j int) bool {
		return emailsByCountry[i].DeliveryTime < emailsByCountry[j].DeliveryTime
	})

	if len(emailsByCountry) < 3 {
		return emailsByCountry, emailsByCountry
	}

	return emailsByCountry[len(emailsByCountry)-3:], emailsByCountry[:3]
}
