# sms
Command line tool written in Go that sends an SMS of the given data.

Download:
  * [Mac](https://github.com/codeitloadit/sms/blob/master/bin/sms?raw=true)

Environment Variables:
  * SMS_DEFAULT_MESSAGE 
  * SMS_DEFAULT_NUMBER

Usages:
  * $ sms "This is a message" 5551234567
  * $ sms "This is a message" <- Uses SMS_DEFAULT_NUMBER
  * $ sms <- Uses SMS_DEFAULT_MESSAGE and SMS_DEFAULT_NUMBER
  * $ pwd | sms
  * $ ls | grep test | sms
