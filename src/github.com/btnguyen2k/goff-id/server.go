package main

import (
	"bytes"
	"github.com/btnguyen2k/olaf"
	"github.com/labstack/echo"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
)

func getMacAddr() (addr string) {
	interfaces, err := net.Interfaces()
	if err == nil {
		for _, i := range interfaces {
			if i.Flags&net.FlagUp != 0 && bytes.Compare(i.HardwareAddr, nil) != 0 {
				// Don't use random as we have a real address
				addr = i.HardwareAddr.String()
				break
			}
		}
	}
	return
}

func getMacAddrAsLong() int64 {
	mac, _ := strconv.ParseInt(strings.Replace(getMacAddr(), ":", "", -1), 16, 64)
	return mac
}

var Olaf = olaf.NewOlaf(getMacAddrAsLong())

func main() {
	readmeIn, _ := ioutil.ReadFile("./README.md")
	readmeOut := blackfriday.Run(readmeIn)

	e := echo.New()
	e.GET("/sf64", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Olaf.Id64())
	})
	e.GET("/sf64hex", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Olaf.Id64Hex())
	})
	e.GET("/sf64ascii", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Olaf.Id64Ascii())
	})
	e.GET("/sf128", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Olaf.Id128())
	})
	e.GET("/sf128hex", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Olaf.Id128Hex())
	})
	e.GET("/sf128ascii", func(c echo.Context) error {
		return c.JSON(http.StatusOK, Olaf.Id128Ascii())
	})
	e.GET("/", func(c echo.Context) error {
		if os.Getenv("DEV_MODE") != "" {
			readmeIn, _ = ioutil.ReadFile("./README.md")
			readmeOut = blackfriday.Run(readmeIn)
		}
		return c.Blob(http.StatusOK, "text/html", readmeOut)
	})
	e.Logger.Fatal(e.Start(":8080"))
}

/*----------------------------------------------------------------------*/
