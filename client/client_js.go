// +build js

package client

import (
	"fmt"

	"github.com/the729/go-libra/config"
	"github.com/the729/go-libra/generated/pbac"
	"github.com/the729/go-libra/types/validator"
)

func (c *Client) connect(server string) error {
	c.ac = pbac.NewAdmissionControlClient(server)
	return nil
}

func (c *Client) loadTrustedPeers(tomlData string) error {
	peerconf, err := config.LoadTrustedPeers(tomlData)
	if err != nil {
		return fmt.Errorf("load conf err: %v", err)
	}
	verifier, err := validator.NewConsensusVerifier(peerconf)
	if err != nil {
		return fmt.Errorf("new verifier err: %v", err)
	}
	c.verifier = verifier
	return nil
}
