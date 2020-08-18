package util

import (
	"errors"
	"fmt"
	"strconv"
)

func GetServerAddress(port string) (string, error) {
	serverPort, err := strconv.Atoi(port)
	if serverPort < 0 || serverPort > 65535 {
		err = errors.New("port out of range")
	}
	if err != nil {
		return "", fmt.Errorf("server port invalid: %w", err)
	}
	return fmt.Sprintf(":%v", serverPort), nil
}

func PrintBanner() {
	banner := `

  _      _ _   _   _        ____                           
 | |    (_) | | | | |      |  _ \                          
 | |     _| |_| |_| | ___  | |_) |_ __ _____      ___ __   
 | |    | | __| __| |/ _ \ |  _ <| '__/ _ \ \ /\ / / '_ \  
 | |____| | |_| |_| |  __/ | |_) | | | (_) \ V  V /| | | | 
 |______|_|\__|\__|_|\___| |____/|_|  \___/ \_/\_/ |_| |_| 
 |  _ \            | |     / ____| |                       
 | |_) | ___   ___ | | __ | (___ | |__   ___  _ __         
 |  _ < / _ \ / _ \| |/ /  \___ \| '_ \ / _ \| '_ \        
 | |_) | (_) | (_) |   <   ____) | | | | (_) | |_) |       
 |____/ \___/ \___/|_|\_\ |_____/|_| |_|\___/| .__/        
                                             | |           
                                             |_|

`
	fmt.Print(banner)
}
