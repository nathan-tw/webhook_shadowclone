# webhook_shadowclone

A webhook forwarder to multiple urls. Shadow clone is a Ninjutsu that copy oneself multiple times.

## Usage 

Add webhook urls in `webhook_list` file

```
https://url1/api1
https://url1/api2
https://url2/api1
...
...
```

Run the shadowclone server
```bash
git clone https://github.com/nathan-tw/webhook_shadowclone.git
cd webhook_shadowclone.git
go build && ./webhook_shadowclone
```
