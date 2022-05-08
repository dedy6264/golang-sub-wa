package usecase

import (
	"fmt"
	"sub/modelWA"
	"sub/system/examples/sendImage"
	"sub/system/examples/sendTextMessages"
	"time"
)

func Sending(req modelWA.User) modelWA.ResGlobal {
	var result modelWA.ResGlobal
	t := time.Now()
	fmt.Println("Start here: Req===========", req)
	// otpNumb, err := helperWA.GetRandNum()
	// if err != nil {
	// 	result.Status = "31"
	// 	result.StatusDesc = "Error Generate Number" + otpNumb
	// 	return result
	// }
	// _, redisStatus, redisDesc := repository.SetOTP(otpNumb)
	// if redisStatus != "00" {
	// 	result.Status = "31"
	// 	result.StatusDesc = redisDesc
	// 	return result
	// }
	// req.Text = "This is your OTP Number : *" + otpNumb + "*"
	hasil := sendTextMessages.SendText(req)
	result.Status = hasil.Status
	result.StatusDateTime = t
	result.StatusDesc = hasil.StatusDesc
	result.Result = hasil
	fmt.Println("Result : ", result)
	return result
}

func SendingWithImage(req modelWA.ReqSendWAWithImage) modelWA.ResGlobal {
	var result modelWA.ResGlobal
	t := time.Now()
	fmt.Println("Start here: Req===========", req)
	hasil := sendImage.SendWithImage(req)
	result.Status = hasil.Status
	result.StatusDateTime = t
	result.StatusDesc = hasil.StatusDesc
	result.Result = hasil
	return result
}
