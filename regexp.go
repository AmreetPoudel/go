package regular expression

import (
	"fmt"
	"regexp"
)

func main() {

	// match email
	// re := regexp.MustCompile(`[a-zA-Z0-9_.-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}`)
	// email1 := "amrit@gmail.com"
	// email2 := "a@b.c@"
	// fmt.Println(re.MatchString(email1))
	// fmt.Println(re.MatchString(email2))

	//extract date
	data := "On 2025-06-30 at 09:45 aM, system admin Amrit Poudel logged into the server from IP 192.168.1.50. He used the email address amrit.poudel@example.com to report a failed login attempt for user devadmin at 10:15 PM on 2025-06-29. Another login attempt from IP 10.0.0.254 occurred at 3:40 AM on the same day. The system generated log files like /var/log/auth-2025-06-30.log and /var/log/syslog-2025-06-29.log. Meanwhile, a customer support ticket (#ID-A12345) was raised by bishal123@gmail.com regarding payment failure on invoice INV-2025-6633. Phone number provided was +977-9812345678. A scheduled downtime is planned on 2025/07/01 from 01:00 AM to 04:00 AM. Please contact support@mycompany.co or dial 01-5556789 for urgent issues. Temporary access was granted to user `tempuser1`, and log file temp_access_30_Jun_2025.txt was created."
	// re = regexp.MustCompile(`[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9.-]+`)
	// fmt.Println(re.FindAllString(data, -1))

	// extract all ip addresses
	// re := regexp.MustCompile(`\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}`)
	// fmt.Println(re.FindAllString(data, -1))

	// all dates 2024-06-07
	re := regexp.MustCompile(`\d{4}[/.\-]\d{1,2}[/.\-]\d{1,2}`)
	fmt.Println(re.FindAllString(data, -1))
	// all times 00:00 AM or PM
	re = regexp.MustCompile(`(?i)\d{1,2}:\d{1,2}\s[AP][M]`)

	fmt.Println(re.FindAllString(data, -1))

}
