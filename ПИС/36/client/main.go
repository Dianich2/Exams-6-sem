package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type WebDAVClient struct {
	URL      string
	Username string
	Password string
	Client   *http.Client
}

func CreateClient(
	url string,
	username string,
	password string,
) *WebDAVClient {
	return &WebDAVClient{
		URL:      url,
		Username: username,
		Password: password,
		Client:   &http.Client{},
	}
}

// MKCOL
func (c *WebDAVClient) MKCOL(
	path string,
) error {
	fullUrl := c.URL + path

	req, err := http.NewRequest(
		"MKCOL",
		fullUrl,
		nil,
	)

	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Password)

	response, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	fmt.Println("MKCOL: ", response.Status)

	if response.StatusCode != 201 && response.StatusCode != 409 {
		return fmt.Errorf("MKCOL failed: %s", response.Status)
	}

	return nil
}

// PUT
func (c *WebDAVClient) PUT(
	fileName string,
	path string,
) error {
	fullUrl := c.URL + path

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	defer file.Close()

	req, err := http.NewRequest(
		"PUT",
		fullUrl,
		file,
	)

	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Password)

	response, err := c.Client.Do(req)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	fmt.Println("PUT: ", response.Status)

	if response.StatusCode != 201 && response.StatusCode != 200 {
		return fmt.Errorf("PUT failed: %s", response.Status)
	}

	return nil
}

func (c *WebDAVClient) GET(
	fileName string,
	path string,
) error {
	fullUrl := c.URL + path

	req, err := http.NewRequest(
		"GET",
		fullUrl,
		nil,
	)

	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Password)

	response, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	getFile, err := os.Create(fileName)
	if err != nil {
		return err
	}

	defer getFile.Close()

	_, err = io.Copy(getFile, response.Body)
	if err != nil {
		return err
	}

	fmt.Println("GET: ", response.Status)

	return nil
}

func (c *WebDAVClient) DELETE(
	path string,
) error {
	fullUrl := c.URL + path

	req, err := http.NewRequest(
		"DELETE",
		fullUrl,
		nil,
	)

	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Password)

	response, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	fmt.Println("DELETE: ", response.Status)

	return nil
}

func (c *WebDAVClient) COPY(
	srcPath string,
	destPath string,
) error {
	srcFullPath := c.URL + srcPath
	destFullPath := c.URL + destPath

	req, err := http.NewRequest(
		"COPY",
		srcFullPath,
		nil,
	)

	if err != nil {
		return err
	}

	req.Header.Set("Destination", destFullPath)
	req.SetBasicAuth(c.Username, c.Password)

	response, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	fmt.Println("COPY: ", response.Status)

	return nil
}

func (c *WebDAVClient) MOVE(
	srcPath string,
	destPath string,
) error {
	srcFullPath := c.URL + srcPath
	destFullPath := c.URL + destPath

	req, err := http.NewRequest(
		"MOVE",
		srcFullPath,
		nil,
	)

	if err != nil {
		return err
	}

	req.Header.Set("Destination", destFullPath)
	req.SetBasicAuth(c.Username, c.Password)

	response, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	fmt.Println("MOVE: ", response.Status)

	return nil
}

func (c *WebDAVClient) RMCOL(
	path string,
) error {
	fullUrl := c.URL + path

	req, err := http.NewRequest(
		"DELETE",
		fullUrl,
		nil,
	)

	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Password)

	response, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer response.Body.Close()

	fmt.Println("RMCOL: ", response.Status)

	return nil
}

func main() {
	// client := CreateClient(
	// 	"http://localhost:8080",
	// 	"myuser",
	// 	"2",
	// )

	// client.MKCOL("/test/")

	// os.WriteFile("test.txt", []byte("hello"), 0644)

	// client.PUT("test.txt", "/test/test.txt")

	// client.GET("downloaded.txt", "/test/test.txt")

	// client.COPY("/test/test.txt", "/test/copy.txt")

	// client.MOVE("/test/copy.txt", "/test/moved.txt")

	// client.DELETE("/test/moved.txt")
	// client.DELETE("/test/test.txt")
	// client.RMCOL("/test/")

	client := CreateClient(
		"http://localhost:8090",
		"myuser",
		"2",
	)

	client.MKCOL("/test/")

	os.WriteFile("test.txt", []byte("hello"), 0644)

	client.PUT("test.txt", "/test/test.txt")

	client.GET("downloaded.txt", "/test/test.txt")

	client.COPY("/test/test.txt", "/test/copy.txt")

	client.MOVE("/test/copy.txt", "/test/moved.txt")

	client.DELETE("/test/moved.txt")
	client.DELETE("/test/test.txt")
	client.RMCOL("/test/")
}
