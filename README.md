# sms
Command line tool written in Go that sends an SMS of the given data.

Environment Variables:
  * SMS_DEFAULT_MESSAGE 
  * SMS_DEFAULT_NUMBER

Usages:
  * $ sms "This is a message" 5551234567
  * pwd | sms
  * ls | grep test | sms
