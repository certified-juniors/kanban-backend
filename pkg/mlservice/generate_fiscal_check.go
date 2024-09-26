package mlservice

import "fmt"

func (c *clientImpl) GenerateFiscalCheck(uuid string) *string {
	switch c.opts.Endpoint {
	case EndpointTest:
		checkUrl := fmt.Sprintf("%s/%s", CheckUrlTest, uuid)
		return &checkUrl
	case EndpointProd:
		checkUrl := fmt.Sprintf("%s/%s", CheckUrlProd, uuid)
		return &checkUrl
	default:
		return nil
	}
}
