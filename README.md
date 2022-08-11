# LOGGING 
HOOK

Hook : struct berupa callback will run if terjd kejadian untuk level log tertentu (semacam event jika di event-driven kafka) . sejenis callback/listener jika di js

ex jika terjd log error maka callback akan mengirim notifikasi ke programmer

use AddHook()

ex di hook_test.go we make struct implementasi dr hook, hook ini bth implementasi func Levels & Fire<br>
1. Levels return level apa saja / event apa saja as trigger
2. Fire, inilah callback yg akan dirun/eksekusi/jlnkan ketika event dari Level terjd. fire bth param berupa entry
