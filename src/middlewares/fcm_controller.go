package middlewares

import (
	"fmt"

	"github.com/appleboy/go-fcm"
)

func FCM() {
	msg := &fcm.Message{
		To: "fW8tkCbXS-apEvZPB-Ut_T:APA91bHiAmZxfGemGBhg7ogdsqgz57P8HXC1eP5-hOs0rzAOb3jXp-IYPSlQGjZgOFBwmW_GqKrbsLQfF0SKAD7INpas6IQyjczf6gVfhmefP3W8uykMTyKkNc5dbg8Of73oEDBImf3Y",
		// RegistrationIDs: []string{"fXoJADGTQOmSe-gsuyiqv7:APA91bE5M88Den-mHupYI3HwcFM6nmI3U0nONLKE6O-IKPhMn5kCCjsNTFGOS0LhZ3xX84RuYg9yj8qmg6bvacWfeXZ6JNg_ZMDrFUMv1A7qpRZHlUy4zgJr6U3Y3Y-2VQPqcICNYdbL", "fXoJADGTQOmSe-gsuyiqv7:APA91bE5M88Den-mHupYI3HwcFM6nmI3U0nONLKE6O-IKPhMn5kCCjsNTFGOS0LhZ3xX84RuYg9yj8qmg6bvacWfeXZ6JNg_ZMDrFUMv1A7qpRZHlUy4zgJr6U3Y3Y-2VQPqcICNYdbL"},
		Data: map[string]interface{}{
			"foo": "bar",
		},
		Notification: &fcm.Notification{
			Title: "Nadi",
			Body:  "Selamat kamu berhasil bergabung!!",
		},
	}

	client, err := fcm.NewClient("AAAAw2cOOyI:APA91bEv4Z9hU_BttouGHPKtSOg2eEoEf_azjit9CzzKxQ86XiLvxHrSIpqADF2gUSgaqSG6NSw8_QLm_-Ck3Z0C1PnZzWI5buy-LBLQJiBme-qChfsOnoQkMogPsa4TNDvJu_RBJECZ")
	if err != nil {
		fmt.Println("1")
		fmt.Println(err)
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {
		fmt.Println("2")
		fmt.Println(err)
	}

	fmt.Printf("%#v\n", response)

}
