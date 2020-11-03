package packages

import "fmt"

// Client struct
type Client struct{}

// computer interface
type computer interface{
	InsertIntoLightningPort()
}

// InsertLightningConnectorIntoComputer meth
func (c *Client) InsertLightningConnectorIntoComputer(com computer)  {
	fmt.Println("Client inserts Lightning connector into computer.")
	com.InsertIntoLightningPort()
}