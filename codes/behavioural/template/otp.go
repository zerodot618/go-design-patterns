package template

type iOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
	publishMetric()
}

type otp struct {
	iOtp iOtp
}

func (o *otp) genAndSendOTP(optLength int) error {
	opt := o.iOtp.genRandomOTP(optLength)
	o.iOtp.saveOTPCache(opt)
	message := o.iOtp.getMessage(opt)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}
	o.iOtp.publishMetric()
	return nil
}
