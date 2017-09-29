### Problem3: Web page word frequency counter

```
wordfreq [-help|-h]
wordfreq -urls=<comma-separated-one-or-more-urls>
```

wordfreq reads multiple URLs and generates for each of them a word frequency table. Output is stored in url<N>.txt files and formatted as follows:

```
  url: <url>
    word1: count
    word2: count
    ...
```

Run
cd samples/problem3/wordfreq
go get ./...
go build
./problem3

**Console output**

```
Mini:wordfreq doru$ ./wordfreq -urls https://www.aporeto.com/caesars-story-microbursting-cloud-security-services/,https://www.aporeto.com/trust-centric-security-authentication-and-authorization/,https://www.aporeto.com/accelerating-business-devops-and-microservices-part-ii-running-safer/
Performing text analysis on:
	https://www.aporeto.com/caesars-story-microbursting-cloud-security-services/
	https://www.aporeto.com/trust-centric-security-authentication-and-authorization/
	https://www.aporeto.com/accelerating-business-devops-and-microservices-part-ii-running-safer/

[2] Processing: https://www.aporeto.com/accelerating-business-devops-and-microservices-part-ii-running-safer/
[0] Processing: https://www.aporeto.com/caesars-story-microbursting-cloud-security-services/
[1] Processing: https://www.aporeto.com/trust-centric-security-authentication-and-authorization/
[2] Wrote: url2.txt
[0] Wrote: url0.txt
[1] Wrote: url1.txt

Done!
```
```
Mini:problem3 doru$ head -5 url*.txt
==> url0.txt <==
url: https://www.aporeto.com/caesars-story-microbursting-cloud-security-services/
  hard: 1
  morning: 2
  key: 1
  want: 1

==> url1.txt <==
url: https://www.aporeto.com/trust-centric-security-authentication-and-authorization/
  partners: 2
  moat: 3
  mobility: 1
  source: 1

==> url2.txt <==
url: https://www.aporeto.com/accelerating-business-devops-and-microservices-part-ii-running-safer/
  keynote: 1
  security: 34
  author: 2
  years: 1

```
