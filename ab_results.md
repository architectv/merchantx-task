## Нагрузочное тестирование

```
$ cat > put.txt
{"id":1,"link":"https://docs.google.com/spreadsheets/d/e/2PACX-1vRmOaivfZYZqJCgnS6Dnjw8kLvRtgMELipP9r7m8nE_Te6N06glcNaGyNVw73f0VuKi8mgoErSploTZ/pub?output=xlsx"}
^C
$ ./ab -c 10 -n 1000 -T application/json -u put.txt localhost:8001/api/offers/
This is ApacheBench, Version 2.3 <$Revision: 1879490 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)


Server Software:        
Server Hostname:        localhost
Server Port:            8001

Document Path:          /api/offers/
Document Length:        68 bytes

Concurrency Level:      10
Time taken for tests:   86.663 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      195000 bytes
Total body sent:        307000
HTML transferred:       68000 bytes
Requests per second:    11.54 [#/sec] (mean)
Time per request:       866.634 [ms] (mean)
Time per request:       86.663 [ms] (mean, across all concurrent requests)
Transfer rate:          2.20 [Kbytes/sec] received
                        3.46 kb/s sent
                        5.66 kb/s total

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    0   0.5      0       2
Processing:   590  852 145.6    829    1698
Waiting:      590  852 145.6    828    1698
Total:        590  852 145.6    829    1698

Percentage of the requests served within a certain time (ms)
  50%    829
  66%    911
  75%    935
  80%    946
  90%    982
  95%   1053
  98%   1243
  99%   1423
 100%   1698 (longest request)
```