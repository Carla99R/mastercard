package signature

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	oauth "github.com/mastercard/oauth1-signer-go"
	oauthUtils "github.com/mastercard/oauth1-signer-go/utils"
	"log"
	openapi "masterCard/out"
	"masterCard/utils"
	"net/http"
)

func Sign(ctx iris.Context) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Print(utils.ReadXml("01"))
		return
	}

	body, err := ctx.GetBody()
	var result map[string]interface{}
	err = json.Unmarshal(body, &result)

	e, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(e))

	signingKey, err := oauthUtils.LoadSigningKey(
		"utils/files/pruebaaa-sandbox.p12",
		"Carlyta99!!")

	consumerKey := "Wx6SZLN4oIoYSan3J5NDGieJmg4IXIaF9_CjohPa801cfe56!ff972d42dce445a299f7fbe5a05ba1d50000000000000000"
	//url, _ := url.Parse("https://sandbox.api.mastercard.com/send/v1/partners/ptnr_BEeCrYJHh2BXTXPy_PEtp-8DBOo/transfers/payment")
	//method := "POST"
	payload := e
	//authHeader, err := oauth.GetAuthorizationHeader(url, method, []byte(payload), consumerKey, signingKey)
	//fmt.Println("auth")
	//log.Print(authHeader)
	//fmt.Println(authHeader)
	request, _ := http.NewRequest("POST", "https://sandbox.api.mastercard.com/send/v1/partners/ptnr_BEeCrYJHh2BXTXPy_PEtp-8DBOo/transfers/payment", bytes.NewBuffer(payload))
	signer := &oauth.Signer{
		ConsumerKey: consumerKey,
		SigningKey:  signingKey,
	}
	config := openapi.NewConfiguration()
	apiClient := openapi.NewAPIClient(config)

	err = signer.Sign(request)
	req := apiClient.PaymentTransferApi.CreatePayment(context.Background(), "ptnr_BEeCrYJHh2BXTXPy_PEtp-8DBOo")
	wrapper, resp, err := req.ApiService.CreatePaymentExecute(req)
	fmt.Println(wrapper)
	fmt.Println(resp)

	fmt.Println(request.Header["Authorization"])
	fmt.Println(len(payload))
	//request.Header.Set("Content-Type", "application/json")
	//request.Header.Set("Content-Length", strconv.Itoa(len(payload)))
	//client := &http.Client{Timeout: 30 * time.Second}
	//resp, err := client.Do(request)
	//defer resp.Body.Close()
	//bodyBytes, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println("RESPUESTAAA")
	//fmt.Println(string(bodyBytes))
}
