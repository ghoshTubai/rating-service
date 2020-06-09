package consul

import (
	"errors"
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"os"
	"rating-service/config"
	"rating-service/leonidas"
	"strconv"
	"strings"
)

func GetServiceProperties() error {
	resource := os.Getenv("ServiceProperties")
	if resource == "" {
		return errors.New(fmt.Sprintf("Value not found for key [%s] from environment variable", "SystemProperties"))
	}
	clientKV, err := retrieveKeysByPrefix(resource)
	if err != nil {
		return err
	}
	for k, v := range clientKV {
		index := strings.LastIndexAny(k, "/")
		key := strings.TrimSpace(k[index+1:])
		os.Setenv(key, v)
	}
	return nil
}

func GetSystemProperties() error {
	resource := os.Getenv("SystemProperties")
	if resource == "" {
		return errors.New(fmt.Sprintf("Value not found for key [%s] from environment variable", "SystemProperties"))
	}
	clientKV, err := retrieveKeysByPrefix(resource)
	if err != nil {
		return err
	}
	for k, v := range clientKV {
		index := strings.LastIndexAny(k, "/")
		key := strings.TrimSpace(k[index+1:])
		os.Setenv(key, v)
	}
	return nil
}

func retrieveKeysByPrefix(resource string) (map[string]string, error) {
	client, err := consulapi.NewClient(consulapi.DefaultConfig())
	if err != nil {
		leonidas.Logging(leonidas.ERROR, nil,
			fmt.Sprintf("Unable to connnect to consul for resource path[%s]. Error[]%s", resource, err.Error()))
		return nil, err
	}
	kv := client.KV()
	keys, _, err := kv.Keys(resource, "\n", nil)
	if err != nil {
		return nil, errors.New("unable to retrieve key value pair from consul for resource path[" + resource + "]. error:" + err.Error())
	}
	maps := make(map[string]string)
	if keys != nil {
		for _, k := range keys {
			pair, _, err := kv.Get(k, nil)
			if err != nil {
				return nil, errors.New("unable to retrieve value for key[" + k + "]. error- " + err.Error())
			}
			maps[pair.Key] = string(pair.Value)
		}
	}
	return maps, nil
}

func RegisterServiceWithConsul() error {
	conf := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(conf)
	if err != nil {
		return err
	}
	registration := new(consulapi.AgentServiceRegistration)

	registration.ID = os.Getenv(config.ServiceName)
	registration.Name = os.Getenv(config.ServiceName)
	address := hostname()
	registration.Address = address
	p, err := strconv.Atoi(os.Getenv(config.HttpPort))
	if err != nil {
		return errors.New("error while converting consul port "+err.Error())
	}
	registration.Port = p
	registration.Check = new(consulapi.AgentServiceCheck)
	registration.Check.HTTP = fmt.Sprintf("http://%s:%v%s", hostname(), p, os.Getenv(config.ServiceHealthEndpoint))
	registration.Check.Interval = os.Getenv(config.ServiceRegistrationCheckInterval)
	registration.Check.Timeout = os.Getenv(config.ServiceRegistrationCheckTimeout)
	if err = consul.Agent().ServiceRegister(registration); err!=nil {
		leonidas.Logging(leonidas.ERROR, nil, "service register is unsuccessful with consul "+ err.Error())
		return err
	}
	leonidas.Logging("INFO", nil, "service registered successful with consul, health path: "+registration.Check.HTTP)
	return nil
}

func LookupServiceWithConsul(serviceName string) (string, error) {
	config := consulapi.DefaultConfig()
	consul, err := consulapi.NewClient(config)
	if err != nil {
		return "", err
	}
	services, err := consul.Agent().Services()
	if err != nil {
		return "", err
	}
	srvc := services["product-service"]
	address := srvc.Address
	port := srvc.Port
	return fmt.Sprintf("http://%s:%v", address, port), nil
}

//func port() string {
//	p := os.Getenv(config.LISTEN_PORT)
//	if len(strings.TrimSpace(p)) == 0 {
//		return ":6666"
//	}
//	return fmt.Sprintf(":%s", p)
//}

func hostname() string {
	hn, err := os.Hostname()
	if err != nil {
		log.Fatalln(err)
	}
	return hn
}
