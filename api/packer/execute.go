package packer

import (
	"encoding/json"
	"log"

	"github.com/SevereCloud/vksdk/v2/api"
)

type packedExecuteResponse struct {
	Responses     map[string]json.RawMessage
	ExecuteErrors api.ExecuteErrors
}

func (p *Packer) execute(code string) (packedExecuteResponse, error) {
	resp, err := p.vkHandler("execute", api.Params{
		"access_token": p.tokenPool.Get(),
		"v":            api.Version,
		"code":         code,
	})
	if err != nil {
		return packedExecuteResponse{}, err
	}

	if p.debug {
		log.Printf("packer: execute: response: \n%s\n", resp.Response)
	}

	execResponses := make(map[string]json.RawMessage)
	if err := json.Unmarshal(resp.Response, &execResponses); err != nil {
		return packedExecuteResponse{}, err
	}

	return packedExecuteResponse{
		execResponses,
		resp.ExecuteErrors,
	}, nil
}
