package manager

import (
	"fmt"
	"go-cli-test/constants"
	"go-cli-test/utils"
	"net/http"
)



func RestTest() {

	countries:=utils.RestHelper(http.MethodGet, constants.GetCountriesURL()+"all", nil,nil)


	afghanistan:=utils.RestHelper(http.MethodGet, constants.GetCountriesURL()+"iso3code/" +countries["RestResponse"].(map[string]interface{})["result"].([]interface{})[0].(map[string]interface{})["alpha3_code"].(string), nil,nil)
	fmt.Println(afghanistan["RestResponse"].(map[string]interface{})["result"].(map[string]interface{}))

}

