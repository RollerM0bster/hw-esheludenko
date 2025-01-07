package hw10programoptimization

import (
	"bufio"
	"io"
	"regexp"
	"strings"

	"github.com/goccy/go-json"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	var domainRegex = regexp.MustCompile(`@.*\.` + domain)
	result := make(DomainStat)
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}
		var user User
		if err := json.Unmarshal(line, &user); err != nil {
			return nil, err
		}
		if domainRegex.MatchString(user.Email) {
			index := strings.LastIndex(user.Email, "@")
			if index != -1 {
				domainPart := strings.ToLower(user.Email[index+1:])
				result[domainPart]++
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return result, nil
}
